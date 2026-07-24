package config

import (
	"errors"
	"os"
	"strings"
)

// loadEnvironment aplica as sobrescritas de variáveis de ambiente
// sobre a configuração padrão. Retorna erro se a variável TASKMANAGER_DB
// estiver definida mas vazia.
func loadEnvironment(cfg *Config) error {
	if path, ok := os.LookupEnv(EnvDatabasePath); ok {
		if strings.TrimSpace(path) == "" {
			return errors.New("database path cannot be empty")
		}

		cfg.Database.Path = path
	}

	return nil
}
