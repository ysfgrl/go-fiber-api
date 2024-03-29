package tasks

import (
	"context"
	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/backends/result"
	"github.com/RichardKnop/machinery/v2/log"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/opentracing/opentracing-go"
	response2 "github.com/ysfgrl/go-fiber-api/src/models"
	"github.com/ysfgrl/go-fiber-api/src/pkg/response"
	"github.com/ysfgrl/go-fiber-api/src/services"
	"reflect"
	"time"
)

type downloadTask struct {
	BaseTask
	service services.ResourceService
}

func NewDownloadTask(service services.ResourceService) TaskInterface {
	return &downloadTask{
		BaseTask{
			Name: "download_helper",
		},
		service,
	}
}

func (d *downloadTask) GetName() string {
	return d.Name
}

func (d *downloadTask) run(id string) error {

	resource, err := d.service.GetResource(id)
	if err != nil {
		log.INFO.Println(err.Detail)
		return nil
	}
	log.INFO.Println(resource.Url)
	log.INFO.Println("I am " + d.BaseTask.Name + " Task")
	log.INFO.Println("Id: " + id + " Task")
	return nil
}

func (d *downloadTask) Register() *response2.Error {
	err := d.server.RegisterTask(d.BaseTask.Name, d.run)
	if err != nil {
		return response.GetError(err)
	}
	return nil
}

func (d *downloadTask) Test() {
	d.SendAsync([]tasks.Arg{
		{
			Name:  "id",
			Type:  "string",
			Value: "64dd7a3831779826103cb3ac",
		},
	})
}

func (d *downloadTask) SetServer(server *machinery.Server) {
	d.server = server
}

func (d *downloadTask) SendSync(args []tasks.Arg) ([]reflect.Value, *response2.Error) {
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

func (d *downloadTask) SendAsync(args []tasks.Arg) (*result.AsyncResult, *response2.Error) {

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
