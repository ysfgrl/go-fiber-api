package mongo_repository

import (
	"context"
	"go-fiber-api/src/models"
	"go-fiber-api/src/models/mongo_collections"
	"go-fiber-api/src/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type resourceRepository[CType models.MongoCollections] struct {
	MongoRepository[CType]
}

func NewResourceRepo(collection *mongo.Collection) repository.Repository[mongo_collections.ResourceListItem] {
	return &resourceRepository[mongo_collections.ResourceListItem]{
		MongoRepository: MongoRepository[mongo_collections.ResourceListItem]{
			collection,
			context.TODO(),
		},
	}
}
