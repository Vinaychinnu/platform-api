package application

import "github.com/Vinaychinnu/platform-api/internal/domain"

type EnvironmentService struct {
	repo EnvironmentRepository
}

func NewEnvironmentService(repo EnvironmentRepository) *EnvironmentService {
	return &EnvironmentService{repo: repo}
}

func (s *EnvironmentService) CreateEnvironment(id, projectID, name string) (*domain.Environment, error) {
	env, err := domain.NewEnvironment(id, projectID, name)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(env); err != nil {
		return nil, err
	}

	return env, nil
}
