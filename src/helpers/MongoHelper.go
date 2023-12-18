package helpers

import (
	"context"
	"go-fiber-api/src/config"
	response2 "go-fiber-api/src/models"
	"go-fiber-api/src/utils/response"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoHelper struct {
	BaseHelper[mongo.Client]
}

func NewMongoHelper(ctx context.Context) (*mongoHelper, *response2.MyError) {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.Mongo.Url))
	if err != nil {
		return nil, response.GetError(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, response.GetError(err)
	}
	return &mongoHelper{
		BaseHelper[mongo.Client]{
			client: client,
			ctx:    context.TODO(),
		},
	}, nil
}

func (m *mongoHelper) GetCollection(name string) *mongo.Collection {
	return m.client.Database(config.Mongo.Db).Collection(name)
}
