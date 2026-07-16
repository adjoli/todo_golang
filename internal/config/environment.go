package config

import (
	"errors"
	"os"
	"strings"
)

func loadEnvironment(cfg *Config) error {
	if path, ok := os.LookupEnv(EnvDatabasePath); ok {
		if strings.TrimSpace(path) == "" {
			return errors.New("database path cannot be empty")
		}

		cfg.Database.Path = path
	}

	return nil
}
