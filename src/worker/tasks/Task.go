package tasks

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/backends/result"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/ysfgrl/go-fiber-api/src/models"
	"reflect"
)

type TaskInterface interface {
	run(id string) error
	SendSync(args []tasks.Arg) ([]reflect.Value, *models.Error)
	SendAsync(args []tasks.Arg) (*result.AsyncResult, *models.Error)
	Register() *models.Error
	Test()
	GetName() string
	SetServer(server *machinery.Server)
}

type BaseTask struct {
	server *machinery.Server
	Name   string
}
