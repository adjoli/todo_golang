package logger

import (
	"log/slog"
	"os"
)

// New cria o logger padrão da aplicação.
//
// Inicialmente utilizamos um TextHandler escrevendo em stdout.
// No futuro esta implementação poderá evoluir para suportar:
//
//   - JSON
//   - níveis de log
//   - arquivos
//   - múltiplos handlers
//
// sem impactar os consumidores.
func New() *slog.Logger {
	handler := slog.NewTextHandler(os.Stderr, nil)

	return slog.New(handler)
}
