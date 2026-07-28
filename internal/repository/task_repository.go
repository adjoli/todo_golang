// Package repository implementa a persistência de tarefas via
// database/sql. Cada operação recebe context.Context para suportar
// cancelamento e timeouts.
package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/adjoli/todo_chatgpt/internal/models"
)

// TaskRepository é o repositório de persistência de tarefas.
// Ele encapsula as operações SQL e mapeia resultados para models.Task.
type TaskRepository struct {
	db      *sql.DB
	dialect dialect
	builder *sqlBuilder
}

// New cria um novo TaskRepository com a conexão de banco e o dialect
// fornecidos.
func New(db *sql.DB, d dialect) *TaskRepository {
	return &TaskRepository{
		db:      db,
		dialect: d,
		builder: newSQLBuilder(d),
	}
}

// Create insere uma nova tarefa no banco e popula o campo ID
// do objeto passado como referência.
func (r *TaskRepository) Create(
	ctx context.Context,
	task *models.Task,
) error {
	if r.dialect.supportsLastInsertID() {
		return r.createWithLastInsertID(ctx, task)
	}
	return r.createWithReturning(ctx, task)
}

// createWithLastInsertID insere a tarefa usando result.LastInsertId()
// (suportado por SQLite).
func (r *TaskRepository) createWithLastInsertID(
	ctx context.Context,
	task *models.Task,
) error {
	result, err := r.db.ExecContext(
		ctx,
		r.builder.insertTask(),
		task.Title,
		task.Completed,
		task.CreatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = id

	return nil
}

// createWithReturning insere a tarefa usando RETURNING id
// (necessário para Postgres).
func (r *TaskRepository) createWithReturning(
	ctx context.Context,
	task *models.Task,
) error {
	row := r.db.QueryRowContext(
		ctx,
		r.builder.insertTaskReturningID(),
		task.Title,
		task.Completed,
		task.CreatedAt,
	)

	var id int64
	if err := row.Scan(&id); err != nil {
		return fmt.Errorf("scan returning id: %w", err)
	}

	task.ID = id

	return nil
}

// FindByID busca uma tarefa pelo seu ID.
// Retorna sql.ErrNoRows se a tarefa não existir.
func (r *TaskRepository) FindByID(
	ctx context.Context,
	id int64,
) (*models.Task, error) {
	row := r.db.QueryRowContext(ctx, r.builder.findTaskByID(), id)

	task := &models.Task{}

	if err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.CreatedAt,
	); err != nil {
		return nil, err
	}

	return task, nil
}

// List retorna todas as tarefas que atendem ao filtro informado.
// O resultado é ordenado por ID crescente.
func (r *TaskRepository) List(
	ctx context.Context,
	filter models.TaskFilter,
) ([]models.Task, error) {
	query := r.builder.selectTasks()
	args := []any{}

	if filter.Completed != nil {
		query += r.builder.filterByCompleted()
		args = append(args, *filter.Completed)
	}

	query += r.builder.orderTasks()

	rows, err := r.db.QueryContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task

		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Completed,
			&task.CreatedAt,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Update atualiza o título e o status de uma tarefa existente.
// Retorna sql.ErrNoRows se o ID não existir no banco.
func (r *TaskRepository) Update(
	ctx context.Context,
	task *models.Task,
) error {
	result, err := r.db.ExecContext(
		ctx,
		r.builder.updateTask(),
		task.Title,
		task.Completed,
		task.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Delete remove uma tarefa pelo seu ID.
// Retorna sql.ErrNoRows se o ID não existir no banco.
func (r *TaskRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	result, err := r.db.ExecContext(ctx, r.builder.deleteTask(), id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
