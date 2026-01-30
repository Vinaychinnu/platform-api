package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenPostgres(
	host, port, user, pass, name string,
) (*sql.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
