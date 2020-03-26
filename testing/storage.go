package testing

import "github.com/haorenfsa/tasks/storage/mysql/engine"

// NewDBEngineForTest create a default db engine for test
func NewDBEngineForTest() *engine.Default {
	dsn := engine.DSN{
		Address:  "",
		UserName: "root",
		Password: "",
		DBName:   "tasks_test",
	}
	ret, err := engine.NewDefault(dsn)
	if err != nil {
		panic(err)
	}
	// TODO: @shaoyue.chen init
	ret.Exec("")
	return ret
}
