package http

import (
	"net/http"

	"github.com/Vinaychinnu/platform-api/internal/application"
	"github.com/Vinaychinnu/platform-api/internal/transport/http/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// project wiring (temporary in-memory)
	projectRepo := application.NewInMemoryProjectRepository()
	projectService := application.NewProjectService(projectRepo)
	projectHandler := handlers.NewProjectHandler(projectService)

	mux.HandleFunc("/projects", projectHandler.CreateProject)

	// environment wiring (temporary in-memory)
	envRepo := application.NewInMemoryEnvironmentRepository()
	envService := application.NewEnvironmentService(envRepo)
	envHandler := handlers.NewEnvironmentHandler(envService)

	mux.HandleFunc("/environments", envHandler.CreateEnvironment)

	// release wiring (temporary in-memory)
	releaseRepo := application.NewInMemoryReleaseRepository()
	releaseService := application.NewReleaseService(releaseRepo)
	releaseHandler := handlers.NewReleaseHandler(releaseService)

	mux.HandleFunc("/releases", releaseHandler.CreateRelease)

	return mux
}
