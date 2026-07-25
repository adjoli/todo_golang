package api

import "net/http"

// registerRoutes registra todas as rotas da API no ServeMux fornecido.
func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /", s.home)

	mux.HandleFunc("GET /tasks", s.listTasks)
	mux.HandleFunc("POST /tasks", s.createTask)
	mux.HandleFunc("GET /tasks/{id}", s.getTask)
	mux.HandleFunc("PUT /tasks/{id}", s.updateTask)
}
