// Package config carrega e valida as configurações da aplicação.
// As configurações são definidas por variáveis de ambiente, com
// valores padrão sensatos quando a variável não está presente.
package config

import "fmt"

// EnvDatabasePath é o nome da variável de ambiente que sobrescreve
// o caminho padrão do banco de dados SQLite.
const EnvDatabasePath = "TASKMANAGER_DB"

// DefaultDatabasePath é o caminho padrão do banco de dados SQLite
// quando a variável de ambiente não está definida.
const DefaultDatabasePath = "data/tasks.db"

// Config armazena as configurações carregadas da aplicação.
type Config struct {
	Database DatabaseConfig
}

// DatabaseConfig armazena as configurações de conexão com o banco de dados.
type DatabaseConfig struct {
	Path string
}

// New cria uma nova Config aplicando os valores padrão, carregando
// sobrescritas de variáveis de ambiente e preparando o diretório
// do banco de dados.
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

// defaultConfig retorna uma Config com os valores padrão.
func defaultConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Path: DefaultDatabasePath,
		},
	}
}
