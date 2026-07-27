// Package app é o composition root da aplicação.
// Ele inicializa e conecta todas as camadas (config, banco de dados,
// logger, repository e service), expondo a estrutura App como
// container de dependências e ponto de lifecycle (Close).
package app

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/adjoli/todo_chatgpt/internal/config"
	"github.com/adjoli/todo_chatgpt/internal/database"
	"github.com/adjoli/todo_chatgpt/internal/logger"
	"github.com/adjoli/todo_chatgpt/internal/repository"
	"github.com/adjoli/todo_chatgpt/internal/service"
)

// App é o container de dependências e lifecycle da aplicação.
// Ele mantém as referências para a configuração, conexão com o banco,
// logger e serviços. A instância deve ser criada com New e
// liberada com Close quando não for mais necessária.
type App struct {
	cfg         *config.Config
	db          *sql.DB
	logger      *slog.Logger
	taskService *service.TaskService
}

// TaskService retorna o service de gerenciamento de tarefas.
func (a *App) TaskService() *service.TaskService {
	return a.taskService
}

// Close fecha a conexão com o banco de dados.
// É nil-safe: retorna nil imediatamente se o banco não foi inicializado.
func (a *App) Close() error {
	if a.db == nil {
		return nil
	}
	return a.db.Close()
}

// Config retorna a configuração carregada da aplicação.
func (a *App) Config() *config.Config {
	return a.cfg
}

// Logger retorna o logger estruturado da aplicação.
func (a *App) Logger() *slog.Logger {
	return a.logger
}

// New inicializa a aplicação criando e conectando todas as camadas
// na ordem: configuração, banco de dados, logger, repository e service.
// A função também cria o schema do banco de dados automaticamente.
// O caller deve chamar Close quando a instância não for mais necessária.
func New() (*App, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("initialize application: %w", err)
	}

	db, err := database.New(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		return nil, fmt.Errorf("initialize application: %w", err)
	}

	d, err := repository.NewDialect(cfg.Database.Driver)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("initialize application: %w", err)
	}

	logger := logger.New()
	repo := repository.New(db, d)
	svc := service.New(repo, logger)

	return &App{
		cfg:         cfg,
		db:          db,
		logger:      logger,
		taskService: svc,
	}, nil
}
