package service

import "errors"

// ErrEmptyTitle é retornado quando o título fornecido para uma tarefa
// é vazio ou contém apenas espaços em branco.
var ErrEmptyTitle = errors.New("title cannot be empty")

// ErrTaskNotFound é retornado quando a tarefa procurada pelo ID
// não existe no banco de dados.
var ErrTaskNotFound = errors.New("task not found")

// ErrTaskAlreadyCompleted é retornado quando se tenta concluir uma
// tarefa que já está marcada como concluída.
var ErrTaskAlreadyCompleted = errors.New("task already completed")
