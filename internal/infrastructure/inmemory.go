package infrastructure

import "github.com/Vinaychinnu/platform-api/internal/application"

func NewInMemoryRepositories() *Repositories {
	return &Repositories{
		Projects:     application.NewInMemoryProjectRepository(),
		Environments: application.NewInMemoryEnvironmentRepository(),
		Releases:     application.NewInMemoryReleaseRepository(),
	}
}
