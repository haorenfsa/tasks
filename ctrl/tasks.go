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
	g.POST("/", a.handleAddTask)
	g.GET("", a.handleQueryTasks)
	g.PATCH("", a.handleUpdateTask)
	g.DELETE("/:id", a.handleDeleteTask)
	g.PUT("/:id/position/:position", a.handleChangePosition)
}

func (a *TaskController) handleAddTask(c *gin.Context) {
	task := new(models.Task)
	err := c.ShouldBindJSON(task)
	if err != nil {
		log.Warn("handleUpdateTask:", err)
		c.JSON(http.StatusBadRequest, errMsgBadBody)
		return
	}
	err = a.core.Add(task)
	writeObjectResponseByError(c, task, err)
}

func (a *TaskController) handleQueryTasks(c *gin.Context) {
	ret, err := a.core.QueryAll()
	writeObjectResponseByError(c, ret, err)
}

func (a *TaskController) handleUpdateTask(c *gin.Context) {
	task := new(models.Task)
	err := c.ShouldBindJSON(task)
	if err != nil {
		log.Warn("handleUpdateTask:", err)
		c.JSON(http.StatusBadRequest, errMsgBadBody)
		return
	}
	err = a.core.UpdateTask(*task)
	writeMsgResponseByError(c, err)
}

func (a *TaskController) handleDeleteTask(c *gin.Context) {
	id, err := parseIntParam(c, "id")
	if err != nil {
		writeMsgResponseByError(c, err)
		return
	}
	err = a.core.DeleteTask(id)
	writeMsgResponseByError(c, err)
}

func (a *TaskController) handleChangePosition(c *gin.Context) {
	id, err := parseIntParam(c, "id")
	if err != nil {
		writeMsgResponseByError(c, err)
	}
	position, err := parseIntParam(c, "position")
	if err != nil {
		writeMsgResponseByError(c, err)
	}
	err = a.core.ChangePosition(id, int(position))
	writeMsgResponseByError(c, err)
}
