package core

import (
	"github.com/haorenfsa/tasks/storage"
)

// Tasks implements tasks methods interfaces
type Tasks struct {
	storage.TaskStorage
}

// NewTasks builds a New Tasks
func NewTasks(storage storage.TaskStorage) *Tasks {
	return &Tasks{TaskStorage: storage}
}
