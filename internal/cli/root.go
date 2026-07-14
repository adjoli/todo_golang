package cli

import (
	"github.com/adjoli/todo_chatgpt/internal/app"
	"github.com/spf13/cobra"
)

type CLI struct {
	app  *app.App
	root *cobra.Command
}

func Execute(app *app.App) error {
	cli := &CLI{
		app: app,
	}

	cli.root = cli.newRootCommand()

	return cli.root.Execute()
}

func (c *CLI) newRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "taskmanager",
		Short: "Manage your tasks from the command line.",
		Long: `Task Manager is a command-line application built in Go.

It demonstrates idiomatic Go practices while providing
a simple and useful task management tool.`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(
		c.newListCommand(),
		c.newAddCommand(),
		c.newDoneCommand(),
		c.newRemoveCommand(),
	)

	return rootCmd
}
