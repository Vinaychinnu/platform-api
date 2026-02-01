package infrastructure

import (
	"database/sql"

	"github.com/Vinaychinnu/platform-api/internal/domain"
)

type PostgresEnvironmentRepository struct {
	db *sql.DB
}

func NewPostgresEnvironmentRepository(db *sql.DB) *PostgresEnvironmentRepository {
	return &PostgresEnvironmentRepository{db: db}
}

func (r *PostgresEnvironmentRepository) Save(e *domain.Environment) error {
	_, err := r.db.Exec(
		`INSERT INTO environments (id, project_id, name, created_at)
		 VALUES ($1, $2, $3, $4)`,
		e.ID, e.ProjectID, e.Name, e.CreatedAt,
	)
	return err
}

func (r *PostgresEnvironmentRepository) FindByID(id string) (*domain.Environment, error) {
	row := r.db.QueryRow(
		`SELECT id, project_id, name, created_at FROM environments WHERE id = $1`,
		id,
	)

	var e domain.Environment
	if err := row.Scan(&e.ID, &e.ProjectID, &e.Name, &e.CreatedAt); err != nil {
		return nil, err
	}

	return &e, nil
}
