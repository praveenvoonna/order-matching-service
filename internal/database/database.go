package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	DB *sql.DB
}

func (p *PostgresDatabase) Connect() *sql.DB {
	var err error
	// change below connection string as sql.Open("postgres", "database://username:password@localhost/postgres?sslmode=disable")
	p.DB, err = sql.Open("postgres", "postgres://postgres:112233@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = p.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return p.DB
}

func (p *PostgresDatabase) Execute(query string, args ...interface{}) (sql.Result, error) {
	result, err := p.DB.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func (p *PostgresDatabase) QueryRows(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := p.DB.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	return rows, nil
}

func (p *PostgresDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
	row := p.DB.QueryRow(query, args...)
	return row
}
