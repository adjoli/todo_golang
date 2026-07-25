// Package cli define os comandos da interface de linha de comando
// da aplicação Task Manager, utilizando Cobra como framework de CLI.
package cli

import (
	"github.com/adjoli/todo_chatgpt/internal/app"
	"github.com/spf13/cobra"
)

// CLI é o container que conecta a aplicação (app.App) à interface
// de linha de comando. Ele encapsula a instância da aplicação e
// o comando raiz do Cobra.
type CLI struct {
	app  *app.App
	root *cobra.Command
}

// Execute é o ponto de entrada da CLI. Ele configura o comando raiz,
// registra os subcomandos e executa a CLI.
func Execute(app *app.App) error {
	cli := &CLI{
		app: app,
	}

	cli.root = cli.newRootCommand()

	return cli.root.Execute()
}

// newRootCommand cria o comando raiz "taskmanager" e registra
// todos os subcomandos disponíveis.
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
		c.newUpdateCommand(),
	)

	return rootCmd
}
