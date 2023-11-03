package database

import "database/sql"

type Database interface {
	ConnectToDatabase() *sql.DB
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}
