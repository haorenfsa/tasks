package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/haorenfsa/tasks/core"
	"github.com/haorenfsa/tasks/ctrl"

	"github.com/haorenfsa/tasks/server"
	"github.com/haorenfsa/tasks/storage/mysql"
	"github.com/haorenfsa/tasks/storage/mysql/engine"
)

func main() {
	var staticPath string
	var port int
	flag.StringVar(&staticPath, "s", "", "static file path to serve, not serve when empty")
	flag.IntVar(&port, "p", 3001, "server port")
	flag.Parse()
	log.Print(staticPath)

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
	theServer.ServeStaticPath(staticPath)
	theServer.Run(fmt.Sprintf(":%d", port))
}
