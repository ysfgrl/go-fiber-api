package helpers

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go-fiber-api/models"
	"go-fiber-api/myUtils/response"
	"time"
)

type RedisHelper struct {
	Client *redis.Client
}

func (cache *RedisHelper) Ping(ctx context.Context) error {
	if _, err := cache.Client.Ping(ctx).Result(); err != nil {
		return err
	}
	if err := cache.Client.Set(ctx, "test", "test redis", 0); err.Err() != nil {
		return err.Err()
	}
	return nil
}

func (cache *RedisHelper) Set(ctx context.Context, key string, value interface{}) *models.MyError {
	p, err := json.Marshal(value)
	if err != nil {
		return response.GetError(err)
	}
	if err := cache.Client.Set(ctx, key, p, time.Second); err != nil {
		return response.GetError(err.Err())
	}
	return nil
}

func (cache *RedisHelper) SetTempFile(ctx context.Context, key string, file models.UploadedFile) *models.MyError {

	p, err := json.Marshal(file)
	if err != nil {
		return response.GetError(err)
	}
	if err1 := cache.Client.Set(ctx, key, p, time.Second*30); err1.Err() != nil {
		return response.GetError(err1.Err())
	}
	return nil
}
