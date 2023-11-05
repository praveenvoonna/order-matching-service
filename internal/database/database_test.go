package database

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestPostgresDatabase_Connect(t *testing.T) {
	var p PostgresDatabase
	db := p.Connect()
	if db == nil {
		t.Error("Connection failed")
	}
}

func TestPostgresDatabase_Execute(t *testing.T) {
	var p PostgresDatabase
	p.Connect()

	query := "CREATE TABLE test_table(id SERIAL PRIMARY KEY, name VARCHAR);"
	_, err := p.Execute(query)
	if err != nil {
		t.Error("Execute function failed")
	}

	_, err = p.Execute("DROP TABLE test_table;")
	if err != nil {
		t.Error("Execute function failed")
	}
}

func TestPostgresDatabase_QueryRows(t *testing.T) {
	var p PostgresDatabase
	p.Connect()

	query := "SELECT 1"
	rows, err := p.QueryRows(query)
	if err != nil {
		t.Error("QueryRows function failed")
	}
	defer rows.Close()
}

func TestPostgresDatabase_QueryRow(t *testing.T) {
	var p PostgresDatabase
	p.Connect()

	query := "SELECT 1"
	row := p.QueryRow(query)
	if row == nil {
		t.Error("QueryRow function failed")
	}
}
