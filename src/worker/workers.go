package worker

import (
	"github.com/RichardKnop/machinery/v2"
	amqpbackend "github.com/RichardKnop/machinery/v2/backends/amqp"
	amqpbroker "github.com/RichardKnop/machinery/v2/brokers/amqp"
	"github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/example/tracers"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/RichardKnop/machinery/v2/log"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/ysfgrl/go-fiber-api/src/clients"
	appconfig "github.com/ysfgrl/go-fiber-api/src/config"
	"github.com/ysfgrl/go-fiber-api/src/models/elastic_collections"
	apptask "github.com/ysfgrl/go-fiber-api/src/worker/tasks"
)

var myWorker *worker = nil

type worker struct {
	server *machinery.Server
	tasks  map[string]apptask.TaskInterface
}

func NewWorker() *worker {
	if myWorker != nil {
		return myWorker
	}
	cnf := &config.Config{
		Broker:          appconfig.Rabbit.Url,
		DefaultQueue:    appconfig.Rabbit.Que,
		ResultBackend:   appconfig.Rabbit.Url,
		ResultsExpireIn: appconfig.Rabbit.Expire,
		AMQP: &config.AMQPConfig{
			Exchange:      appconfig.Rabbit.Exchange,
			ExchangeType:  appconfig.Rabbit.ExchangeType,
			BindingKey:    appconfig.Rabbit.BiddingKey,
			PrefetchCount: appconfig.Rabbit.PreFetchCount,
		},
	}
	server := machinery.NewServer(
		cnf,
		amqpbroker.New(cnf),
		amqpbackend.New(cnf),
		eagerlock.New(),
	)
	myWorker = &worker{
		server: server,
		tasks:  map[string]apptask.TaskInterface{},
	}
	return myWorker
}

func (w *worker) errorHandler(err error) {
	log.ERROR.Println("I am an error handler:", err)
}
func (w *worker) preTaskHandler(signature *tasks.Signature) {
	taskState := elastic_collections.TaskState{
		TaskUUID:   signature.UUID,
		TaskName:   signature.Name,
		Args:       elastic_collections.NewTaskArgs(signature.Args),
		Results:    []elastic_collections.TaskResult{},
		State:      "STARTING",
		RoutingKey: signature.RoutingKey,
		CreatedAt:  *signature.ETA,
		Error:      "",
	}
	clients.AppRequest.Post("tasks/add", nil, taskState)
	log.INFO.Println("I am a start of task handler for:", signature.UUID)
}

func (w *worker) postTaskHandler(signature *tasks.Signature) {
	clients.AppRequest.Put("tasks/status/"+signature.UUID, nil, nil)
	log.INFO.Println("I am an end of task handler for:", signature.UUID)
}

//	func (w *worker) GetStatusById(id string) {
//		state, _ := w.
//	}
func (w *worker) RegisterAll() {
	for _, t := range w.tasks {
		t.Register()
	}
}
func (w *worker) AddTask(newTask apptask.TaskInterface) {
	newTask.SetServer(w.server)
	w.tasks[newTask.GetName()] = newTask
}
func (w *worker) AddTaskAndRegister(newTask apptask.TaskInterface) {
	newTask.SetServer(w.server)
	w.tasks[newTask.GetName()] = newTask
	newTask.Register()
}
func (w *worker) Launch() error {
	cleanup, err := tracers.SetupTracer(appconfig.Rabbit.ConsumerTag)
	if err != nil {
		log.FATAL.Fatalln("Unable to instantiate a tracer:", err)
	}
	defer cleanup()
	wor := w.server.NewWorker(appconfig.Rabbit.ConsumerTag, 0)
	wor.SetPostTaskHandler(w.postTaskHandler)
	wor.SetErrorHandler(w.errorHandler)
	wor.SetPreTaskHandler(w.preTaskHandler)
	return wor.Launch()
}

func Launch() error {
	return myWorker.Launch()
}

func (w *worker) TaskNameList() []string {
	return w.server.GetRegisteredTaskNames()
}
func (w *worker) Test() error {
	cleanup, err := tracers.SetupTracer("sender")
	if err != nil {
		log.FATAL.Fatalln("Unable to instantiate a tracer:", err)
	}
	defer cleanup()
	for name, taskInterface := range w.tasks {
		log.INFO.Println("Sent Test for " + name)
		taskInterface.Test()
	}
	return nil
}

func Test() error {
	return myWorker.Test()
}
