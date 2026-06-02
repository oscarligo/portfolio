package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"slices"
	"syscall"
	"time"

	"backend/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type application struct {
	logger  *slog.Logger
	db      *pgxpool.Pool
	queries *repository.Queries
}

type projectResponse struct {
	ID               int32  `json:"id"`
	Title            string `json:"title"`
	DescriptionShort string `json:"descriptionShort"`
	DescriptionLong  string `json:"descriptionLong"`
	RepoURL          string `json:"repoUrl"`
	LiveURL          string `json:"liveUrl"`
	CreatedAt        string `json:"createdAt,omitempty"`
	Featured         bool   `json:"featured"`
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	serverAddress := getEnv("SERVER_ADDRESS", "0.0.0.0:8080")
	dbSource := os.Getenv("DB_SOURCE")
	migrationsDir := getEnv("MIGRATIONS_DIR", "database/migration")

	if dbSource == "" {
		logger.Error("DB_SOURCE is required")
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := connectDatabase(ctx, logger, dbSource)
	if err != nil {
		logger.Error("unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := ensureSchema(ctx, db, migrationsDir); err != nil {
		logger.Error("unable to initialize database schema", "error", err)
		os.Exit(1)
	}

	app := &application{
		logger:  logger,
		db:      db,
		queries: repository.New(db),
	}

	server := &http.Server{
		Addr:              serverAddress,
		Handler:           app.routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 1)

	go func() {
		logger.Info("backend listening", "addr", serverAddress)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			logger.Error("server stopped unexpectedly", "error", err)
			os.Exit(1)
		}
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("graceful shutdown failed", "error", err)
			os.Exit(1)
		}
	}
}

func connectDatabase(ctx context.Context, logger *slog.Logger, dbSource string) (*pgxpool.Pool, error) {
	var lastErr error

	for attempt := 1; attempt <= 15; attempt++ {
		pool, err := pgxpool.New(ctx, dbSource)
		if err != nil {
			return nil, fmt.Errorf("parse db config: %w", err)
		}

		pingCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
		err = pool.Ping(pingCtx)
		cancel()

		if err == nil {
			return pool, nil
		}

		pool.Close()
		lastErr = err

		logger.Warn("database not ready yet", "attempt", attempt, "error", err)

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}

	return nil, fmt.Errorf("database never became ready: %w", lastErr)
}

func ensureSchema(ctx context.Context, db *pgxpool.Pool, migrationsDir string) error {
	pattern := filepath.Join(migrationsDir, "*.up.sql")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("scan migrations: %w", err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no migration files found in %s", migrationsDir)
	}

	slices.Sort(files)

	for _, file := range files {
		query, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("read migration %s: %w", file, err)
		}

		if _, err := db.Exec(ctx, string(query)); err != nil {
			return fmt.Errorf("apply migration %s: %w", file, err)
		}
	}

	return nil
}

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.handleRoot)
	mux.HandleFunc("/healthz", app.handleHealth)
	mux.HandleFunc("/api/projects", app.handleProjects)
	return mux
}

func (app *application) handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.writeJSON(w, http.StatusOK, map[string]string{
		"service": "portfolio-backend",
		"status":  "running",
	})
}

func (app *application) handleHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := app.db.Ping(ctx); err != nil {
		app.writeJSON(w, http.StatusServiceUnavailable, map[string]string{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	app.writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (app *application) handleProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	projects, err := app.queries.ListProjects(r.Context())
	if err != nil {
		app.logger.Error("list projects failed", "error", err)
		app.writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "unable to list projects",
		})
		return
	}

	response := make([]projectResponse, 0, len(projects))
	for _, project := range projects {
		item := projectResponse{
			ID:               project.ID,
			Title:            project.Title,
			DescriptionShort: project.DescriptionShort,
			DescriptionLong:  project.DescriptionLong,
			RepoURL:          project.RepoUrl,
			LiveURL:          project.LiveUrl,
			Featured:         project.Featured,
		}

		if project.CreatedAt.Valid {
			item.CreatedAt = project.CreatedAt.Time.UTC().Format(time.RFC3339)
		}

		response = append(response, item)
	}

	app.writeJSON(w, http.StatusOK, response)
}

func (app *application) writeJSON(w http.ResponseWriter, status int, v any) {
	payload, err := json.Marshal(v)
	if err != nil {
		app.logger.Error("json marshal failed", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(payload)
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
