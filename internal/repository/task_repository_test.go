package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/adjoli/todo_golang/internal/models"
)

// ----------------------------------------------
func TestCreate(t *testing.T) {
	repo := newTestRepository(t)

	task := newTestTask()

	err := repo.Create(context.Background(), task)
	if err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	if task.ID == 0 {
		t.Fatalf("expected ID to be populated")
	}
}

// ----------------------------------------------
func TestFindByID(t *testing.T) {
	repo := newTestRepository(t)
	ctx := context.Background()

	// Arrange
	expected := newTestTask()

	if err := repo.Create(ctx, expected); err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	// Act
	actual, err := repo.FindByID(ctx, expected.ID)
	// Assert
	if err != nil {
		t.Fatalf("FindByID() returned error: %v", err)
	}

	if actual.ID != expected.ID {
		t.Errorf("expected ID %d, got %d", expected.ID, actual.ID)
	}

	if actual.Title != expected.Title {
		t.Errorf("expected Title %q, got %q", expected.Title, actual.Title)
	}

	if actual.Completed != expected.Completed {
		t.Errorf("expected Completed %t, got %t", expected.Completed, actual.Completed)
	}

	if !actual.CreatedAt.Equal(expected.CreatedAt) {
		t.Errorf("expected CreatedAt %v, got %v", expected.CreatedAt, actual.CreatedAt)
	}
}

// ----------------------------------------------
func TestList_AllTasks(t *testing.T) {
	repo := newTestRepository(t)
	ctx := context.Background()

	// Arrange
	first := newTestTask()
	first.Title = "Primeira tarefa"

	second := newTestTask()
	second.Title = "Segunda tarefa"

	if err := repo.Create(ctx, first); err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	if err := repo.Create(ctx, second); err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	// Act
	tasks, err := repo.List(ctx, models.TaskFilter{})
	// Assert
	if err != nil {
		t.Fatalf("List() returned error: %v", err)
	}

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].ID != first.ID {
		t.Errorf("expected first ID %d, got %d", first.ID, tasks[0].ID)
	}

	if tasks[1].ID != second.ID {
		t.Errorf("expected second ID %d, got %d", second.ID, tasks[1].ID)
	}
}

// ----------------------------------------------
func TestList_OnlyPendingTasks(t *testing.T) {
	repo := newTestRepository(t)
	ctx := context.Background()
	myTasks := make([]*models.Task, 3)

	// Arrange
	myTasks[0] = newTestTask()
	myTasks[0].Title = "Primeira tarefa"

	myTasks[1] = newTestTask()
	myTasks[1].Title = "Segunda tarefa"
	myTasks[1].Completed = true

	myTasks[2] = newTestTask()
	myTasks[2].Title = "Terceira tarefa"

	for i, task := range myTasks {
		if err := repo.Create(ctx, task); err != nil {
			t.Fatalf("#%d: Create() returned error: %v", i, err)
		}
	}

	completed := false

	// Act
	tasks, err := repo.List(ctx, models.TaskFilter{Completed: &completed})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Assert
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	for _, task := range tasks {
		if task.Completed {
			t.Fatalf("expected only pending tasks")
		}
	}
}

// ----------------------------------------------
func TestList_OnlyCompleted(t *testing.T) {
	repo := newTestRepository(t)
	ctx := context.Background()
	myTasks := make([]*models.Task, 3)

	// Arrange
	myTasks[0] = newTestTask()
	myTasks[0].Title = "Primeira tarefa"

	myTasks[1] = newTestTask()
	myTasks[1].Title = "Segunda tarefa"
	myTasks[1].Completed = true

	myTasks[2] = newTestTask()
	myTasks[2].Title = "Terceira tarefa"

	for i, task := range myTasks {
		if err := repo.Create(ctx, task); err != nil {
			t.Fatalf("#%d: Create() returned error: %v", i, err)
		}
	}

	completed := true

	// Act
	tasks, err := repo.List(ctx, models.TaskFilter{Completed: &completed})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Assert
	if len(tasks) != 1 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	for _, task := range tasks {
		if !task.Completed {
			t.Fatalf("expected only completed tasks")
		}
	}
}

// ----------------------------------------------
func TestUpdate(t *testing.T) {
	repo := newTestRepository(t)
	ctx := context.Background()

	// Arrange
	task := newTestTask()

	if err := repo.Create(ctx, task); err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	task.Title = "Titulo atualizado"
	task.Completed = true

	// Act
	if err := repo.Update(ctx, task); err != nil {
		t.Fatalf("Update() returned error: %v", err)
	}

	updated, err := repo.FindByID(ctx, task.ID)
	// Assert
	if err != nil {
		t.Fatalf("FindByID() returned error: %v", err)
	}

	if updated.Title != task.Title {
		t.Errorf("expected title %q, got %q", task.Title, updated.Title)
	}

	if updated.Completed != task.Completed {
		t.Errorf("expected completed %t, got %t", task.Completed, updated.Completed)
	}
}

// ----------------------------------------------
func TestDelete(t *testing.T) {
	repo := newTestRepository(t)
	ctx := context.Background()

	// Arrange
	task := newTestTask()

	if err := repo.Create(ctx, task); err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	// Act
	if err := repo.Delete(ctx, task.ID); err != nil {
		t.Fatalf("Delete() returned error: %v", err)
	}

	_, err := repo.FindByID(ctx, task.ID)

	// Assert
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("expected sql.ErrNoRows, got %v", err)
	}
}
