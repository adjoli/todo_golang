package models

import "time"

type Task struct {
	ID        int64
	Title     string
	Completed bool
	CreatedAt time.Time
}
