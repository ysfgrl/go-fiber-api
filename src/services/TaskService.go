package services

import (
	"context"
	"github.com/ysfgrl/fibersocket"
	"github.com/ysfgrl/go-fiber-api/src/clients"
	"github.com/ysfgrl/go-fiber-api/src/models"
	"github.com/ysfgrl/go-fiber-api/src/models/elastic_collections"
	"github.com/ysfgrl/go-fiber-api/src/repository"
	"github.com/ysfgrl/go-fiber-api/src/repository/elastic_repository"
)

type TaskService struct {
	repo repository.Repository[elastic_collections.TaskState]
	ctx  context.Context
}

func NewTaskService() TaskService {
	return TaskService{
		elastic_repository.NewTaskStateRepo("task_state", clients.Elastic.GetClient()),
		context.TODO(),
	}
}

func (service *TaskService) Add(schema elastic_collections.TaskState) (*elastic_collections.TaskState, *models.Error) {
	socketServer := fibersocket.GetServerByName("task_server")
	if socketServer != nil {
		//socketServer.Emit(shcema)
	}
	return service.repo.Add(schema)
}

func (service *TaskService) GetList(schema models.ListRequest) (*models.ListResponse[elastic_collections.TaskState], *models.Error) {
	return service.repo.List(schema)
}
