package mongo_repository

import (
	"context"
	"go-fiber-api/src/models"
	"go-fiber-api/src/utils/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository[CType models.MongoCollections] struct {
	collection *mongo.Collection
	ctx        context.Context
}

func (repo *MongoRepository[CType]) Get(id string) (*CType, *models.MyError) {
	obId, _ := primitive.ObjectIDFromHex(id)
	return repo.GetByFirst("_id", obId)
}

func (repo *MongoRepository[CType]) GetByFirst(key string, value any) (*CType, *models.MyError) {
	query := bson.M{key: value}
	var user CType
	res := repo.collection.FindOne(repo.ctx, query)
	if err := res.Decode(&user); err != nil {
		return nil, response.GetError(err)
	}
	return &user, nil
}

func (repo *MongoRepository[CType]) List(schema models.ListRequest) (*models.ListResponse[CType], *models.MyError) {
	opt := options.FindOptions{}
	opt.SetLimit(int64(schema.PageSize))
	opt.SetSkip(int64(schema.Page - 1))
	//opt.SetSort(bson.M{"_id": -1})

	//query := bson.M{"eventDateTime":bson.M{"$gte": schema.Gte, "$lt":schema.Lte}}
	query := bson.M{}

	cursor, err := repo.collection.Find(repo.ctx, query, &opt)
	if err != nil {
		return nil, response.GetError(err)
	}

	defer cursor.Close(repo.ctx)

	var list []CType

	for cursor.Next(repo.ctx) {
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
	return &models.ListResponse[CType]{
		List:     list,
		Page:     schema.Page,
		PageSize: schema.PageSize,
		Total:    len(list),
	}, nil
}

func (repo *MongoRepository[CType]) Add(schema CType) (*CType, *models.MyError) {
	res, err := repo.collection.InsertOne(repo.ctx, schema)
	if err != nil {
		return nil, response.GetError(err)
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return repo.Get(id)
}

func (repo *MongoRepository[CType]) Update(id string, schema CType) (*CType, *models.MyError) {
	obId, _ := primitive.ObjectIDFromHex(id)
	_, err := repo.collection.ReplaceOne(
		repo.ctx,
		bson.M{"_id": obId},
		schema,
	)
	if err != nil {
		return nil, response.GetError(err)
	}
	return repo.Get(id)
}

func (repo *MongoRepository[CType]) UpdateField(id string, field string, value any) (*CType, *models.MyError) {
	obId, _ := primitive.ObjectIDFromHex(id)
	_, err := repo.collection.UpdateByID(
		repo.ctx,
		bson.M{"_id": obId},
		bson.D{{"$set", bson.D{{field, value}}}},
	)
	if err != nil {
		return nil, response.GetError(err)
	}
	return repo.Get(id)
}

func (repo *MongoRepository[CType]) Delete(id string) (bool, *models.MyError) {
	obId, _ := primitive.ObjectIDFromHex(id)
	_, err := repo.collection.DeleteOne(
		repo.ctx,
		bson.M{"_id": obId},
	)
	if err != nil {
		return false, response.GetError(err)
	}
	return true, nil
}
