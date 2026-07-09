package repository

import (
	"context"
	"database/sql"

	"github.com/adjoli/todo_chatgpt/internal/models"
)

const sqlInsertTask = `
INSERT INTO tasks (
	title,
	completed,
	created_at
)
VALUES (?, ?, ?);
`

type TaskRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(ctx context.Context, task *models.Task) error {
	result, err := r.db.ExecContext(
		ctx,
		sqlInsertTask,
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
