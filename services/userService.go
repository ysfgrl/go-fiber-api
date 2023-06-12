package services

import (
	"context"
	"fmt"
	"go-fiber-api/helpers"
	"go-fiber-api/models"
	"go-fiber-api/myUtils/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mime/multipart"
	"time"
)

type UserService struct {
	Collection *mongo.Collection
	Cache      *helpers.RedisHelper
	Minio      *helpers.MinioHelper
	Ctx        context.Context
}

func (userService *UserService) AddUser(user models.UserListItem) (*models.UserListItem, *models.MyError) {
	res, err := userService.Collection.InsertOne(userService.Ctx, user)
	if err != nil {
		return nil, response.GetError(err)
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return userService.GetUser(id)
}

func (userService *UserService) GetUser(id string) (*models.UserListItem, *models.MyError) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}
	user := models.UserListItem{}
	res := userService.Collection.FindOne(userService.Ctx, query)
	if err := res.Decode(&user); err != nil {
		return nil, response.GetError(err)
	}
	return &user, nil
}

func (userService *UserService) GetByUserName(username string) (*models.UserListItem, *models.MyError) {
	query := bson.M{"userName": username}
	user := models.UserListItem{}
	res := userService.Collection.FindOne(userService.Ctx, query)
	if err := res.Decode(&user); err != nil {
		return nil, response.GetError(err)
	}
	return &user, nil
}

func (userService *UserService) GetList(schema models.ListRequest) (*models.ListResponse, *models.MyError) {

	opt := options.FindOptions{}
	opt.SetLimit(int64(schema.PageSize))
	opt.SetSkip(int64(schema.Page - 1))
	//opt.SetSort(bson.M{"_id": -1})

	//query := bson.M{"eventDateTime":bson.M{"$gte": schema.Gte, "$lt":schema.Lte}}
	query := bson.M{}

	cursor, err := userService.Collection.Find(userService.Ctx, query, &opt)
	if err != nil {
		return nil, response.GetError(err)
	}

	defer cursor.Close(userService.Ctx)

	var list []models.UserListItem

	for cursor.Next(userService.Ctx) {
		post := models.UserListItem{}
		err := cursor.Decode(&post)
		if err != nil {
			return nil, response.GetError(err)
		}
		list = append(list, post)
	}

	if err := cursor.Err(); err != nil {
		return nil, response.GetError(err)
	}

	if len(list) == 0 {
		list = []models.UserListItem{}
	}
	return &models.ListResponse{
		List:     list,
		Page:     schema.Page,
		PageSize: schema.PageSize,
		Total:    len(list),
	}, nil
}

func (userService *UserService) UploadProfile(file *multipart.FileHeader) (string, *models.MyError) {

	info, err := userService.Minio.PutHeaderObject("temp", file)
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
	ctx := context.TODO()
	if err := userService.Cache.SetTempFile(ctx, key, result); err != nil {
		return "", err
	}
	return "minio://" + key, nil

}
