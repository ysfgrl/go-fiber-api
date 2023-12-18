package tasks

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/backends/result"
	"github.com/RichardKnop/machinery/v2/tasks"
	"go-fiber-api/src/models"
	"reflect"
)

type TaskInterface interface {
	run(id string) error
	SendSync(args []tasks.Arg) ([]reflect.Value, *models.MyError)
	SendAsync(args []tasks.Arg) (*result.AsyncResult, *models.MyError)
	Register() *models.MyError
	Test()
	GetName() string
	SetServer(server *machinery.Server)
}

type BaseTask struct {
	server *machinery.Server
	Name   string
}
