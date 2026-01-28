package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Vinaychinnu/platform-api/internal/application"
	"github.com/Vinaychinnu/platform-api/internal/domain"
)

type ProjectHandler struct {
	service *application.ProjectService
}

func NewProjectHandler(service *application.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req CreateProjectRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	project, err := h.service.CreateProject(req.ID, req.Name)
	if err != nil {
		if err == domain.ErrInvalidName {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	resp := CreateProjectResponse{
		ID:   project.ID,
		Name: project.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
