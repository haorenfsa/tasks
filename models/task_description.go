package models

import "time"

// TaskDescription ..
type TaskDescription struct {
	ID          int64     `json:"id" db:"id"`
	TaskID      int64     `json:"task_id" db:"task_id"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
