package repository

import (
	"context"
	"github.com/ysfgrl/go-fiber-api/src/models"
	"github.com/ysfgrl/go-fiber-api/src/pkg/response"
	"github.com/ysfgrl/go-fiber-api/src/repository/user_repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCollections interface {
	user_repository.User
}

type MongoRepository[CType MongoCollections] struct {
	Collection *mongo.Collection
}

func (repo *MongoRepository[CType]) Get(ctx context.Context, id string) (*CType, *models.Error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	return repo.GetByFirst(ctx, "_id", obId)
}

func (repo *MongoRepository[CType]) GetByFirst(ctx context.Context, key string, value any) (*CType, *models.Error) {
	query := bson.M{key: value}
	var user CType
	res := repo.Collection.FindOne(ctx, query)
	if err := res.Decode(&user); err != nil {
		return nil, response.GetError(err)
	}
	return &user, nil
}

func (repo *MongoRepository[CType]) List(ctx context.Context, schema models.ListRequest) ([]CType, *models.Error) {
	opt := options.FindOptions{}
	opt.SetLimit(int64(schema.PageSize))
	opt.SetSkip(int64(schema.Page - 1))
	//opt.SetSort(bson.M{"_id": -1})

	//query := bson.M{"eventDateTime":bson.M{"$gte": schema.Gte, "$lt":schema.Lte}}
	query := bson.M{}

	cursor, err := repo.Collection.Find(ctx, query, &opt)
	if err != nil {
		return nil, response.GetError(err)
	}

	defer cursor.Close(ctx)

	var list []CType

	for cursor.Next(ctx) {
		var item CType
		err := cursor.Decode(&item)
		if err != nil {
			return nil, response.GetError(err)
		}
		list = append(list, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, response.GetError(err)
	}

	if len(list) == 0 {
		list = []CType{}
	}
	return list, nil
}

func (repo *MongoRepository[CType]) Add(ctx context.Context, schema CType) (*CType, *models.Error) {
	res, err := repo.Collection.InsertOne(ctx, schema)
	if err != nil {
		return nil, response.GetError(err)
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return repo.Get(ctx, id)
}

func (repo *MongoRepository[CType]) Update(ctx context.Context, id string, schema CType) (*CType, *models.Error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	_, err := repo.Collection.ReplaceOne(
		ctx,
		bson.M{"_id": obId},
		schema,
	)
	if err != nil {
		return nil, response.GetError(err)
	}
	return repo.Get(ctx, id)
}

func (repo *MongoRepository[CType]) UpdateField(ctx context.Context, id string, field string, value any) (*CType, *models.Error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	_, err := repo.Collection.UpdateByID(
		ctx,
		bson.M{"_id": obId},
		bson.D{{"$set", bson.D{{field, value}}}},
	)
	if err != nil {
		return nil, response.GetError(err)
	}
	return repo.Get(ctx, id)
}

func (repo *MongoRepository[CType]) Delete(ctx context.Context, id string) (bool, *models.Error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	_, err := repo.Collection.DeleteOne(
		ctx,
		bson.M{"_id": obId},
	)
	if err != nil {
		return false, response.GetError(err)
	}
	return true, nil
}
