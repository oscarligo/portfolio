package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"backend/internal/repository"
	"backend/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProjectHandler struct {
	Store *store.Store
}

func NewProjectHandler(store *store.Store) *ProjectHandler {
	return &ProjectHandler{Store: store}
}

// Helper function to convert pgtype.Text to *string, handling NULL values correctly
func pgtypeToStringPtr(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}
	return &t.String
}

// Helper function to convert string to pgtype.Text, treating empty strings as NULL
func stringToPgtypeText(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{Valid: false} // Se guardará como NULL en PostgreSQL
	}
	return pgtype.Text{String: s, Valid: true}
}

// Helper function to safely convert an interface{} to json.RawMessage, handling both []byte and other types
func safeJSONRawMessage(v interface{}) json.RawMessage {
	if bytes, ok := v.([]byte); ok {
		return json.RawMessage(bytes)
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		return json.RawMessage("[]")
	}
	return json.RawMessage(bytes)
}

// Struct espejo global para la respuesta HTTP unificada
type fullProjectResponse struct {
	ID             int32           `json:"id"`
	Title          string          `json:"title"`
	TranslationKey string          `json:"translation_key"`
	RepoUrl        *string         `json:"repo_url"`
	LiveUrl        *string         `json:"live_url"`
	VideoUrl       *string         `json:"video_url"`
	CreatedAt      time.Time       `json:"created_at"`
	Featured       bool            `json:"featured"`
	Images         json.RawMessage `json:"images"`
	Technologies   json.RawMessage `json:"technologies"`
}

// GET /api/projects
func (h *ProjectHandler) ListProjects(w http.ResponseWriter, r *http.Request) {
	rows, err := h.Store.ListProjects(r.Context())
	if err != nil {
		http.Error(w, "Error while fetching projects: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]fullProjectResponse, len(rows))
	for i, row := range rows {
		response[i] = fullProjectResponse{
			ID:             row.ID,
			Title:          row.Title,
			TranslationKey: row.TranslationKey,
			RepoUrl:        pgtypeToStringPtr(row.RepoUrl),
			LiveUrl:        pgtypeToStringPtr(row.LiveUrl),
			VideoUrl:       pgtypeToStringPtr(row.VideoUrl),
			CreatedAt:      row.CreatedAt.Time,
			Featured:       row.Featured,
			Images:         safeJSONRawMessage(row.Images),
			Technologies:   safeJSONRawMessage(row.Technologies),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GET /api/projects/featured
func (h *ProjectHandler) ListFeaturedProjects(w http.ResponseWriter, r *http.Request) {
	rows, err := h.Store.ListFeaturedProjects(r.Context())
	if err != nil {
		http.Error(w, "Error while fetching featured projects: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]fullProjectResponse, len(rows))
	for i, row := range rows {
		response[i] = fullProjectResponse{
			ID:             row.ID,
			Title:          row.Title,
			TranslationKey: row.TranslationKey,
			RepoUrl:        pgtypeToStringPtr(row.RepoUrl),
			LiveUrl:        pgtypeToStringPtr(row.LiveUrl),
			VideoUrl:       pgtypeToStringPtr(row.VideoUrl),
			CreatedAt:      row.CreatedAt.Time,
			Featured:       row.Featured,
			Images:         safeJSONRawMessage(row.Images),
			Technologies:   safeJSONRawMessage(row.Technologies),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GET /api/projects/{id}
func (h *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	// Get the project ID from the URL path and convert it to int32
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, "ID de proyecto inválido", http.StatusBadRequest)
		return
	}

	row, err := h.Store.GetProject(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) || errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error while fetching the project: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := fullProjectResponse{
		ID:             row.ID,
		Title:          row.Title,
		TranslationKey: row.TranslationKey,
		RepoUrl:        pgtypeToStringPtr(row.RepoUrl),
		LiveUrl:        pgtypeToStringPtr(row.LiveUrl),
		VideoUrl:       pgtypeToStringPtr(row.VideoUrl),
		CreatedAt:      row.CreatedAt.Time,
		Featured:       row.Featured,
		Images:         safeJSONRawMessage(row.Images),
		Technologies:   safeJSONRawMessage(row.Technologies),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

type createProjectRequest struct {
	Title          string   `json:"title"`
	TranslationKey string   `json:"translation_key"`
	RepoUrl        string   `json:"repo_url"`
	LiveUrl        string   `json:"live_url"`
	VideoUrl       string   `json:"video_url"`
	Featured       bool     `json:"featured"`
	Images         []string `json:"images"`
	Technologies   []int32  `json:"technologies"`
}

// POST /api/projects
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req createProjectRequest

	// Decode JSON body into the request struct, handling errors gracefully
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields (title, description_short, description_long) and return a 400 Bad Request if any are missing
	if req.Title == "" || req.TranslationKey == "" {
		http.Error(w, "Fields 'title' and 'translation_key' are required", http.StatusBadRequest)
		return
	}

	// Map the incoming request to the CreateProjectTxParams struct, transforming optional string fields to pgtype.Text using our helper function
	arg := store.CreateProjectTxParams{
		CreateProjectParams: repository.CreateProjectParams{
			Title:          req.Title,
			TranslationKey: req.TranslationKey,
			RepoUrl:        stringToPgtypeText(req.RepoUrl),
			LiveUrl:        stringToPgtypeText(req.LiveUrl),
			VideoUrl:       stringToPgtypeText(req.VideoUrl),
			Featured:       req.Featured,
		},
		ImageURLs:     req.Images,
		TechnologyIDs: req.Technologies,
	}

	// Run tranction
	project, err := h.Store.CreateProjectTx(r.Context(), arg)
	if err != nil {
		http.Error(w, "Error while creating the project: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// API response with the created project
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}
