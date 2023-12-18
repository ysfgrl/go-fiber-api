package services

import (
	"context"
	"fmt"
	"go-fiber-api/src/helpers"
	"go-fiber-api/src/models"
	"go-fiber-api/src/models/mongo_collections"
	"go-fiber-api/src/repository"
	"mime/multipart"
	"time"
)

type UserService struct {
	repo repository.Repository[mongo_collections.UserListItem]
	ctx  context.Context
}

func NewUserService(repo repository.Repository[mongo_collections.UserListItem]) UserService {
	return UserService{
		repo,
		context.TODO(),
	}
}

func (service *UserService) AddUser(schema mongo_collections.UserListItem) (*mongo_collections.UserListItem, *models.MyError) {
	newUser, err := service.repo.Add(schema)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (service *UserService) GetUser(id string) (*mongo_collections.UserListItem, *models.MyError) {
	user, err := service.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) GetByUserName(username string) (*mongo_collections.UserListItem, *models.MyError) {
	user, err := service.repo.GetByFirst("userName", username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) GetList(schema models.ListRequest) (*models.ListResponse[mongo_collections.UserListItem], *models.MyError) {
	list, err := service.repo.List(schema)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (service *UserService) UploadProfile(file *multipart.FileHeader) (string, *models.MyError) {

	info, err := helpers.Minio.PutHeaderObject("temp", file)
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("temp_%v", time.Now().UnixNano())
	result := models.UploadedFile{
		Bucket: info.Bucket,
		Name:   info.Key,
		Size:   info.Size,
		Type:   file.Header["Content-Type"][0],
	}
	if err := helpers.Redis.SetTempFile(key, result); err != nil {
		return "", err
	}
	return "minio://" + key, nil

}
