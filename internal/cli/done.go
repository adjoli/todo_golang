package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// newDoneCommand cria o subcomando "done", que recebe um ID
// como argumento positional e marca a tarefa como concluída.
func (c *CLI) newDoneCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "done <id>",
		Short: "Mark a task as completed",
		Args:  cobra.ExactArgs(1),
		RunE:  c.runDone,
	}
	return cmd
}

// runDone executa a marcação de uma tarefa como concluída
// pelo seu ID.
func (c *CLI) runDone(
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
	err = c.app.TaskService().CompleteTask(
		context.Background(),
		id,
	)
	if err != nil {
		return handleError(cmd, err)
	}

	cmd.Printf("Task %d completed successfully.\n", id)

	return nil
}
