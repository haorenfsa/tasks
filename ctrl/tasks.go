package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tevino/log"

	"github.com/haorenfsa/tasks/core"
	"github.com/haorenfsa/tasks/models"
)

// TaskController is controller for Task
// implements server.Controller
type TaskController struct {
	core core.Tasks
}

// NewTaskController builds a New TaskController
func NewTaskController(core core.Tasks) *TaskController {
	return &TaskController{core: core}
}

// Register registers request handler
func (a *TaskController) Register(root gin.IRouter) {
	g := root.Group("/tasks")
	g.PUT("/:name", a.handleAddTask)
	g.GET("", a.handleQueryTasks)
	g.PATCH("/:name", a.handleUpdateTask)
	g.DELETE("/:name", a.handleDeleteTask)
}

func (a *TaskController) handleAddTask(c *gin.Context) {
	name := c.Param("name")
	err := a.core.Add(name)
	writeMsgResponseByError(c, err)
}

func (a *TaskController) handleQueryTasks(c *gin.Context) {
	ret, err := a.core.QueryAll()
	writeObjectResponseByError(c, ret, err)
}

func (a *TaskController) handleUpdateTask(c *gin.Context) {
	name := c.Param("name")
	task := new(models.Task)
	err := c.ShouldBindJSON(task)
	if err != nil {
		log.Warn("handleUpdateTask:", err)
		c.JSON(http.StatusBadRequest, errMsgBadBody)
		return
	}
	err = a.core.UpdateTask(name, *task)
	writeMsgResponseByError(c, err)
}

func (a *TaskController) handleDeleteTask(c *gin.Context) {
	name := c.Param("name")
	err := a.core.DeleteTask(name)
	writeMsgResponseByError(c, err)
}
