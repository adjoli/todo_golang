package config

import (
	"os"
	"path/filepath"
)

func prepareDatabasePath(cfg *Config) error {
	dir := filepath.Dir(cfg.Database.Path)

	if dir == "." {
		return nil
	}

	return os.MkdirAll(dir, 0o755)
}
