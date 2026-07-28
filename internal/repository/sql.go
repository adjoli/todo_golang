package repository

import "fmt"

// sqlBuilder constrói as consultas SQL utilizando os placeholders
// do dialect configurado.
type sqlBuilder struct {
	dialect dialect
}

// newSQLBuilder cria um sqlBuilder com o dialect informado.
func newSQLBuilder(d dialect) *sqlBuilder {
	return &sqlBuilder{dialect: d}
}

func (b *sqlBuilder) insertTask() string {
	return fmt.Sprintf(
		`INSERT INTO tasks (title, completed, created_at) VALUES (%s, %s, %s)`,
		b.dialect.placeholder(1),
		b.dialect.placeholder(2),
		b.dialect.placeholder(3),
	)
}

func (b *sqlBuilder) insertTaskReturningID() string {
	return fmt.Sprintf(
		`INSERT INTO tasks (title, completed, created_at) VALUES (%s, %s, %s) RETURNING id`,
		b.dialect.placeholder(1),
		b.dialect.placeholder(2),
		b.dialect.placeholder(3),
	)
}

func (b *sqlBuilder) findTaskByID() string {
	return fmt.Sprintf(
		`SELECT id, title, completed, created_at FROM tasks WHERE id = %s`,
		b.dialect.placeholder(1),
	)
}

func (b *sqlBuilder) selectTasks() string {
	return `SELECT id, title, completed, created_at FROM tasks`
}

func (b *sqlBuilder) orderTasks() string {
	return ` ORDER BY id`
}

func (b *sqlBuilder) filterByCompleted() string {
	return fmt.Sprintf(` WHERE completed = %s`, b.dialect.placeholder(1))
}

func (b *sqlBuilder) updateTask() string {
	return fmt.Sprintf(
		`UPDATE tasks SET title = %s, completed = %s WHERE id = %s`,
		b.dialect.placeholder(1),
		b.dialect.placeholder(2),
		b.dialect.placeholder(3),
	)
}

func (b *sqlBuilder) deleteTask() string {
	return fmt.Sprintf(
		`DELETE FROM tasks WHERE id = %s`,
		b.dialect.placeholder(1),
	)
}
