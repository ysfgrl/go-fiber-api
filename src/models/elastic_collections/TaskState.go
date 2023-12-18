package elastic_collections

import (
	"github.com/RichardKnop/machinery/v2/tasks"
	"time"
)

type TaskArg struct {
	Name  string      `bson:"name"`
	Type  string      `bson:"type"`
	Value interface{} `bson:"value"`
}

func NewTaskArgs(args []tasks.Arg) []TaskArg {
	var result []TaskArg
	for _, arg := range args {
		result = append(result, TaskArg{
			Name:  arg.Name,
			Type:  arg.Type,
			Value: arg.Value,
		})
	}
	return result
}

type TaskResult struct {
	Type  string      `bson:"type"`
	Value interface{} `bson:"value"`
}

func NewTaskResult(args []tasks.TaskResult) []TaskResult {
	var result []TaskResult
	for _, arg := range args {
		result = append(result, TaskResult{
			Type:  arg.Type,
			Value: arg.Value,
		})
	}
	return result
}

type TaskState struct {
	TaskUUID   string       `bson:"taskUUID"`
	TaskName   string       `bson:"taskName"`
	State      string       `bson:"state"`
	Results    []TaskResult `bson:"results"`
	Error      string       `bson:"error"`
	CreatedAt  time.Time    `bson:"createdAt"`
	RoutingKey string       `bson:"routingKey"`
	Args       []TaskArg    `bson:"args"`
}
