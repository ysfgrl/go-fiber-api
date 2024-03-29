package clients

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"github.com/ysfgrl/go-fiber-api/src/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoDb *mongo.Client              = nil
	Elastic *elasticsearch.TypedClient = nil
	MiniIO  *minio.Client              = nil
	Redis   *redis.Client              = nil
)

func GetCollection(name string) *mongo.Collection {
	return MongoDb.Database(config.AppConf.Mongo.Db).Collection(name)
}

func init() {
	ctx := context.TODO()

	var err error
	MongoDb, err = initMongo(ctx)
	if err != nil {
		panic(err.Error())
	}

	Elastic, err = initElastic(ctx)
	if err != nil {
		panic(err.Error())
	}

	MiniIO, err = initMinio(ctx)
	if err != nil {
		panic(err.Error())
	}
	Redis, err = initRedis(ctx)
	if err != nil {
		panic(err.Error())
	}
}
