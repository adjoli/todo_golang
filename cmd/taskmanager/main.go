package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/adjoli/todo_chatgpt/internal/database"
	"github.com/adjoli/todo_chatgpt/internal/models"
	"github.com/adjoli/todo_chatgpt/internal/repository"
)

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.New(db)

	ctx := context.Background()

	task := newTask("Terminar de ler o livro de Go")

	if err := repo.Create(ctx, task); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task criada com sucesso!\n")
	fmt.Printf("ID: %d\n", task.ID)
}

func newTask(title string) *models.Task {
	return &models.Task{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
}
