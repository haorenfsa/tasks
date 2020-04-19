package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefault(t *testing.T) {
	dsn := DSN{
		Address:  "",
		UserName: "root",
		Password: "",
		DBName:   "tasks",
	}
	ret, err := NewDefault(dsn)
	assert.NoError(t, err)
	assert.Nil(t, ret)
}
