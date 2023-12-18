package repository

import (
	"go-fiber-api/src/models"
)

type Repository[CType models.MongoCollections | models.ElasticCollections] interface {
	Get(id string) (*CType, *models.MyError)
	GetByFirst(key string, value any) (*CType, *models.MyError)
	List(schema models.ListRequest) (*models.ListResponse[CType], *models.MyError)
	Add(schema CType) (*CType, *models.MyError)
	Delete(id string) (bool, *models.MyError)
	Update(id string, schema CType) (*CType, *models.MyError)
	UpdateField(id string, field string, value any) (*CType, *models.MyError)
}
