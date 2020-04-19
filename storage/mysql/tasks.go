package mysql

import (
	"fmt"

	"github.com/haorenfsa/tasks/models"
	"github.com/haorenfsa/tasks/storage/mysql/engine"
)

// Tasks stores tasks
type Tasks struct {
	engine *engine.Default
}

// NewTasks builds a New Tasks
func NewTasks(engine *engine.Default) *Tasks {
	return &Tasks{engine: engine}
}

const taskQueryFields = "id,name,status,created_at,updated_at"

// Add a task
func (t *Tasks) Add(name string) error {
	_, err := t.engine.Exec(
		"INSERT INTO task (name) VALUES(?)",
		name,
	)
	return err
}

// QueryAll tasks
func (t *Tasks) QueryAll() (ret []models.Task, err error) {
	sql := fmt.Sprintf("SELECT %s FROM task ORDER BY id DESC", taskQueryFields)
	err = t.engine.Select(&ret, sql)
	return
}

// UpdateTask ...
func (t *Tasks) UpdateTask(name string, task models.Task) error {
	SQL := fmt.Sprintf(`UPDATE task SET name=?, status=?, year=?, month=?, week=?, day=? WHERE name=?`)
	_, err := t.engine.Exec(SQL, task.Name, task.Status, task.Year, task.Month, task.Week, task.Day, name)
	return err
}

// DeleteTask ...
func (t *Tasks) DeleteTask(name string) error {
	SQL := fmt.Sprintf(`DELETE FROM task WHERE name=?`)
	_, err := t.engine.Exec(SQL, name)
	return err
}
