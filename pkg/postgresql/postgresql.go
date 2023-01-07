// Package postgres implements postgres connection.
package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Postgres -.
type Postgres struct {
	Db *sql.DB
}

// New -.
func New(connectionString string) (*Postgres, error) {
	db, err := sql.Open("postgres", connectionString)
	pg := &Postgres{db}
	if err != nil {
		return pg, err
	}
	err = pg.Db.Ping()
	if err != nil {
		return pg, err
	}
	return pg, nil
}

// Close -.
func (pg *Postgres) Close() {
	if pg.Db != nil {
		pg.Db.Close()
	}
}
