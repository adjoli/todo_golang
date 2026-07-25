// Package api implementa o servidor HTTP que expõe a API REST
// para gerenciamento de tarefas, utilizando a API padrão net/http.
package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/adjoli/todo_chatgpt/internal/app"
	"github.com/adjoli/todo_chatgpt/internal/service"
)

// Server é o servidor HTTP da API de tarefas.
// Ele encapsula o logger e o service, e registra as rotas HTTP.
type Server struct {
	logger      *slog.Logger
	taskService *service.TaskService
}

// New cria um Server a partir da instância da aplicação,
// extraindo o logger e o TaskService.
func New(app *app.App) *Server {
	return &Server{
		logger:      app.Logger(),
		taskService: app.TaskService(),
	}
}

// Start registra as rotas e inicia o servidor HTTP no endereço
// informado. O servidor bloqueia até ocorrer um erro.
func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()

	s.registerRoutes(mux)

	fmt.Printf("HTTP server listening on %s\n", addr)

	return http.ListenAndServe(addr, mux)
}

// home é o handler da rota raiz. Retorna uma mensagem simples
// indicando que a API está funcionando.
func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Task Manager API")
}
