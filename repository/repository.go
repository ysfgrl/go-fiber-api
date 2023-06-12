package repository

import (
	"context"
	"go-fiber-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoRepository struct {
	client *mongo.Client
	db     *mongo.Database
	Config *config.Config
	Error  error
}

func (mongoRepository *MongoRepository) Connect() error {
	mongoRepository.client, mongoRepository.Error = mongo.NewClient(options.Client().ApplyURI(mongoRepository.Config.MongoUrl))
	if mongoRepository.Error != nil {
		return mongoRepository.Error
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mongoRepository.Error = mongoRepository.client.Connect(ctx)
	mongoRepository.db = mongoRepository.client.Database(mongoRepository.Config.MongoDb)

	if mongoRepository.client != nil {
		return mongoRepository.Error
	}
	return nil
}

func (mongoRepository *MongoRepository) GetCollection(name string) *mongo.Collection {
	return mongoRepository.db.Collection(name)
}
