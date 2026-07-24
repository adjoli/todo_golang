package config

import (
	"os"
	"path/filepath"
)

// prepareDatabasePath cria o diretório pai do caminho do banco de dados
// caso não exista. Se o caminho não contém diretório (ex: "tasks.db"),
// a função é uma operação no-op.
func prepareDatabasePath(cfg *Config) error {
	dir := filepath.Dir(cfg.Database.Path)

	if dir == "." {
		return nil
	}

	return os.MkdirAll(dir, 0o755)
}
