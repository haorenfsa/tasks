package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/haorenfsa/tasks/storage/mysql"
	"github.com/haorenfsa/tasks/test"
)

func TestTasks_Add_OK(t *testing.T) {
	engine := test.NewTestEngine()
	defer engine.CleanUp()
	storage := mysql.NewTasks(engine.Default)
	tasks := NewTasks(storage)
	err := tasks.Add("test")
	assert.NoError(t, err)
}

func TestTasks_Add_FailedDuplicated(t *testing.T) {
	engine := test.NewTestEngine()
	defer engine.CleanUp()
	storage := mysql.NewTasks(engine.Default)
	tasks := NewTasks(storage)

	err := tasks.Add("test")
	assert.NoError(t, err)
	err = tasks.Add("test")
	assert.Error(t, err)
}

func TestTasks_QueryAll_OKZero(t *testing.T) {
	engine := test.NewTestEngine()
	defer engine.CleanUp()
	storage := mysql.NewTasks(engine.Default)
	tasks := NewTasks(storage)

	ret, err := tasks.QueryAll()
	assert.NoError(t, err)
	assert.Len(t, ret, 0)
}

func TestTasks_QueryAll_OKOne(t *testing.T) {
	engine := test.NewTestEngine()
	defer engine.CleanUp()
	storage := mysql.NewTasks(engine.Default)
	tasks := NewTasks(storage)
	err := tasks.Add("test")
	assert.NoError(t, err)

	ret, err := tasks.QueryAll()
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
}

func TestTasks_QueryAll_OKTwo(t *testing.T) {
	engine := test.NewTestEngine()
	defer engine.CleanUp()
	storage := mysql.NewTasks(engine.Default)
	tasks := NewTasks(storage)

	err := tasks.Add("test")
	err = tasks.Add("test2")
	assert.NoError(t, err)

	ret, err := tasks.QueryAll()
	assert.NoError(t, err)
	assert.Len(t, ret, 2)
}
