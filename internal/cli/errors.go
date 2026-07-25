package cli

import (
	"errors"
	"fmt"

	"github.com/adjoli/todo_chatgpt/internal/service"
	"github.com/spf13/cobra"
)

// handleError mapeia erros conhecidos do service para mensagens
// amigáveis exibidas ao usuário. Erros inesperados são embrulhados
// com contexto e retornados ao caller.
func handleError(
	cmd *cobra.Command,
	err error,
) error {
	switch {
	case errors.Is(err, service.ErrTaskNotFound):
		cmd.Println("Task not found.")
		return nil

	case errors.Is(err, service.ErrTaskAlreadyCompleted):
		cmd.Println("Task is already completed.")
		return nil

	case errors.Is(err, service.ErrEmptyTitle):
		cmd.Println("Title cannot be empty.")
		return nil
	}

	return fmt.Errorf("unexpected error: %w", err)
}
