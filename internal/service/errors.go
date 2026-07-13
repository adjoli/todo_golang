package service

import "errors"

var (
	ErrEmptyTitle           = errors.New("title cannot be empty")
	ErrTaskNotFound         = errors.New("task not found")
	ErrTaskAlreadyCompleted = errors.New("task already completed")
)
