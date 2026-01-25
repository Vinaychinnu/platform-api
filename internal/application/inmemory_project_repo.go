package application

import (
	"errors"
	"sync"

	"github.com/Vinaychinnu/platform-api/internal/domain"
)

type InMemoryProjectRepository struct {
	mu       sync.Mutex
	projects map[string]*domain.Project
}

func NewInMemoryProjectRepository() *InMemoryProjectRepository {
	return &InMemoryProjectRepository{
		projects: make(map[string]*domain.Project),
	}
}

func (r *InMemoryProjectRepository) Save(project *domain.Project) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.projects[project.ID]; exists {
		return errors.New("project already exists")
	}

	r.projects[project.ID] = project
	return nil
}

func (r *InMemoryProjectRepository) FindByID(id string) (*domain.Project, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	project, ok := r.projects[id]
	if !ok {
		return nil, errors.New("project not found")
	}

	return project, nil
}
