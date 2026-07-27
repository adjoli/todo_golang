package repository

import "fmt"

// dialect define o contrato para adaptações de SQL específicas
// de cada driver de banco de dados.
type dialect interface {
	// placeholder retorna o placeholder correto para o índice
	// informado (1-based). SQLite usa "?"; Postgres usa "$1", "$2", etc.
	placeholder(index int) string

	// supportsLastInsertID indica se o driver suporta
	// result.LastInsertId() após um INSERT.
	supportsLastInsertID() bool
}

// NewDialect retorna o dialect correspondente ao driver informado.
func NewDialect(driver string) (dialect, error) {
	switch driver {
	case "sqlite":
		return &sqliteDialect{}, nil
	case "postgres":
		return &postgresDialect{}, nil
	default:
		return nil, fmt.Errorf("unsupported driver %q", driver)
	}
}

// sqliteDialect implementa dialect para SQLite.
type sqliteDialect struct{}

func (d *sqliteDialect) placeholder(_ int) string {
	return "?"
}

func (d *sqliteDialect) supportsLastInsertID() bool {
	return true
}

// postgresDialect implementa dialect para Postgres.
type postgresDialect struct{}

func (d *postgresDialect) placeholder(index int) string {
	return fmt.Sprintf("$%d", index)
}

func (d *postgresDialect) supportsLastInsertID() bool {
	return false
}
