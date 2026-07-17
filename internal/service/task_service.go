package service

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"strings"
	"time"

	"github.com/adjoli/todo_chatgpt/internal/models"
	"github.com/adjoli/todo_chatgpt/internal/repository"
)

type TaskService struct {
	repo   *repository.TaskRepository
	logger *slog.Logger
}

// New creates a new TaskService
func New(
	repo *repository.TaskRepository,
	logger *slog.Logger,
) *TaskService {
	return &TaskService{
		repo:   repo,
		logger: logger,
	}
}

// ----------------------------------------------
func (s *TaskService) CreateTask(
	ctx context.Context,
	title string,
) (*models.Task, error) {
	title, err := validateTitle(title)
	if err != nil {
		return nil, err
	}

	task := &models.Task{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}

	s.logger.Info(
		"task created",
		slog.Int64("id", task.ID),
		slog.String("title", task.Title),
	)
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
		err = mapRepositoryError(err)
		if errors.Is(err, ErrTaskNotFound) {
			s.logger.Warn(
				"task not found",
				slog.Int64("id", id),
			)
		}
		return err
	}

	if task.Completed {
		s.logger.Warn(
			"task already completed",
			slog.Int64("id", task.ID),
			slog.String("title", task.Title),
		)
		return ErrTaskAlreadyCompleted
	}

	task.Completed = true

	if err := s.repo.Update(ctx, task); err != nil {
		return err
	}

	s.logger.Info(
		"task completed",
		slog.Int64("id", task.ID),
		slog.String("title", task.Title),
	)

	return nil
}

// ----------------------------------------------
func (s *TaskService) DeleteTask(
	ctx context.Context,
	id int64,
) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		err = mapRepositoryError(err)
		if errors.Is(err, ErrTaskNotFound) {
			s.logger.Warn(
				"task not found",
				slog.Int64("id", id),
			)
		}
		return err
	}

	s.logger.Info(
		"task removed",
		slog.Int64("id", id),
	)

	return nil
}

// ----------------------------------------------
func (s *TaskService) UpdateTask(
	ctx context.Context,
	id int64,
	title string,
) error {
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		err = mapRepositoryError(err)
		if errors.Is(err, ErrTaskNotFound) {
			s.logger.Warn(
				"task not found",
				slog.Int64("id", id),
			)
		}
		return err
	}

	title, err = validateTitle(title)
	if err != nil {
		return err
	}

	oldTitle := task.Title
	
	if oldTitle == title {
		return nil
	}

	task.Title = title

	if err := s.repo.Update(ctx, task); err != nil {
		return err
	}

	s.logger.Info(
		"task update",
		slog.Int64("id", task.ID),
		slog.String("old_title", oldTitle),
		slog.String("new_title", task.Title),
	)

	return nil
}

// ----------------------------------------------
// HELPER FUNCTIONS
// ----------------------------------------------
func mapRepositoryError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrTaskNotFound
	}
	return err
}

func validateTitle(title string) (string, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return "", ErrEmptyTitle
	}

	return title, nil
}
