package tasks

import (
	"context"
	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/backends/result"
	"github.com/RichardKnop/machinery/v2/log"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/opentracing/opentracing-go"
	response2 "go-fiber-api/src/models"
	"go-fiber-api/src/utils/response"
	"reflect"
	"time"
)

type testTask struct {
	BaseTask
}

func NewTestTask() TaskInterface {
	return &testTask{
		BaseTask{
			Name: "test_task",
		},
	}
}

func (d *testTask) GetName() string {
	return d.Name
}

func (d *testTask) run(id string) error {

	log.INFO.Println("Test ")
	log.INFO.Println("I am " + d.BaseTask.Name + " Task")
	log.INFO.Println("Id: " + id + " Task")
	return nil
}

func (d *testTask) Register() *response2.MyError {
	err := d.server.RegisterTask(d.BaseTask.Name, d.run)
	if err != nil {
		return response.GetError(err)
	}
	return nil
}

func (d *testTask) Test() {
	d.SendAsync([]tasks.Arg{
		{
			Name:  "id",
			Type:  "string",
			Value: "64dd7a3831779826103cb3ac",
		},
	})
}
func (d *testTask) SetServer(server *machinery.Server) {
	d.server = server
}

func (d *testTask) SendSync(args []tasks.Arg) ([]reflect.Value, *response2.MyError) {
	asyncResult, err := d.SendAsync(args)
	if err != nil {
		return nil, err
	}
	results, err1 := asyncResult.Get(time.Millisecond * 5)
	if err != nil {
		return nil, response.GetError(err1)
	}
	return results, nil
}

func (d *testTask) SendAsync(args []tasks.Arg) (*result.AsyncResult, *response2.MyError) {

	span, ctx := opentracing.StartSpanFromContext(context.Background(), "send")
	defer span.Finish()
	asyncResult, err := d.server.SendTaskWithContext(ctx, &tasks.Signature{
		Name: d.Name,
		Args: args,
	})
	if err != nil {
		return nil, response.GetError(err)
	}
	return asyncResult, nil
}
