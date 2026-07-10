package repository

import (
	"context"
	"database/sql"

	"github.com/adjoli/todo_chatgpt/internal/models"
)

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

func (r *TaskRepository) FindByID(ctx context.Context, id int64) (*models.Task, error) {
	row := r.db.QueryRowContext(ctx, sqlFindTaskByID, id)

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

func (r *TaskRepository) List(ctx context.Context) ([]models.Task, error) {
	rows, err := r.db.QueryContext(ctx, sqlListTasks)
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

func (r *TaskRepository) Update(ctx context.Context, task *models.Task) error {
	result, err := r.db.ExecContext(
		ctx,
		sqlUpdateTask,
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

func (r *TaskRepository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, sqlDeleteTask, id)
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
