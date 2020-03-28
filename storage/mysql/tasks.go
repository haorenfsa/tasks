package mysql

import (
	"time"

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

// DBTask model of task in db
type DBTask struct {
	ID        int64             `db:"id"`
	Name      string            `db:"name"`
	DueTime   time.Time         `db:"due_time"`
	Status    models.TaskStatus `db:"status"`
	CreatedAt time.Time         `db:"created_at"`
	UpdatedAt time.Time         `db:"updated_at"`
}

// Add a task
func (t *Tasks) Add(task models.Task) error {
	_, err := t.engine.Exec(
		"INSERT INTO tasks (name,due_time,status) VALUES(?,?,?)",
		task.Name, task.DueTime, task.Status,
	)
	return err
}

// Get a task
func (t *Tasks) Get(name string) (models.Task, error) {
	// TODO: @shaoyue.chen
	// _, err := t.engine.Exec("SELECT * FROM ")
	return models.Task{}, nil
}

// QueryAll tasks
func (t *Tasks) QueryAll() ([]models.Task, error) {
	// TODO: @shaoyue.chen
	return nil, nil
}

// QueryByWeek ...
func (t *Tasks) QueryByWeek(models.Week) ([]models.Task, error) {
	// TODO: @shaoyue.chen
	return nil, nil
}

// QueryByDay ...
func (t *Tasks) QueryByDay(models.Day) ([]models.Task, error) {
	// TODO: @shaoyue.chen
	return nil, nil
}

// QueryByProject ...
func (t *Tasks) QueryByProject(models.Project) ([]models.Task, error) {
	// TODO: @shaoyue.chen
	return nil, nil
}
