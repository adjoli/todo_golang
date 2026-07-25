package repository

import (
	"testing"
	"time"

	"github.com/adjoli/todo_golang/internal/database"
	"github.com/adjoli/todo_golang/internal/models"
)

func newTestRepository(t *testing.T) *TaskRepository {
	t.Helper()

	db, err := database.New(":memory:")
	if err != nil {
		t.Fatalf("opening database: %v", err)
	}

	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Errorf("closing database: %v", err)
		}
	})

	return New(db)
}

func newTestTask() *models.Task {
	return &models.Task{
		Title:     "Estudar Go",
		Completed: false,
		CreatedAt: time.Now(),
	}
}
