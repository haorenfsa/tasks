package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/haorenfsa/tasks/errs"
	"github.com/haorenfsa/tasks/models"
	"github.com/haorenfsa/tasks/storage/mysql"
	"github.com/haorenfsa/tasks/test"
)

func TestTasks_AddOK(t *testing.T) {
	engine := test.NewTestEngine()
	defer engine.CleanUp()
	storage := mysql.NewTasks(engine.Default)
	tasks := NewTasks(storage)

	task := &models.Task{
		Name: "test",
	}
	err := tasks.Add(task)
	assert.NoError(t, err)
	assert.True(t, time.Since(task.DueTime) <= 0)
}

func TestTasks_AddEmptyName(t *testing.T) {
	engine := test.NewTestEngine()
	defer engine.CleanUp()
	storage := mysql.NewTasks(engine.Default)
	tasks := NewTasks(storage)

	task := new(models.Task)
	err := tasks.Add(task)
	assert.Error(t, err)
	assert.Equal(t, errs.ErrBadRequest, err)
}
