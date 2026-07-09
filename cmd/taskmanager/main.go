package main

import (
	"fmt"
	"log"

	"github.com/adjoli/todo_chatgpt/internal/database"
)

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Task Manager iniciado com sucesso!")
}
