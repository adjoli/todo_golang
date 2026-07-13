package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/adjoli/todo_chatgpt/internal/models"
	"github.com/adjoli/todo_chatgpt/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

// New creates a new TaskService
func New(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// ----------------------------------------------
func (s *TaskService) CreateTask(
	ctx context.Context,
	title string,
) (*models.Task, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return nil, ErrEmptyTitle
	}

	task := &models.Task{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

// ----------------------------------------------
func (s *TaskService) ListTasks(
	ctx context.Context,
) ([]models.Task, error) {
	return s.repo.List(ctx)
}

// ----------------------------------------------
func (s *TaskService) CompleteTask(
	ctx context.Context,
	id int64,
) error {
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return mapRepositoryError(err)
	}

	if task.Completed {
		return ErrTaskAlreadyCompleted
	}

	task.Completed = true

	return s.repo.Update(ctx, task)
}

// ----------------------------------------------
func (s *TaskService) DeleteTask(
	ctx context.Context,
	id int64,
) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return mapRepositoryError(err)
	}
	return nil
}

// ----------------------------------------------
func mapRepositoryError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrTaskNotFound
	}
	return err
}
