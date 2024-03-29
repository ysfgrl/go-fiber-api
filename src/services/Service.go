package services

import (
	"context"
	"github.com/ysfgrl/go-fiber-api/src/models"
	"github.com/ysfgrl/go-fiber-api/src/repository"
)

type Service[CType repository.MongoCollections] interface {
	Add(id string) (*CType, *models.Error)
}

type BaseService struct {
	ctx context.Context
}
