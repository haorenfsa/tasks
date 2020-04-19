package interfaces

import (
	"time"

	"github.com/haorenfsa/tasks/models"
)

// Tasks methods
type Tasks interface {
	Add(models.Task) error

	Get(name string) (models.Task, error)
	QueryAll() ([]models.Task, error)
	QueryByWeek(models.Week) ([]models.Task, error)
	QueryByDay(models.Day) ([]models.Task, error)
	QueryByProject(models.Project) ([]models.Task, error)

	UpdateStatus(name string, status models.TaskStatus) error
	DoAtCertainTime(name string, time time.Time) error

	AddProject(task string, project string)
	AssignToWeekPlan(task string, week models.Week)
	AssignToDayPlan(task string, week models.Day)
}
