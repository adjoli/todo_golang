package service

import (
	"testing"

	"github.com/adjoli/todo_golang/internal/database"
	"github.com/adjoli/todo_golang/internal/logger"
	"github.com/adjoli/todo_golang/internal/repository"
)

func newTestService(t *testing.T) *TaskService {
	t.Helper()

	db, err := database.New("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("opening database: %v", err)
	}

	t.Cleanup(func() {
		_ = db.Close()
	})

	d, _ := repository.NewDialect("sqlite")

	logger := logger.New()
	repo := repository.New(db, d)

	return New(repo, logger)
}
