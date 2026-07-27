// Package database gerencia a conexão com o banco de dados e a
// criação automática do schema. Suporta SQLite (modernc.org/sqlite)
// e Postgres (pgx) como drivers.
package database

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5/stdlib"
	_ "modernc.org/sqlite"
)

func init() {
	sql.Register("postgres", stdlib.GetDefaultDriver())
}

// DDL por driver.
const (
	createTasksTableSQLite = `
CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	completed BOOLEAN NOT NULL DEFAULT FALSE,
	created_at DATETIME NOT NULL
);`

	createTasksTablePostgres = `
CREATE TABLE IF NOT EXISTS tasks (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	completed BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMPTZ NOT NULL
);`
)

// createSchema executa o SQL de criação das tabelas de acordo
// com o driver informado.
func createSchema(db *sql.DB, driver string) error {
	var query string

	switch driver {
	case "sqlite":
		query = createTasksTableSQLite
	case "postgres":
		query = createTasksTablePostgres
	default:
		return fmt.Errorf("unsupported driver %q for schema creation", driver)
	}

	_, err := db.Exec(query)
	return err
}

// New abre uma conexão com o banco de dados utilizando o driver
// e a string de conexão fornecidos, verifica a acessibilidade
// com Ping e cria as tabelas necessárias.
// O caller é responsável por chamar db.Close() quando não for
// mais necessário.
func New(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("open %s database: %w", driver, err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping %s database: %w", driver, err)
	}

	if err := createSchema(db, driver); err != nil {
		db.Close()
		return nil, fmt.Errorf("create %s database tables: %w", driver, err)
	}

	return db, nil
}
