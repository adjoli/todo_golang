package service

import (
	"testing"

	"github.com/adjoli/todo_chatgpt/internal/database"
	"github.com/adjoli/todo_chatgpt/internal/repository"
)

func newTestService(t *testing.T) *TaskService {
	t.Helper()

	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("opening database: %v", err)
	}

	t.Cleanup(func() {
		_ = db.Close()
	})

	repo := repository.New(db)

	return New(repo)
}
