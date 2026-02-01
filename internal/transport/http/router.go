package http

import (
	"net/http"

	"database/sql"
	"github.com/Vinaychinnu/platform-api/internal/application"
	"github.com/Vinaychinnu/platform-api/internal/infrastructure"
	"github.com/Vinaychinnu/platform-api/internal/transport/http/handlers"
)

func NewRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	// health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// repositories (composition root)
	repos := infrastructure.NewPostgresRepositories(db)

	// services
	projectService := application.NewProjectService(repos.Projects)
	envService := application.NewEnvironmentService(repos.Environments)
	releaseService := application.NewReleaseService(repos.Releases)

	// handlers
	projectHandler := handlers.NewProjectHandler(projectService)
	envHandler := handlers.NewEnvironmentHandler(envService)
	releaseHandler := handlers.NewReleaseHandler(releaseService)

	// routes
	mux.HandleFunc("/projects", projectHandler.CreateProject)
	mux.HandleFunc("/environments", envHandler.CreateEnvironment)
	mux.HandleFunc("/releases", releaseHandler.CreateRelease)

	return mux
}
