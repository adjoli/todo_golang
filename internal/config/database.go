package config

import (
	"os"
	"path/filepath"
)

// prepareDatabasePath cria o diretório pai do caminho do banco de dados
// caso não exista. A operação é um no-op para drivers que não utilizam
// arquivos (ex: Postgres).
func prepareDatabasePath(cfg *Config) error {
	if cfg.Database.Driver != "sqlite" {
		return nil
	}

	dir := filepath.Dir(cfg.Database.DSN)

	if dir == "." {
		return nil
	}

	return os.MkdirAll(dir, 0o755)
}
