package cli

import (
	"context"

	"github.com/adjoli/todo_chatgpt/internal/service"
	"github.com/spf13/cobra"
)

// newAddCommand cria o subcomando "add", que recebe um título
// como argumento positional e cria uma nova tarefa.
func (c *CLI) newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <title>",
		Short: "Add a new task",
		Args:  cobra.ExactArgs(1),
		RunE:  c.runAdd,
	}
	return cmd
}

// runAdd executa a criação de uma tarefa a partir do título
// informado no argumento positional.
func (c *CLI) runAdd(
	cmd *cobra.Command,
	args []string,
) error {
	task, err := c.app.TaskService().CreateTask(
		context.Background(),
		service.CreateTaskInput{
			Title: args[0],
		},
	)
	if err != nil {
		return handleError(cmd, err)
	}

	cmd.Printf(
		"Task %d created successfully.\n",
		task.ID,
	)
	return nil
}
