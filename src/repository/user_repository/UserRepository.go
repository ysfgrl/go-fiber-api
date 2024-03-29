package user_repository

import (
	"github.com/ysfgrl/go-fiber-api/src/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	repository.MongoRepository[User]
}

func NewUserRepo(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		MongoRepository: repository.MongoRepository[User]{
			Collection: collection,
		},
	}
}

func NewUserRepoImpl(collection *mongo.Collection) repository.Repository[User] {
	return &UserRepository{
		MongoRepository: repository.MongoRepository[User]{
			Collection: collection,
		},
	}
}
