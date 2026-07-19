package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *CLI) newUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <id> <title>",
		Short: "Update a task",
		Args:  cobra.ExactArgs(2),
		RunE:  c.runUpdate,
	}
	return cmd
}

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
		args[1],
	); err != nil {
		return handleError(cmd, err)
	}

	cmd.Printf("Task %d updated successfully.\n", id)

	return nil
}
