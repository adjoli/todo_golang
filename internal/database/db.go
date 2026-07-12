package database

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const (
	dbDir = "data"
	// dbFile = "data/tasks.db"
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

func Open(path string) (*sql.DB, error) {
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	if err := createSchema(db); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
