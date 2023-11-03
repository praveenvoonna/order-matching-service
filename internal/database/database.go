package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	DB *sql.DB
}

func (p *PostgresDatabase) ConnectToDatabase() *sql.DB {
	var err error
	p.DB, err = sql.Open("postgres", "postgres://postgres:112233@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = p.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return p.DB
}

func (p *PostgresDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := p.DB.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func (p *PostgresDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
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
