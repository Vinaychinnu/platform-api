package domain

import "time"

type Environment struct {
	ID        string
	ProjectID string
	Name      string
	CreatedAt time.Time
}

func NewEnvironment(id, projectID, name string) (*Environment, error) {
	if projectID == "" {
		return nil, ErrInvalidEnvironment
	}

	if name == "" {
		return nil, ErrInvalidEnvironment
	}

	return &Environment{
		ID:        id,
		ProjectID: projectID,
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}, nil
}
