// Package service implementa as regras de negócio de gerenciamento
// de tarefas. Ele orquestra operações entre a camada de interface
// (CLI/API) e a camada de persistência (repository), validando
// dados e traduzindo erros de infraestrutura em erros de domínio.
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

// TaskService é o service que orquestra as operações de tarefas.
// Ele valida regras de negócio, delega persistência ao repository
// e registra eventos relevantes no logger.
type TaskService struct {
	repo   *repository.TaskRepository
	logger *slog.Logger
}

// New cria um TaskService com o repository e o logger fornecidos.
func New(
	repo *repository.TaskRepository,
	logger *slog.Logger,
) *TaskService {
	return &TaskService{
		repo:   repo,
		logger: logger,
	}
}

// CreateTask valida o título, cria uma nova tarefa com status
// pendente e a persiste no banco. Retorna a tarefa criada com
// o ID populado.
func (s *TaskService) CreateTask(
	ctx context.Context,
	input CreateTaskInput,
) (*models.Task, error) {
	title, err := validateTitle(input.Title)
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

// ListTasks retorna todas as tarefas que atendem ao filtro informado,
// delegando a consulta diretamente ao repository.
func (s *TaskService) ListTasks(
	ctx context.Context,
	filter models.TaskFilter,
) ([]models.Task, error) {
	return s.repo.List(ctx, filter)
}

// GetTask busca uma tarefa pelo seu ID. Retorna ErrTaskNotFound
// se a tarefa não existir.
func (s *TaskService) GetTask(
	ctx context.Context,
	id int64,
) (models.Task, error) {
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		err = mapRepositoryError(err)
		if errors.Is(err, ErrTaskNotFound) {
			s.logger.Warn(
				"task not found",
				slog.Int64("id", id),
			)
		}
		return models.Task{}, err
	}

	return *task, nil
}

// CompleteTask marca uma tarefa como concluída pelo seu ID.
// Retorna ErrTaskNotFound se a tarefa não existir ou
// ErrTaskAlreadyCompleted se já estiver concluída.
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

// DeleteTask remove uma tarefa permanentemente pelo seu ID.
// Retorna ErrTaskNotFound se a tarefa não existir.
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

// UpdateTask atualiza o título de uma tarefa existente.
// Retorna ErrTaskNotFound se a tarefa não existir ou ErrEmptyTitle
// se o novo título for vazio. Se o título não mudar, a operação
// é um no-op.
func (s *TaskService) UpdateTask(
	ctx context.Context,
	id int64,
	input UpdateTaskInput,
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

	title, err := validateTitle(input.Title)
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

// mapRepositoryError converte erros de infraestrutura do repository
// em erros de domínio do service. Atualmente, sql.ErrNoRows é
// mapeado para ErrTaskNotFound.
func mapRepositoryError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrTaskNotFound
	}
	return err
}

// validateTitle remove espaços em branco das extremidades e valida
// que o título não está vazio. Retorna ErrEmptyTitle se o resultado
// for uma string vazia.
func validateTitle(title string) (string, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return "", ErrEmptyTitle
	}

	return title, nil
}
