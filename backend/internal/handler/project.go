package handler

import (
	"encoding/json"
	"net/http"

	"backend/internal/store"
)

type ProjectHandler struct {
	Store *store.Store
}

func NewProjectHandler(store *store.Store) *ProjectHandler {
	return &ProjectHandler{Store: store}
}

// GET /api/projects
func (h *ProjectHandler) ListProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.Store.ListProjects(r.Context())
	if err != nil {
		http.Error(w, "Error while fetching projects: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
