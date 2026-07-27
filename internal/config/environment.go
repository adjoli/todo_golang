package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// validDrivers contém os drivers de banco de dados suportados.
var validDrivers = map[string]bool{
	"sqlite":   true,
	"postgres": true,
}

// loadEnvironment aplica as sobrescritas de variáveis de ambiente
// sobre a configuração padrão.
//
// As variáveis são lidas do ambiente do sistema. Caso um arquivo .env
// exista no diretório de trabalho, ele é carregado previamente pelo
// pacote godotenv, mas variáveis de ambiente reais sempre prevalecem.
//
// Variáveis suportadas:
//   - TASKMANAGER_DB_DRIVER: define o driver ("sqlite" ou "postgres")
//   - TASKMANAGER_DB_DSN: define a string de conexão
//   - TASKMANAGER_DB: alias legado para TASKMANAGER_DB_DSN
func loadEnvironment(cfg *Config) error {
	if driver, ok := os.LookupEnv(EnvDatabaseDriver); ok {
		driver = strings.TrimSpace(driver)
		if driver == "" {
			return errors.New("database driver cannot be empty")
		}
		cfg.Database.Driver = driver
	}

	dsn, dsnSet := os.LookupEnv(EnvDatabaseDSN)

	// Retrocompatibilidade: TASKMANAGER_DB é alias para TASKMANAGER_DB_DSN
	if !dsnSet {
		if legacy, ok := os.LookupEnv(EnvDatabasePath); ok {
			dsn = legacy
			dsnSet = true
		}
	}

	if dsnSet {
		if strings.TrimSpace(dsn) == "" {
			return errors.New("database DSN cannot be empty")
		}
		cfg.Database.DSN = dsn
	}

	return nil
}

// validateDriver verifica se o driver configurado é suportado.
func validateDriver(cfg *Config) error {
	if !validDrivers[cfg.Database.Driver] {
		return fmt.Errorf(
			"unsupported database driver %q: must be \"sqlite\" or \"postgres\"",
			cfg.Database.Driver,
		)
	}
	return nil
}
