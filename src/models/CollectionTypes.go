package models

import (
	"go-fiber-api/src/models/elastic_collections"
	"go-fiber-api/src/models/mongo_collections"
)

type MongoCollections interface {
	mongo_collections.UserListItem | mongo_collections.ResourceListItem
}

type ElasticCollections interface {
	elastic_collections.TaskState
}
