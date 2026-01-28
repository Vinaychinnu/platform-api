package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Vinaychinnu/platform-api/internal/application"
	"github.com/Vinaychinnu/platform-api/internal/domain"
)

type EnvironmentHandler struct {
	service *application.EnvironmentService
}

func NewEnvironmentHandler(service *application.EnvironmentService) *EnvironmentHandler {
	return &EnvironmentHandler{service: service}
}

func (h *EnvironmentHandler) CreateEnvironment(w http.ResponseWriter, r *http.Request) {
	var req CreateEnvironmentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	env, err := h.service.CreateEnvironment(req.ID, req.ProjectID, req.Name)
	if err != nil {
		if err == domain.ErrInvalidEnvironment {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	resp := CreateEnvironmentResponse{
		ID:        env.ID,
		ProjectID: env.ProjectID,
		Name:      env.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
