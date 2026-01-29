package infrastructure

import "github.com/Vinaychinnu/platform-api/internal/application"

type Repositories struct {
	Projects     application.ProjectRepository
	Environments application.EnvironmentRepository
	Releases     application.ReleaseRepository
}
