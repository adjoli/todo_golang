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

type App struct {
	cfg         *config.Config
	db          *sql.DB
	logger      *slog.Logger
	taskService *service.TaskService
}

func (a *App) TaskService() *service.TaskService {
	return a.taskService
}

func (a *App) Close() error {
	if a.db == nil {
		return nil
	}
	return a.db.Close()
}

func (a *App) Config() *config.Config {
	return a.cfg
}

func New() (*App, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("initialize application: %w", err)
	}

	db, err := database.New(cfg.Database.Path)
	if err != nil {
		return nil, fmt.Errorf("initialize application: %w", err)
	}

	logger := logger.New()
	repo := repository.New(db)
	service := service.New(repo, logger)

	return &App{
		db:          db,
		taskService: service,
	}, nil
}
