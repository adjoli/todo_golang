package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/adjoli/todo_golang/internal/service"
	"github.com/spf13/cobra"
)

// newUpdateCommand cria o subcomando "update", que recebe um ID
// e um novo título como argumentos posicionais.
func (c *CLI) newUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <id> <title>",
		Short: "Update a task",
		Args:  cobra.ExactArgs(2),
		RunE:  c.runUpdate,
	}
	return cmd
}

// runUpdate atualiza o título de uma tarefa existente
// identificada pelo ID.
func (c *CLI) runUpdate(
	cmd *cobra.Command,
	args []string,
) error {
	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf(
			"invalid task ID %q: must be a number",
			args[0],
		)
	}

	if err := c.app.TaskService().UpdateTask(
		context.Background(),
		id,
		service.UpdateTaskInput{
			Title: args[1],
		},
	); err != nil {
		return handleError(cmd, err)
	}

	cmd.Printf("Task %d updated successfully.\n", id)

	return nil
}
