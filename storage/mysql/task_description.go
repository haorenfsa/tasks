package mysql

import (
	"fmt"

	"github.com/haorenfsa/tasks/models"
	"github.com/haorenfsa/tasks/storage/mysql/engine"
)

// TaskDescriptions stores tasks descriptions
type TaskDescriptions struct {
	engine *engine.Default
}

const tableName = "task_description"

const taskDescQueryFields = "id,task_id,description,created_at,updated_at"

func (t TaskDescriptions) SetByTaskID(id int64, desc *models.TaskDescription) error {
	_, err := t.engine.Exec(fmt.Sprintf("REPLACE INTO %s (task_id,description) values(%d, \"%s\")", tableName, id, desc.Description))
	return err
}

func (t TaskDescriptions) GetByTaskID(id int64) (*models.TaskDescription, error) {
	var ret *models.TaskDescription
	err := t.engine.Get(ret, fmt.Sprintf("SELECT %s from %s WHERE task_id=%d", tableName, taskDescQueryFields, id))
	return ret, err
}
