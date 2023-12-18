package elastic_repository

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"go-fiber-api/src/models"
	"go-fiber-api/src/models/elastic_collections"
	"go-fiber-api/src/repository"
)

type taskStateRepository[CType models.ElasticCollections] struct {
	ElasticRepository[CType]
}

func NewTaskStateRepo(index string, client *elasticsearch.TypedClient) repository.Repository[elastic_collections.TaskState] {
	return &taskStateRepository[elastic_collections.TaskState]{
		ElasticRepository[elastic_collections.TaskState]{
			client: client,
			index:  index,
			ctx:    context.TODO(),
		},
	}
}
