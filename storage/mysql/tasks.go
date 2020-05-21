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

const taskQueryFields = "id,name,status,position,year,month,week,day,created_at,updated_at"

// Add a task
func (t *Tasks) Add(task *models.Task) (err error) {
	var position int
	tx, err := t.engine.Beginx()
	if err != nil {
		return
	}
	defer func() {
		if err == nil {
			return
		}
		if errRb := tx.Rollback(); errRb != nil {
			log.Warnf("transaction rollback error: %v", errRb)
		}
	}()

	err = tx.Get(&position, "SELECT IFNULL(MAX(position), 0)+1 FROM task")
	if err != nil {
		return
	}

	res, err := t.engine.Exec(
		"INSERT INTO task (name, position) VALUES(?,?)",
		task.Name, position,
	)
	if err != nil {
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		return
	}
	task.ID = id

	err = tx.Commit()
	return
}

// Task DB model
type Task struct {
	ID        int64             `db:"id"`
	Name      string            `db:"name"`
	Status    models.TaskStatus `db:"status"`
	Position  int64             `db:"position"`
	Year      int               `db:"year"`
	Month     int               `db:"month"`
	Week      int               `db:"week"`
	Day       int               `db:"day"`
	CreatedAt time.Time         `db:"created_at"`
	UpdatedAt time.Time         `db:"updated_at"`
}

func (t Task) fromModel(task *models.Task) {
	t.ID = task.ID
	t.Name = task.Name
	t.Status = task.Status
	t.Position = task.Position
	t.Year = task.Plan.Year
	t.Month = task.Plan.Month
	t.Week = task.Plan.Week
	t.Day = task.Plan.Day
	t.CreatedAt = task.CreatedAt
	t.UpdatedAt = task.UpdatedAt
}

func (t Task) toModel() models.Task {
	ret := new(models.Task)
	ret.ID = t.ID
	ret.Name = t.Name
	ret.Status = t.Status
	ret.Position = t.Position
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
	sql := fmt.Sprintf("SELECT %s FROM task ORDER BY position DESC", taskQueryFields)
	err = t.engine.Select(&tasks, sql)
	ret = tasksToModels(tasks)
	return
}

// UpdateTask ...
func (t *Tasks) UpdateTask(task models.Task) error {
	SQL := fmt.Sprintf(`UPDATE task SET name=?, status=?, year=?, month=?, week=?, day=? WHERE id=?`)
	_, err := t.engine.Exec(SQL, task.Name, task.Status, task.Plan.Year, task.Plan.Month, task.Plan.Week, task.Plan.Day, task.ID)
	return err
}

// ChangePosition by given id, tPos (target postion)
func (t *Tasks) ChangePosition(id int64, tPos int) error {
	tx, err := t.engine.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			err = tx.Commit()
			if err == nil {
				return
			}
		}
		if errRb := tx.Rollback(); errRb != nil {
			log.Warnf("transaction rollback error: %v", errRb)
		}
	}()
	var cPos int // current position
	err = tx.Get(&cPos, "SELECT position FROM task WHERE id=?", id)
	if err != nil {
		return err
	}

	if tPos > cPos {
		_, err = tx.Exec("UPDATE task SET position=position - 1 WHERE position > ? and position <= ?", cPos, tPos)
		if err != nil {
			return err
		}
	} else if tPos < cPos {
		_, err = tx.Exec("UPDATE task SET position=position + 1 WHERE position >= ? and position < ?", tPos, cPos)
		if err != nil {
			return err
		}
	} else {
		return nil
	}

	_, err = tx.Exec("UPDATE task SET position=? WHERE id=?", tPos, id)
	return err
}

// DeleteTask ...
func (t *Tasks) DeleteTask(id int64) error {
	tx, err := t.engine.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			err = tx.Commit()
			if err == nil {
				return
			}
		}
		if errRb := tx.Rollback(); errRb != nil {
			log.Warnf("transaction rollback error: %v", errRb)
		}
	}()

	var cPos int // current position
	err = tx.Get(&cPos, "SELECT position FROM task WHERE id=?", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM task WHERE id=?`, id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE task SET position = position - 1 WHERE position > ?`, cPos)
	return err
}
