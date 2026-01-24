package application

import "github.com/Vinaychinnu/platform-api/internal/domain"

type ReleaseService struct {
	repo ReleaseRepository
}

func NewReleaseService(repo ReleaseRepository) *ReleaseService {
	return &ReleaseService{repo: repo}
}

func (s *ReleaseService) CreateRelease(id, envID, version string) (*domain.Release, error) {
	release, err := domain.NewRelease(id, envID, version)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(release); err != nil {
		return nil, err
	}

	return release, nil
}
