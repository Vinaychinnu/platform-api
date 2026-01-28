package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Vinaychinnu/platform-api/internal/application"
	"github.com/Vinaychinnu/platform-api/internal/domain"
)

type ReleaseHandler struct {
	service *application.ReleaseService
}

func NewReleaseHandler(service *application.ReleaseService) *ReleaseHandler {
	return &ReleaseHandler{service: service}
}

func (h *ReleaseHandler) CreateRelease(w http.ResponseWriter, r *http.Request) {
	var req CreateReleaseRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	release, err := h.service.CreateRelease(req.ID, req.EnvironmentID, req.Version)
	if err != nil {
		if err == domain.ErrInvalidEnvironment || err == domain.ErrInvalidVersion {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	resp := CreateReleaseResponse{
		ID:            release.ID,
		EnvironmentID: release.EnvironmentID,
		Version:       release.Version,
		Status:        string(release.Status),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
