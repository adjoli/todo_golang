package config

import "fmt"

const (
	EnvDatabasePath     = "TASKMANAGER_DB"
	DefaultDatabasePath = "data/tasks.db"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Path string
}

func New() (*Config, error) {
	cfg := defaultConfig()

	if err := loadEnvironment(cfg); err != nil {
		return nil, err
	}

	if err := prepareDatabasePath(cfg); err != nil {
		return nil, fmt.Errorf("prepare database path: %w", err)
	}

	return cfg, nil
}

// ----------------------------------------------
// FUNÇÕES PRIVADAS
// ----------------------------------------------
func defaultConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Path: DefaultDatabasePath,
		},
	}
}
