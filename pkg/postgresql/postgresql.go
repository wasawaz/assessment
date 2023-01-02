package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresqlDB struct {
	Db *sql.DB
}

func New(connectionString string) (*PostgresqlDB, error) {
	db, err := sql.Open("postgres", connectionString)
	pg := &PostgresqlDB{db}
	if err != nil {
		return pg, err
	}
	err = pg.Db.Ping()
	if err != nil {
		return pg, err
	}
	return pg, nil
}

func (pg *PostgresqlDB) Close() {
	if pg.Db != nil {
		pg.Db.Close()
	}
}
