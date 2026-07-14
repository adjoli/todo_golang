package app

import (
	"database/sql"

	"github.com/adjoli/todo_chatgpt/internal/database"
	"github.com/adjoli/todo_chatgpt/internal/repository"
	"github.com/adjoli/todo_chatgpt/internal/service"
)

type App struct {
	db          *sql.DB
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

func New() (*App, error) {
	db, err := database.Open(database.DefaultPath)
	if err != nil {
		return nil, err
	}

	repo := repository.New(db)
	taskService := service.New(repo)

	return &App{
		db:          db,
		taskService: taskService,
	}, nil
}
