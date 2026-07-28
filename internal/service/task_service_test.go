package service

import (
	"context"
	"errors"
	"testing"

	"github.com/adjoli/todo_golang/internal/models"
)

// TestCreateTask valida o comportamento de CreateTask em três cenários:
// título válido, título vazio e título com apenas espaços.
func TestCreateTask(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		wantError error
	}{
		{
			name:      "valid title",
			title:     "Estudar Go",
			wantError: nil,
		},
		{
			name:      "empty title",
			title:     "",
			wantError: ErrEmptyTitle,
		},
		{
			name:      "blank spaces",
			title:     "     ",
			wantError: ErrEmptyTitle,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(f *testing.T) {
			service := newTestService(t)

			task, err := service.CreateTask(
				context.Background(),
				CreateTaskInput{
					Title: tt.title,
				},
			)

			if !errors.Is(err, tt.wantError) {
				t.Fatalf(
					"expected error %v, got %v",
					tt.wantError,
					err,
				)
			}

			if tt.wantError != nil {
				return
			}

			if task.ID == 0 {
				t.Fatal("expected task ID to be populated")
			}

			if task.Title != tt.title {
				t.Errorf(
					"expected title %q, got %q",
					tt.title,
					task.Title,
				)
			}

			if task.Completed {
				t.Error("expected task to be created as not completed")
			}
		})
	}
}

// TestListTasks valida que ListTasks retorna a quantidade correta
// de tarefas, cobrindo lista vazia e lista com duas tarefas.
func TestListTasks(t *testing.T) {
	tests := []struct {
		name  string
		tasks []string
		want  int
	}{
		{
			name:  "empty list",
			tasks: nil,
			want:  0,
		},
		{
			name: "two tasks",
			tasks: []string{
				"Estudar Go",
				"Estudar SQL",
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := newTestService(t)
			ctx := context.Background()

			for _, title := range tt.tasks {
				_, err := service.CreateTask(
					ctx,
					CreateTaskInput{Title: title},
				)
				if err != nil {
					t.Fatalf("CreateTask(): %v", err)
				}
			}

			tasks, err := service.ListTasks(ctx, models.TaskFilter{})
			if err != nil {
				t.Fatalf("ListTasks(): %v", err)
			}

			if len(tasks) != tt.want {
				t.Fatalf(
					"expected %d tasks, got %d",
					tt.want,
					len(tasks),
				)
			}
		})
	}
}

// TestCompleteTask verifica que CompleteTask marca a tarefa
// como concluída.
func TestCompleteTask(t *testing.T) {
	service := newTestService(t)
	ctx := context.Background()

	task, err := service.CreateTask(
		ctx,
		CreateTaskInput{
			Title: "Estudar Go",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := service.CompleteTask(ctx, task.ID); err != nil {
		t.Fatalf("CompleteTask(): %v", err)
	}

	tasks, err := service.ListTasks(ctx, models.TaskFilter{})
	if err != nil {
		t.Fatal(err)
	}

	if !tasks[0].Completed {
		t.Fatal("expected task to be completed")
	}
}

// TestCompleteTask_NotFound verifica que CompleteTask retorna
// ErrTaskNotFound quando o ID não existe.
func TestCompleteTask_NotFound(t *testing.T) {
	service := newTestService(t)

	err := service.CompleteTask(
		context.Background(),
		999,
	)

	if !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf(
			"expected ErrTaskNotFound, got %v",
			err,
		)
	}
}

// TestCompleteTask_AlreadyCompleted verifica que CompleteTask retorna
// ErrTaskAlreadyCompleted quando a tarefa já está concluída.
func TestCompleteTask_AlreadyCompleted(t *testing.T) {
	service := newTestService(t)
	ctx := context.Background()

	task, _ := service.CreateTask(
		ctx,
		CreateTaskInput{
			Title: "Estudar Go",
		},
	)

	if err := service.CompleteTask(ctx, task.ID); err != nil {
		t.Fatal(err)
	}

	err := service.CompleteTask(ctx, task.ID)

	if !errors.Is(err, ErrTaskAlreadyCompleted) {
		t.Fatalf("expected ErrTaskAlreadyCompleted, got %v", err)
	}
}

// TestDeleteTask verifica que DeleteTask remove a tarefa e que
// a lista fica vazia após a remoção.
func TestDeleteTask(t *testing.T) {
	service := newTestService(t)
	ctx := context.Background()

	task, _ := service.CreateTask(
		ctx,
		CreateTaskInput{
			Title: "Estudar Go",
		},
	)

	if err := service.DeleteTask(ctx, task.ID); err != nil {
		t.Fatal(err)
	}

	tasks, err := service.ListTasks(ctx, models.TaskFilter{})
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 0 {
		t.Fatal("expected empty list")
	}
}

// TestDeleteTask_NotFound verifica que DeleteTask retorna
// ErrTaskNotFound quando o ID não existe.
func TestDeleteTask_NotFound(t *testing.T) {
	service := newTestService(t)

	err := service.DeleteTask(
		context.Background(),
		999,
	)

	if !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("expected ErrTaskNotFound, got %v", err)
	}
}

// TestUpdateTask verifica que UpdateTask altera o título
// de uma tarefa existente.
func TestUpdateTask(t *testing.T) {
	service := newTestService(t)
	ctx := context.Background()

	task, err := service.CreateTask(
		ctx, CreateTaskInput{
			Title: "Estudar Go",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := service.UpdateTask(
		ctx,
		task.ID,
		UpdateTaskInput{
			Title: "Estudar Go profundamente",
		},
	); err != nil {
		t.Fatalf("UpdateTask(): %v", err)
	}

	tasks, err := service.ListTasks(ctx, models.TaskFilter{})
	if err != nil {
		t.Fatal(err)
	}

	if tasks[0].Title != "Estudar Go profundamente" {
		t.Fatalf("expected updated title, got %q", tasks[0].Title)
	}
}
