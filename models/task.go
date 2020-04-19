package models

import "time"

// Task model
type Task struct {
	ID        int64      `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	Status    TaskStatus `db:"status" json:"status"`
	Year      int        `db:"year" json:"year"`
	Month     int        `db:"month" json:"month"`
	Week      int        `db:"week" json:"week"`
	Day       int        `db:"day" json:"day"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
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
