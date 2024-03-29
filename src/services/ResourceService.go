package services

import (
	"context"
	"github.com/ysfgrl/go-fiber-api/src/clients"
	"github.com/ysfgrl/go-fiber-api/src/models"
	"github.com/ysfgrl/go-fiber-api/src/models/mongo_collections"
	"github.com/ysfgrl/go-fiber-api/src/repository"
	"mime/multipart"
	"time"
)

type ResourceService struct {
	repo repository.Repository[mongo_collections.ResourceListItem]
	ctx  context.Context
}

func NewResourceService(repo repository.Repository[mongo_collections.ResourceListItem]) ResourceService {
	return ResourceService{
		repo,
		context.TODO(),
	}
}

func (service *ResourceService) AddResource(schema mongo_collections.ResourceListItem) (*mongo_collections.ResourceListItem, *models.Error) {
	newRes, err := service.repo.Add(schema)
	if err != nil {
		return nil, err
	}
	return newRes, nil
}

func (service *ResourceService) GetList(schema models.ListRequest) (*models.ListResponse[mongo_collections.ResourceListItem], *models.Error) {
	list, err := service.repo.List(schema)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (service *ResourceService) GetResource(id string) (*mongo_collections.ResourceListItem, *models.Error) {
	resource, err := service.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (service *ResourceService) UploadResource(file *multipart.FileHeader) (*mongo_collections.ResourceListItem, *models.Error) {
	info, err := clients.Minio.PutHeaderObject("temp", file)
	if err != nil {
		return nil, err
	}
	url := "minio://" + info.Bucket + "/" + info.Key
	newResource := mongo_collections.ResourceListItem{
		Type:      file.Header["Content-Type"][0],
		Url:       url,
		LocalUrl:  url,
		Title:     info.Key,
		Download:  true,
		CreatedAt: time.Now().UTC(),
	}
	return service.AddResource(newResource)

}
