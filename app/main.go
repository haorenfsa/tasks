package main

import (
	"github.com/haorenfsa/tasks/core"
	"github.com/haorenfsa/tasks/ctrl"

	"github.com/haorenfsa/tasks/server"
	"github.com/haorenfsa/tasks/storage/mysql"
	"github.com/haorenfsa/tasks/storage/mysql/engine"
)

func main() {
	theServer := server.NewHTTPServer()

	dsn := engine.DSN{
		UserName: "root",
		DBName:   "tasks",
	}
	engine, err := engine.NewDefault(dsn)
	if err != nil {
		panic(err)
	}
	taskStorage := mysql.NewTasks(engine)
	tasks := core.NewTasks(taskStorage)
	taskController := ctrl.NewTaskController(*tasks)

	ctls := []server.Controller{taskController}
	theServer.UseControllers(ctls)
	theServer.Run(":8080")
}
