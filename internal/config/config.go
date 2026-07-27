// Package config carrega e valida as configurações da aplicação.
// As configurações são definidas por variáveis de ambiente, com
// valores padrão sensatos quando a variável não está presente.
package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

const (
	// EnvDatabaseDriver é o nome da variável de ambiente que define
	// o driver de banco de dados. Valores aceitos: "sqlite" (default)
	// e "postgres".
	EnvDatabaseDriver = "TASKMANAGER_DB_DRIVER"

	// EnvDatabaseDSN é o nome da variável de ambiente que define
	// a string de conexão do banco de dados. Para SQLite, é o caminho
	// do arquivo. Para Postgres, é a connection string DSN.
	EnvDatabaseDSN = "TASKMANAGER_DB_DSN"

	// EnvDatabasePath é a variável legada mantida por retrocompatibilidade.
	// Equivale a TASKMANAGER_DB_DSN.
	EnvDatabasePath = "TASKMANAGER_DB"

	// DefaultDriver é o driver padrão quando TASKMANAGER_DB_DRIVER
	// não está definido.
	DefaultDriver = "sqlite"

	// DefaultDSN é a string de conexão padrão para SQLite.
	DefaultDSN = "data/tasks.db"
)

// Config armazena as configurações carregadas da aplicação.
type Config struct {
	Database DatabaseConfig
}

// DatabaseConfig armazena as configurações de conexão com o banco de dados.
type DatabaseConfig struct {
	Driver string // "sqlite" ou "postgres"
	DSN    string // file path (sqlite) ou connection string (postgres)
}

// New cria uma nova Config aplicando os valores padrão, carregando
// sobrescritas de variáveis de ambiente e preparando o diretório
// do banco de dados quando o driver é SQLite.
//
// Se um arquivo .env existir no diretório de trabalho, as variáveis
// definidas nele são carregadas antes da leitura do ambiente.
// Variáveis de ambiente reais SEMPRE SOBREPÕEM valores do arquivo .env.
func New() (*Config, error) {
	cfg := defaultConfig()

	_ = godotenv.Load()

	if err := loadEnvironment(cfg); err != nil {
		return nil, err
	}

	if err := validateDriver(cfg); err != nil {
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
			Driver: DefaultDriver,
			DSN:    DefaultDSN,
		},
	}
}
