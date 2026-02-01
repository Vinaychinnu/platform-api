package infrastructure

import "database/sql"

func NewPostgresRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Projects:     NewPostgresProjectRepository(db),
		Environments: NewPostgresEnvironmentRepository(db),
		Releases:     NewPostgresReleaseRepository(db),
	}
}
