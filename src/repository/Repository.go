package repository

import (
	"context"
	"github.com/ysfgrl/go-fiber-api/src/clients"
	"github.com/ysfgrl/go-fiber-api/src/models"
	"github.com/ysfgrl/go-fiber-api/src/repository/user_repository"
)

var (
	UserRepo *user_repository.UserRepository = nil
)

func init() {
	UserRepo = user_repository.NewUserRepo(clients.GetCollection("users"))
}

type Repository[CType MongoCollections | ElasticCollections] interface {
	Get(ctx context.Context, id string) (*CType, *models.Error)
	GetByFirst(ctx context.Context, key string, value any) (*CType, *models.Error)
	List(ctx context.Context, schema models.ListRequest) ([]CType, *models.Error)
	Add(ctx context.Context, schema CType) (*CType, *models.Error)
	Delete(ctx context.Context, id string) (bool, *models.Error)
	Update(ctx context.Context, id string, schema CType) (*CType, *models.Error)
	UpdateField(ctx context.Context, id string, field string, value any) (*CType, *models.Error)
}
