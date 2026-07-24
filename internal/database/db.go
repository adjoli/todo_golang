// Package database gerencia a conexão com o banco SQLite e a
// criação automática do schema. Utiliza modernc.org/sqlite,
// uma implementação 100% em Go que não requer CGO.
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

// createSchema executa o SQL de criação das tabelas. A tabela tasks
// é criada apenas se não existir (IF NOT EXISTS).
func createSchema(db *sql.DB) error {
	_, err := db.Exec(createTasksTableSQL)
	return err
}

// New abre uma conexão com o banco SQLite no caminho especificado,
// verifica a acessibilidade com Ping e cria as tabelas necessárias.
// O caller é responsável por chamar db.Close() quando não for
// mais necessário.
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
