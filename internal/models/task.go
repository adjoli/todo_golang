// Package models define os tipos de dados do domínio da aplicação.
package models

import "time"

// Task representa uma tarefa gerenciada pela aplicação.
type Task struct {
	ID        int64
	Title     string
	Completed bool
	CreatedAt time.Time
}
