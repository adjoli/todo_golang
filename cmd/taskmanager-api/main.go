package main

import (
	"log"

	"github.com/adjoli/todo_chatgpt/internal/api"
	"github.com/adjoli/todo_chatgpt/internal/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	server := api.New(app)

	log.Fatal(server.Start(":8080"))
}
