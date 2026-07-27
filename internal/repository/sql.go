package repository

import "fmt"

// queries constrói as consultas SQL utilizando os placeholders
// do dialect configurado.
type queries struct {
	dialect dialect
}

// newQueries cria um queries com o dialect informado.
func newQueries(d dialect) *queries {
	return &queries{dialect: d}
}

func (q *queries) insertTask() string {
	return fmt.Sprintf(
		`INSERT INTO tasks (title, completed, created_at) VALUES (%s, %s, %s)`,
		q.dialect.placeholder(1),
		q.dialect.placeholder(2),
		q.dialect.placeholder(3),
	)
}

func (q *queries) insertTaskReturningID() string {
	return fmt.Sprintf(
		`INSERT INTO tasks (title, completed, created_at) VALUES (%s, %s, %s) RETURNING id`,
		q.dialect.placeholder(1),
		q.dialect.placeholder(2),
		q.dialect.placeholder(3),
	)
}

func (q *queries) findTaskByID() string {
	return fmt.Sprintf(
		`SELECT id, title, completed, created_at FROM tasks WHERE id = %s`,
		q.dialect.placeholder(1),
	)
}

func (q *queries) selectTasks() string {
	return `SELECT id, title, completed, created_at FROM tasks`
}

func (q *queries) orderTasks() string {
	return ` ORDER BY id`
}

func (q *queries) filterByCompleted() string {
	return fmt.Sprintf(` WHERE completed = %s`, q.dialect.placeholder(1))
}

func (q *queries) updateTask() string {
	return fmt.Sprintf(
		`UPDATE tasks SET title = %s, completed = %s WHERE id = %s`,
		q.dialect.placeholder(1),
		q.dialect.placeholder(2),
		q.dialect.placeholder(3),
	)
}

func (q *queries) deleteTask() string {
	return fmt.Sprintf(
		`DELETE FROM tasks WHERE id = %s`,
		q.dialect.placeholder(1),
	)
}
