package app

import "database/sql"

// Database is an interface that abstracts database interactions.
type Database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Commit() error
	Rollback() error
}

// RealDatabase is a struct that implements the Database interface using a real *sql.Tx.
type RealDatabase struct {
	tx *sql.Tx
}

func (r *RealDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	return r.tx.Exec(query, args...)
}

func (r *RealDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
	return r.tx.QueryRow(query, args...)
}

func (r *RealDatabase) Commit() error {
	return r.tx.Commit()
}

func (r *RealDatabase) Rollback() error {
	return r.tx.Rollback()
}
