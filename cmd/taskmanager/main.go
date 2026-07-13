package main

import (
	"log"

	"github.com/adjoli/todo_chatgpt/internal/database"
	"github.com/adjoli/todo_chatgpt/internal/repository"
	"github.com/adjoli/todo_chatgpt/internal/service"
)

func main() {
	db, err := database.Open(database.DefaultPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.New(db)

	service := service.New(repo)

	log.Println("Task Manager iniciado.")
	_ = service
}
