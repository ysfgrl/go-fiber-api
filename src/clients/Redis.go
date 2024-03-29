package clients

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/ysfgrl/go-fiber-api/src/config"
)

func initRedis(ctx context.Context) (*redis.Client, error) {
	opt, err := redis.ParseURL(config.AppConf.Redis.Url)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opt)
	return client, nil
}

//func (cache *redisHelper) Ping() error {
//	if _, err := cache.client.Ping(cache.ctx).Result(); err != nil {
//		return err
//	}
//	if err := cache.client.Set(cache.ctx, "test", "test redis", 0); err.Err() != nil {
//		return err.Err()
//	}
//	return nil
//}
//
//func (cache *redisHelper) Set(key string, value interface{}) *response2.Error {
//	p, err := json.Marshal(value)
//	if err != nil {
//		return response.GetError(err)
//	}
//	if err := cache.client.Set(cache.ctx, key, p, time.Second); err != nil {
//		return response.GetError(err.Err())
//	}
//	return nil
//}
//
//func (cache *redisHelper) SetTempFile(key string, file response2.UploadedFile) *response2.Error {
//
//	p, err := json.Marshal(file)
//	if err != nil {
//		return response.GetError(err)
//	}
//	if err1 := cache.client.Set(cache.ctx, key, p, time.Second*30); err1.Err() != nil {
//		return response.GetError(err1.Err())
//	}
//	return nil
//}
//func (cache *redisHelper) PublishResource(resource mongo_collections.ResourceListItem) {
//	cache.client.Publish(cache.ctx, "channel1", resource)
//}
