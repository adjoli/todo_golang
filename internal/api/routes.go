package api

import "net/http"

func (s *Server) registerRoutes(mux *http.ServeMux) {
	// Home
	mux.HandleFunc("GET /", s.home)

	// Tasks
	mux.HandleFunc("GET /tasks", s.listTasks)
	mux.HandleFunc("POST /tasks", s.createTask)
	mux.HandleFunc("GET /tasks/{id}", s.getTask)
	mux.HandleFunc("PUT /tasks/{id}", s.updateTask)
}
