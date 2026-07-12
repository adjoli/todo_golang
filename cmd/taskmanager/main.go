package main

import (
	"log"

	"github.com/adjoli/todo_chatgpt/internal/database"
)

func main() {
	db, err := database.Open(database.DefaultPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Task Manager iniciado.")
}
