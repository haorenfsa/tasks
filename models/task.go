package models

import "time"

// Task model
type Task struct {
	Name    string
	DueTime time.Time
	Status  TaskStatus
}

// TaskStatus enum
type TaskStatus int

// TaskStatus values
const (
	TaskStatusTODO = iota
	TaskStatusDoing
	TaskStatusDone
	TaskStatusPending
	TaskStatusClosed
)
