package infrastructure

import (
	"database/sql"

	"github.com/Vinaychinnu/platform-api/internal/domain"
)

type PostgresProjectRepository struct {
	db *sql.DB
}

func NewPostgresProjectRepository(db *sql.DB) *PostgresProjectRepository {
	return &PostgresProjectRepository{db: db}
}

func (r *PostgresProjectRepository) Save(p *domain.Project) error {
	_, err := r.db.Exec(
		`INSERT INTO projects (id, name, created_at) VALUES ($1, $2, $3)`,
		p.ID, p.Name, p.CreatedAt,
	)
	return err
}

func (r *PostgresProjectRepository) FindByID(id string) (*domain.Project, error) {
	row := r.db.QueryRow(
		`SELECT id, name, created_at FROM projects WHERE id = $1`,
		id,
	)

	var p domain.Project
	if err := row.Scan(&p.ID, &p.Name, &p.CreatedAt); err != nil {
		return nil, err
	}

	return &p, nil
}
