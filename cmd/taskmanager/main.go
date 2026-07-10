package main

import (
	"context"
	"database/sql"
	"errors"
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

	task := testCreate(ctx, repo, "Fazer uma tarefa aleatória")

	testFindByID(ctx, repo, task.ID)

	testFindByIDNotFound(ctx, repo)

	testList(ctx, repo)

	testUpdate(ctx, repo, task.ID)

	testDelete(ctx, repo, task.ID)
}

func newTask(title string) *models.Task {
	return &models.Task{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
}

func printTask(task *models.Task) {
	fmt.Printf("ID.........: %d\n", task.ID)
	fmt.Printf("Título.....: %s\n", task.Title)
	fmt.Printf("Concluída..: %t\n", task.Completed)
	fmt.Printf("Criada em..: %s\n", task.CreatedAt.Format(time.RFC3339))
}

func testCreate(ctx context.Context, repo *repository.TaskRepository, title string) *models.Task {
	fmt.Println("== CREATE ==")

	task := newTask(title)

	if err := repo.Create(ctx, task); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task criada com sucesso! [ID: %d]\n", task.ID)

	return task
}

func testFindByID(ctx context.Context, repo *repository.TaskRepository, id int64) {
	fmt.Println("== FIND BY ID ==")

	task, err := repo.FindByID(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	printTask(task)
}

func testFindByIDNotFound(ctx context.Context, repo *repository.TaskRepository) {
	fmt.Println("== FIND BY ID (NOT FOUND) ==")

	_, err := repo.FindByID(ctx, 9999)
	if err != nil {
		fmt.Printf("Erro retornado: %v\n\n", err)
		return
	}

	log.Fatal("esperava um erro, mas recebeu nil")
}

func testList(ctx context.Context, repo *repository.TaskRepository) {
	fmt.Println("== LIST ==")

	tasks, err := repo.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for i := range tasks {
		printTask(&tasks[i])
		fmt.Println()
	}
}

func testUpdate(ctx context.Context, repo *repository.TaskRepository, id int64) {
	fmt.Println("== UPDATE ==")

	task, err := repo.FindByID(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	task.Title += " [modificada]"
	task.Completed = true

	if err := repo.Update(ctx, task); err != nil {
		log.Fatal(err)
	}

	task, err = repo.FindByID(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	printTask(task)
	fmt.Println()
}

func testDelete(ctx context.Context, repo *repository.TaskRepository, id int64) {
	fmt.Println("== DELETE ==")

	if err := repo.Delete(ctx, id); err != nil {
		log.Fatal(err)
	}

	_, err := repo.FindByID(ctx, id)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Task removida com sucesso.")
		fmt.Println()
		return
	}

	log.Fatal("a tarefa deveria ter sido removida")
}
