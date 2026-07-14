package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *CLI) newDoneCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "done <id>",
		Short: "Mark a task as completed",
		Args:  cobra.ExactArgs(1),
		RunE:  c.runDone,
	}
	return cmd
}

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
