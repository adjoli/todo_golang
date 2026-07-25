package api

// CreateTaskRequest é o payload de entrada para criação de uma tarefa.
type CreateTaskRequest struct {
	Title string `json:"title"`
}

// UpdateTaskRequest é o payload de entrada para atualização de uma tarefa.
type UpdateTaskRequest struct {
	Title string `json:"title"`
}

// TaskResponse é a representação de uma tarefa na resposta HTTP.
type TaskResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
