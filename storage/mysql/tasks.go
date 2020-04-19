package mysql

import (
	"fmt"
	"time"

	"github.com/tevino/log"

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

const taskQueryFields = "id,name,status,year,month,week,day,created_at,updated_at"

// Add a task
func (t *Tasks) Add(name string) error {
	_, err := t.engine.Exec(
		"INSERT INTO task (name) VALUES(?)",
		name,
	)
	return err
}

// Task DB model
type Task struct {
	ID        int64             `db:"id"`
	Name      string            `db:"name"`
	Status    models.TaskStatus `db:"status"`
	Year      int               `db:"year"`
	Month     int               `db:"month"`
	Week      int               `db:"week"`
	Day       int               `db:"day"`
	CreatedAt time.Time         `db:"created_at"`
	UpdatedAt time.Time         `db:"updated_at"`
}

func (t Task) toModel() models.Task {
	ret := new(models.Task)
	ret.ID = t.ID
	ret.Name = t.Name
	ret.Status = t.Status
	ret.Plan.Year = t.Year
	ret.Plan.Month = t.Month
	ret.Plan.Week = t.Week
	ret.Plan.Day = t.Day
	ret.CreatedAt = t.CreatedAt
	ret.UpdatedAt = t.UpdatedAt
	log.Print(ret)
	return *ret
}

func tasksToModels(tasks []Task) []models.Task {
	ret := make([]models.Task, len(tasks))
	for i, task := range tasks {
		ret[i] = task.toModel()
	}
	return ret
}


// QueryAll tasks
func (t *Tasks) QueryAll() (ret []models.Task, err error) {
	var tasks []Task
	sql := fmt.Sprintf("SELECT %s FROM task ORDER BY id DESC", taskQueryFields)
	err = t.engine.Select(&tasks, sql)
	ret = tasksToModels(tasks)
	return
}

// UpdateTask ...
func (t *Tasks) UpdateTask(name string, task models.Task) error {
	SQL := fmt.Sprintf(`UPDATE task SET name=?, status=?, year=?, month=?, week=?, day=? WHERE name=?`)
	log.Print(SQL, task)
	_, err := t.engine.Exec(SQL, task.Name, task.Status, task.Plan.Year, task.Plan.Month, task.Plan.Week, task.Plan.Day, name)
	return err
}

// DeleteTask ...
func (t *Tasks) DeleteTask(name string) error {
	SQL := fmt.Sprintf(`DELETE FROM task WHERE name=?`)
	_, err := t.engine.Exec(SQL, name)
	return err
}
