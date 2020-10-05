package models

import "time"

// Task model
type Task struct {
	ID         int64      `json:"id"`
	ParentTask int64      `json:"parent_task"`
	SubTasks   []int64    `json:"sub_tasks"`
	Name       string     `json:"name"`
	Status     TaskStatus `json:"status"`
	Position   int64      `json:"position"`
	Plan       TaskPlan   `json:"plan"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type TaskPlan struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Week  int `json:"week"`
	Day   int `json:"day"`
}

// NOTSET int default value
const NOTSET = -1

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
