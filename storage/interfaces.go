package storage

import "github.com/haorenfsa/tasks/models"

// TaskStorage abstracts interface for storage of tasks
type TaskStorage interface {
	Add(task *models.Task) (err error)
	DeleteTask(id int64) error
	QueryAll() (ret []models.Task, err error)
	ChangePosition(id int64, tPos int) error
	UpdateTask(task models.Task) error
}
