package application

import "github.com/Vinaychinnu/platform-api/internal/domain"

type ProjectService struct {
	repo ProjectRepository
}

func NewProjectService(repo ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) CreateProject(id, name string) (*domain.Project, error) {
	project, err := domain.NewProject(id, name)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(project); err != nil {
		return nil, err
	}

	return project, nil
}
