package models

import "time"

// Task model
type Task struct {
	ID            int64      `json:"id" db:"id"`
	Name          string     `json:"name" db:"name"`
	Status        TaskStatus `json:"status" db:"status"`
	Position      int64      `json:"position" db:"position"`
	ScheduleLevel TimeLevel  `json:"scheduleLevel" db:"schedule_level"`
	StartTime     time.Time  `json:"startTime" db:"start_time"`
	EndTime       time.Time  `json:"endTime" db:"end_time"`
	CreatedAt     time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time  `json:"updatedAt" db:"updated_at"`
}

// TimeLevel is the level of LeveledTime like year,month,week,day
type TimeLevel int

// TimeLevel enum definitions
const (
	NotSet TimeLevel = iota
	Year
	Month
	Week
	Day
)

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
