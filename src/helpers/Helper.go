package helpers

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"go-fiber-api/src/models"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var Minio *minioHelper = nil
var Mongo *mongoHelper = nil
var Elastic *elasticHelper = nil
var Redis *redisHelper = nil
var AppRequest *httpRequestHelper = nil

type ClientTypes interface {
	minio.Client | mongo.Client | elasticsearch.TypedClient | redis.Client | http.Client
}

type BaseHelper[CType ClientTypes] struct {
	client *CType
	ctx    context.Context
}

func (h *BaseHelper[HType]) GetClient() *HType {
	return h.client
}

func InitHelpers() {
	var err *models.MyError = nil
	if Mongo == nil {
		Mongo, err = NewMongoHelper(context.TODO())
		if err != nil {
			panic(string(err.ToJson()))
		}
	}
	if Redis == nil {
		Redis, err = NewRedisHelper(context.TODO())
		if err != nil {
			panic(err)
		}
	}
	if Minio == nil {
		Minio, err = NewMinioHelper(context.TODO())
		if err != nil {
			panic(err)
		}
	}
	if Elastic == nil {
		Elastic, err = NewElasticHelper(context.TODO())
		if err != nil {
			panic(err)
		}
	}
	if AppRequest == nil {
		AppRequest, err = NewHttpRequestHelper(context.TODO(), "http://localhost:8080/")
		if err != nil {
			panic(err)
		}
	}
}
