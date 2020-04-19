package test

import (
	"fmt"
	"time"

	"github.com/haorenfsa/tasks/storage/mysql/engine"
)

// TestEngine for unit test
type TestEngine struct {
	*engine.Default
	dbName string
}

var tableSQLs = []string{
	`CREATE TABLE IF NOT EXISTS task (
		id BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'primary key',
		name VARCHAR(128) NOT NULL COMMENT 'name, unique',
		year SMALLINT(5) NOT NULL DEFAULT -1 COMMENT 'planed year, -1 means not set',
		month TINYINT(2) NOT NULL DEFAULT -1 COMMENT 'planed month, -1 means not set',
		week TINYINT(2) NOT NULL DEFAULT -1 COMMENT 'planed week, -1 means not set',
		day TINYINT(1) NOT NULL DEFAULT -1 COMMENT 'planed day, -1 means not set',
		status TINYINT NOT NULL DEFAULT 0 COMMENT 'task status',
		created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time auto',
		updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time auto',
		PRIMARY KEY (id),
		UNIQUE KEY uk_name (name)
	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tasks';`,
}

// CleanUp must be called after test done
func (t *TestEngine) CleanUp() {
	t.Default.DB.Exec(fmt.Sprintf("DROP DATABASE %s", t.dbName))
}

// NewTestEngine create a default db engine for test
func NewTestEngine() *TestEngine {
	dsn := engine.DSN{
		Address:  "",
		UserName: "root",
		Password: "",
		DBName:   "",
	}
	engine, err := engine.NewDefault(dsn)
	if err != nil {
		panic(err)
	}

	dbName := fmt.Sprintf("test_tasks_%d", time.Now().Unix())
	_, err = engine.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		panic(err)
	}
	_, err = engine.Exec(fmt.Sprintf("use %s", dbName))
	if err != nil {
		panic(err)
	}

	for _, tableSQL := range tableSQLs {
		engine.Exec(tableSQL)
	}
	return &TestEngine{
		Default: engine,
		dbName:  dbName,
	}
}
