package application

import (
	"errors"
	"sync"

	"github.com/Vinaychinnu/platform-api/internal/domain"
)

type InMemoryReleaseRepository struct {
	mu       sync.Mutex
	releases map[string]*domain.Release
}

func NewInMemoryReleaseRepository() *InMemoryReleaseRepository {
	return &InMemoryReleaseRepository{
		releases: make(map[string]*domain.Release),
	}
}

func (r *InMemoryReleaseRepository) Save(rel *domain.Release) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.releases[rel.ID]; exists {
		return errors.New("release already exists")
	}

	r.releases[rel.ID] = rel
	return nil
}

func (r *InMemoryReleaseRepository) FindByID(id string) (*domain.Release, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	rel, ok := r.releases[id]
	if !ok {
		return nil, errors.New("release not found")
	}

	return rel, nil
}
