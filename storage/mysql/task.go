package mysql

import (
	"fmt"

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

const taskQueryFields = "id,name,status,position,schedule_level,start_time,end_time,created_at,updated_at"

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
type Task = models.Task

// QueryAll tasks
func (t *Tasks) QueryAll() (ret []models.Task, err error) {
	sql := fmt.Sprintf("SELECT %s FROM task ORDER BY position DESC", taskQueryFields)
	err = t.engine.Select(&ret, sql)
	return
}

// UpdateTask ...
func (t *Tasks) UpdateTask(task models.Task) error {
	SQL := fmt.Sprintf(`UPDATE task SET name=?, status=?, schedule_level=?, start_time=?, end_time=? WHERE id=?`)
	_, err := t.engine.Exec(SQL, task.Name, task.Status, task.ScheduleLevel, task.StartTime, task.EndTime, task.ID)
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
