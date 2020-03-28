package core

import (
	"log"
	"time"

	"github.com/haorenfsa/tasks/errs"
	"github.com/haorenfsa/tasks/models"
	"github.com/haorenfsa/tasks/storage/mysql"
)

// Tasks implements tasks methods interfaces
type Tasks struct {
	storage *mysql.Tasks
}

// NewTasks builds a New Tasks
func NewTasks(storage *mysql.Tasks) *Tasks {
	return &Tasks{storage: storage}
}

// Add a task
func (t *Tasks) Add(task *models.Task) error {
	if task.Name == "" {
		return errs.ErrBadRequest
	}
	if time.Since(task.DueTime) > 0 {
		task.DueTime = time.Now().Add(time.Hour * 24 * 30)
	}
	task.Status = models.TaskStatusPending
	err := t.storage.Add(*task)
	if err != nil {
		log.Printf("add task[%s] failed: %s", task.Name, err)
		return errs.ErrStorage
	}
	return nil
}

func (t *Tasks) Get(name string) (models.Task, error) {

}
