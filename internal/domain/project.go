package domain

import "time"

type Project struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

func NewProject(id, name string) (*Project, error) {
	if name == "" {
		return nil, ErrInvalidName
	}

	return &Project{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}, nil
}
