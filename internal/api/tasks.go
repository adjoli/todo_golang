package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/adjoli/todo_chatgpt/internal/models"
	"github.com/adjoli/todo_chatgpt/internal/service"
)

// ----------------------------------------------
func (s *Server) listTasks(
	w http.ResponseWriter,
	r *http.Request,
) {
	tasks, err := s.taskService.ListTasks(
		r.Context(),
		models.TaskFilter{},
	)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	if err := writeJSON(w, http.StatusOK, tasks); err != nil {
		writeError(
			w,
			http.StatusInternalServerError,
			"failed to encode response",
		)
	}
}

// ----------------------------------------------
func (s *Server) createTask(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid JSON",
		)
		return
	}

	if strings.TrimSpace(req.Title) == "" {
		writeError(
			w,
			http.StatusBadRequest,
			"title is required",
		)
		return
	}

	task, err := s.taskService.CreateTask(
		r.Context(),
		service.CreateTaskInput{
			Title: req.Title,
		},
	)
	if err != nil {
		s.logger.Error("create task", "error", err)
		writeError(
			w,
			http.StatusInternalServerError,
			"internal server error",
		)
		return
	}

	resp := TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Completed: task.Completed,
	}

	if err := writeJSON(
		w,
		http.StatusCreated,
		resp,
	); err != nil {
		s.logger.Error("write response", "error", err)
	}
}

// ----------------------------------------------
func (s *Server) getTask(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := r.PathValue("id")

	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid task id",
		)
		return
	}

	task, err := s.taskService.GetTask(
		r.Context(),
		taskID,
	)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			writeError(
				w,
				http.StatusNotFound,
				"task not found",
			)
			return
		}

		s.logger.Error("get task", "error", err)
		writeError(
			w,
			http.StatusInternalServerError,
			"internal server error",
		)
		return
	}

	resp := TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Completed: task.Completed,
	}

	writeJSON(
		w,
		http.StatusOK,
		resp,
	)
}

// ----------------------------------------------
func (s *Server) updateTask(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := r.PathValue("id")

	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid task id",
		)
		return
	}

	var req UpdateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	input := service.UpdateTaskInput{
		Title: req.Title,
	}

	if err = s.taskService.UpdateTask(
		r.Context(),
		taskID,
		input,
	); err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			writeError(
				w,
				http.StatusNotFound,
				"task not found",
			)
			return
		}

		s.logger.Error("update task", "error", err)

		writeError(
			w,
			http.StatusInternalServerError,
			"internal server error",
		)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}