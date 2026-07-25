// Taskmanager é o binário da interface de linha de comando (CLI)
// para gerenciamento de tarefas.
package main

import (
	"log"

	"github.com/adjoli/todo_golang/internal/app"
	"github.com/adjoli/todo_golang/internal/cli"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	if err := cli.Execute(app); err != nil {
		log.Fatal(err)
	}
}
