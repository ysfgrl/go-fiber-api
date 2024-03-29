package services

import (
	"context"
	"fmt"
	"github.com/ysfgrl/go-fiber-api/src/clients"
	"github.com/ysfgrl/go-fiber-api/src/models"
	"github.com/ysfgrl/go-fiber-api/src/repository"
	"github.com/ysfgrl/go-fiber-api/src/repository/user_repository"
	"mime/multipart"
	"time"
)

type UserService struct {
	repo repository.Repository[user_repository.User]
	ctx  context.Context
}

func NewUserService(repo repository.Repository[user_repository.User]) UserService {
	return UserService{
		repo,
		context.TODO(),
	}
}

func (service *UserService) AddUser(schema user_repository.User) (*user_repository.User, *models.Error) {
	newUser, err := service.repo.Add(schema)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (service *UserService) GetUser(id string) (*user_repository.User, *models.Error) {
	user, err := service.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) GetByUserName(username string) (*user_repository.User, *models.Error) {
	user, err := service.repo.GetByFirst("userName", username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) GetList(schema models.ListRequest) (*models.ListResponse[user_repository.User], *models.Error) {
	list, err := service.repo.List(schema)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (service *UserService) UploadProfile(file *multipart.FileHeader) (string, *models.Error) {

	info, err := clients.Minio.PutHeaderObject("temp", file)
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
	if err := clients.Redis.SetTempFile(key, result); err != nil {
		return "", err
	}
	return "minio://" + key, nil

}
