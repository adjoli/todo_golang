package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/adjoli/todo_chatgpt/internal/app"
	"github.com/adjoli/todo_chatgpt/internal/service"
)

type Server struct {
	logger      *slog.Logger
	taskService *service.TaskService
}

func New(app *app.App) *Server {
	return &Server{
		logger:      app.Logger(),
		taskService: app.TaskService(),
	}
}

func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()

	s.registerRoutes(mux)

	fmt.Printf("HTTP server listening on %s\n", addr)

	return http.ListenAndServe(addr, mux)
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Task Manager API")
}
