package application

import "github.com/Vinaychinnu/platform-api/internal/domain"

type ProjectRepository interface {
	Save(project *domain.Project) error
	FindByID(id string) (*domain.Project, error)
}

type EnvironmentRepository interface {
	Save(env *domain.Environment) error
	FindByID(id string) (*domain.Environment, error)
}

type ReleaseRepository interface {
	Save(release *domain.Release) error
	FindByID(id string) (*domain.Release, error)
}
