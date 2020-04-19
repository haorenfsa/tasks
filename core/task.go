package core

import (
	"github.com/haorenfsa/tasks/storage/mysql"
)

// Tasks implements tasks methods interfaces
type Tasks struct {
	*mysql.Tasks
}

// NewTasks builds a New Tasks
func NewTasks(storage *mysql.Tasks) *Tasks {
	return &Tasks{Tasks: storage}
}
