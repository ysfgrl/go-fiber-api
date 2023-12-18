package services

import (
	"context"
	"go-fiber-api/src/models"
)

type Service[CType models.MongoCollections] interface {
	Add(id string) (*CType, *models.MyError)
}

type BaseService struct {
	ctx context.Context
}
