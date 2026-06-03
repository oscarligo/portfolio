package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"backend/internal/handler"
	customMiddleware "backend/internal/middleware"
	"backend/internal/store"
	"backend/internal/utils"
)

func main() {

	// Url from environment variable
	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		log.Fatal("DB_SOURCE Not set in environment variables")
	}

	// JWT secret from environment variable
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET Not set in environment variables")
	}

	ctx := context.Background()

	// Create connection pool
	conn, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("Error while creating connection pool: ", err)
	}
	defer conn.Close()

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal("Error while pinging database: ", err)
	}

	// Store and handlers
	projectStore := store.NewStore(conn)
	r2Uploader, err := utils.NewR2Uploader()
	if err != nil {
		log.Fatalf("Error while creating R2 uploader: %v", err)
	}
	log.Println("R2 uploader initialized successfully.")

	projectHandler := handler.NewProjectHandler(projectStore, r2Uploader)

	authHandler := handler.NewAuthHandler(projectStore, jwtSecret)

	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminUsername != "" && adminPassword != "" {
		err := projectStore.CreateAdminUserIfNotExist(ctx, adminUsername, adminPassword)
		if err != nil {
			log.Printf("Error while creating admin user: %v\n", err)
		} else {
			log.Println("Credentials for admin user are set.")
		}
	}

	// Router setup
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		// Public routes
		r.Get("/projects", projectHandler.ListProjects)
		r.Get("/health", handler.HealthHandler)
		r.Post("/auth/login", authHandler.Login)

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(customMiddleware.AuthRequired(jwtSecret))

			r.Post("/projects", projectHandler.CreateProject)
			r.Delete("/projects/{id}", projectHandler.DeleteProject)
		})
	})

	// Start server
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	// Log server start
	log.Printf("Server running on port %s...", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, r))
}
