package infrastructure

import (
	"database/sql"

	"github.com/Vinaychinnu/platform-api/internal/domain"
)

type PostgresReleaseRepository struct {
	db *sql.DB
}

func NewPostgresReleaseRepository(db *sql.DB) *PostgresReleaseRepository {
	return &PostgresReleaseRepository{db: db}
}

func (r *PostgresReleaseRepository) Save(rel *domain.Release) error {
	_, err := r.db.Exec(
		`INSERT INTO releases (id, environment_id, version, status, created_at)
		 VALUES ($1, $2, $3, $4, $5)`,
		rel.ID, rel.EnvironmentID, rel.Version, rel.Status, rel.CreatedAt,
	)
	return err
}

func (r *PostgresReleaseRepository) FindByID(id string) (*domain.Release, error) {
	row := r.db.QueryRow(
		`SELECT id, environment_id, version, status, created_at
		 FROM releases WHERE id = $1`,
		id,
	)

	var rls domain.Release
	if err := row.Scan(
		&rls.ID,
		&rls.EnvironmentID,
		&rls.Version,
		&rls.Status,
		&rls.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &rls, nil
}
