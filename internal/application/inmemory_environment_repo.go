package application

import (
	"errors"
	"sync"

	"github.com/Vinaychinnu/platform-api/internal/domain"
)

type InMemoryEnvironmentRepository struct {
	mu           sync.Mutex
	environments map[string]*domain.Environment
}

func NewInMemoryEnvironmentRepository() *InMemoryEnvironmentRepository {
	return &InMemoryEnvironmentRepository{
		environments: make(map[string]*domain.Environment),
	}
}

func (r *InMemoryEnvironmentRepository) Save(env *domain.Environment) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.environments[env.ID]; exists {
		return errors.New("environment already exists")
	}

	r.environments[env.ID] = env
	return nil
}

func (r *InMemoryEnvironmentRepository) FindByID(id string) (*domain.Environment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	env, ok := r.environments[id]
	if !ok {
		return nil, errors.New("environment not found")
	}

	return env, nil
}
