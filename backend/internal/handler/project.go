package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"backend/internal/repository"
	"backend/internal/store"
	"backend/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProjectHandler struct {
	Store      *store.Store
	R2Uploader *utils.R2Uploader
}

func NewProjectHandler(store *store.Store, uploader *utils.R2Uploader) *ProjectHandler {
	return &ProjectHandler{Store: store, R2Uploader: uploader}
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

func parseProjectID(r *http.Request) (int32, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid project id: %w", err)
	}

	return int32(id), nil
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
	id, err := parseProjectID(r)
	if err != nil {
		http.Error(w, "ID de proyecto inválido", http.StatusBadRequest)
		return
	}

	row, err := h.Store.GetProject(r.Context(), id)
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
	// Parse multipart form with a  max memory limit
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Extract form values
	title := r.FormValue("title")
	translationKey := r.FormValue("translation_key")
	repoUrl := r.FormValue("repo_url")
	liveUrl := r.FormValue("live_url")
	videoUrl := r.FormValue("video_url")
	featuredStr := r.FormValue("featured")

	// Validate required fields
	if title == "" || translationKey == "" {
		http.Error(w, "Fields 'title' and 'translation_key' are required", http.StatusBadRequest)
		return
	}

	featured := featuredStr == "true"

	// Extract technology IDs from form values
	techStrings := r.MultipartForm.Value["technologies"]
	var technologyIDs []int32

	// Convert string IDs to int32
	for _, ts := range techStrings {
		var id int
		if _, err := fmt.Sscanf(ts, "%d", &id); err == nil {
			technologyIDs = append(technologyIDs, int32(id))
		}
	}

	// Process uploaded images

	files := r.MultipartForm.File["images"]
	var imageURLs []string

	// Iterate over uploaded files, upload them to R2, and collect their URLs
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error opening uploaded file: "+err.Error(), http.StatusInternalServerError)
			return
		}

		objectKey := fmt.Sprintf("projects/%d_%s", time.Now().UnixNano(), fileHeader.Filename)
		contentType := fileHeader.Header.Get("Content-Type")

		url, err := h.R2Uploader.UploadImage(r.Context(), objectKey, file, contentType)
		file.Close()
		if err != nil {
			http.Error(w, "Error uploading image to R2: "+err.Error(), http.StatusInternalServerError)
			return
		}
		// Append the returned URL to the list of image URLs for this project
		imageURLs = append(imageURLs, url)
	}

	// Create the project in the database.
	arg := store.CreateProjectTxParams{
		CreateProjectParams: repository.CreateProjectParams{
			Title:          title,
			TranslationKey: translationKey,
			RepoUrl:        stringToPgtypeText(repoUrl),
			LiveUrl:        stringToPgtypeText(liveUrl),
			VideoUrl:       stringToPgtypeText(videoUrl),
			Featured:       featured,
		},
		ImageURLs:     imageURLs,
		TechnologyIDs: technologyIDs,
	}

	project, err := h.Store.CreateProjectTx(r.Context(), arg)
	if err != nil {
		http.Error(w, "Error while creating the project in database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Return the created project as JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

// DELETE /api/projects/{id}
func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id, err := parseProjectID(r)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	_, err = h.Store.GetProject(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) || errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error while fetching the project: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.Store.DeleteProject(r.Context(), id)
	if err != nil {
		http.Error(w, "Error while deleting the project: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
