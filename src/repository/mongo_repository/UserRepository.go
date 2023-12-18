package mongo_repository

import (
	"context"
	"go-fiber-api/src/models"
	"go-fiber-api/src/models/mongo_collections"
	"go-fiber-api/src/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository[CType models.MongoCollections] struct {
	MongoRepository[CType]
}

func NewUserRepo(collection *mongo.Collection) repository.Repository[mongo_collections.UserListItem] {
	return &userRepository[mongo_collections.UserListItem]{
		MongoRepository: MongoRepository[mongo_collections.UserListItem]{
			collection,
			context.TODO(),
		},
	}
}
