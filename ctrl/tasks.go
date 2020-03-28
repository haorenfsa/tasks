package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haorenfsa/tasks/core"
	"github.com/haorenfsa/tasks/errs"
	"github.com/haorenfsa/tasks/models"
	"github.com/tevino/log"
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
	root.POST("/tasks", a.handleAddTask)
}

func (a *TaskController) handleAddTask(c *gin.Context) {
	task := new(models.Task)
	err := c.ShouldBindJSON(task)
	if err != nil {
		log.Warn("handleAddTask:", err)
		c.JSON(http.StatusBadRequest, errMsgBadBody)
		return
	}
	err = a.core.Add(task)
	if err != nil {
		code := errs.ErrorToHTTPCode(err)
		c.JSON(code, errMsgActionFailed)
		return
	}
	c.JSON(http.StatusOK, msgSuccess)
	return
}

const errMsgBadBody = "bad body format"
const errMsgActionFailed = "action failed"
const msgSuccess = "success"
