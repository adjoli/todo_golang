package cli

import (
	"context"
	"fmt"

	"github.com/adjoli/todo_chatgpt/internal/models"
	"github.com/spf13/cobra"
)

func (c *CLI) newListCommand() *cobra.Command {
	var all bool

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.runList(cmd, args, all)

		},
	}

	cmd.Flags().BoolVar(
		&all,
		"all",
		false,
		"List completed and pending tasks",
	)

	return cmd
}

func (c *CLI) runList(
	cmd *cobra.Command,
	args []string,
	all bool,
) error {
	filter := models.TaskFilter{}

	if !all {
		completed := false
		filter.Completed = &completed
	}

	tasks, err := c.app.TaskService().ListTasks(
		context.Background(),
		filter,
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
