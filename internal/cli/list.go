package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func (c *CLI) newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		RunE:  c.runList,
	}
	return cmd
}

func (c *CLI) runList(
	cmd *cobra.Command,
	args []string,
) error {
	tasks, err := c.app.TaskService().ListTasks(
		context.Background(),
	)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	cmd.Printf("%-4s %-10s %s\n", "ID", "STATUS", "TITLE")
	cmd.Printf("%-4s %-10s %s\n", "--", "------", "-----")

	for _, task := range tasks {
		fmt.Printf(
			"%-4d %-10s %s\n",
			task.ID,
			statusLabel(task.Completed),
			task.Title,
		)
	}
	return nil
}

// ----------------------------------------------
func statusLabel(completed bool) string {
	if completed {
		return "🟢"
	}
	return "🔴"
}