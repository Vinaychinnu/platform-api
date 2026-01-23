package domain

import "time"

type ReleaseStatus string

const (
	ReleasePending  ReleaseStatus = "pending"
	ReleaseDeployed ReleaseStatus = "deployed"
	ReleaseFailed   ReleaseStatus = "failed"
)

type Release struct {
	ID            string
	EnvironmentID string
	Version       string
	Status        ReleaseStatus
	CreatedAt     time.Time
}

func NewRelease(id, envID, version string) (*Release, error) {
	if envID == "" {
		return nil, ErrInvalidEnvironment
	}

	if version == "" {
		return nil, ErrInvalidVersion
	}

	return &Release{
		ID:            id,
		EnvironmentID: envID,
		Version:       version,
		Status:        ReleasePending,
		CreatedAt:     time.Now().UTC(),
	}, nil
}
