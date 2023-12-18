package helpers

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go-fiber-api/src/config"
	response2 "go-fiber-api/src/models"
	"go-fiber-api/src/models/mongo_collections"
	"go-fiber-api/src/utils/response"
	"time"
)

type redisHelper struct {
	BaseHelper[redis.Client]
}

func NewRedisHelper(ctx context.Context) (*redisHelper, *response2.MyError) {
	opt, err := redis.ParseURL(config.Redis.Url)
	if err != nil {
		return nil, response.GetError(err)
	}
	client := redis.NewClient(opt)

	return &redisHelper{
		BaseHelper[redis.Client]{
			client: client,
			ctx:    context.TODO(),
		},
	}, nil
}

func (cache *redisHelper) Ping() error {
	if _, err := cache.client.Ping(cache.ctx).Result(); err != nil {
		return err
	}
	if err := cache.client.Set(cache.ctx, "test", "test redis", 0); err.Err() != nil {
		return err.Err()
	}
	return nil
}

func (cache *redisHelper) Set(key string, value interface{}) *response2.MyError {
	p, err := json.Marshal(value)
	if err != nil {
		return response.GetError(err)
	}
	if err := cache.client.Set(cache.ctx, key, p, time.Second); err != nil {
		return response.GetError(err.Err())
	}
	return nil
}

func (cache *redisHelper) SetTempFile(key string, file response2.UploadedFile) *response2.MyError {

	p, err := json.Marshal(file)
	if err != nil {
		return response.GetError(err)
	}
	if err1 := cache.client.Set(cache.ctx, key, p, time.Second*30); err1.Err() != nil {
		return response.GetError(err1.Err())
	}
	return nil
}
func (cache *redisHelper) PublishResource(resource mongo_collections.ResourceListItem) {
	cache.client.Publish(cache.ctx, "channel1", resource)
}
