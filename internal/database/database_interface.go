package database

import "database/sql"

type Database interface {
	Connect() *sql.DB
	Execute(query string, args ...interface{}) (sql.Result, error)
	QueryRows(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}
