package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

const createTasksTableSQL = `
CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	completed BOOLEAN NOT NULL DEFAULT FALSE,
	created_at DATETIME NOT NULL
);
`

func createSchema(db *sql.DB) error {
	_, err := db.Exec(createTasksTableSQL)
	return err
}

func New(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping sqlite database: %w", err)
	}

	if err := createSchema(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("create database tables: %w", err)
	}

	return db, nil
}
