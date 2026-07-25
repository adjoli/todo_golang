package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// newRemoveCommand cria o subcomando "remove" (alias "rm"),
// que recebe um ID como argumento positional e remove a tarefa
// permanentemente.
func (c *CLI) newRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "remove <id>",
		Aliases: []string{"rm"},
		Short:   "Remove a task",
		Long: `Remove a task permanently.

This operation cannot be undone.`,
		Example: `  taskmanager remove 1
taskmanager rm 1`,
		Args: cobra.ExactArgs(1),
		RunE: c.runRemove,
	}
}

// runRemove remove permanentemente uma tarefa pelo seu ID.
func (c *CLI) runRemove(
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

	if err := c.app.TaskService().DeleteTask(
		context.Background(),
		id,
	); err != nil {
		return handleError(cmd, err)
	}

	cmd.Printf(
		"Task %d removed sucessfully.\n",
		id,
	)

	return nil
}
