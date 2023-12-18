package mongo_collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ResourceListItem struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"-"`
	Url       string             `json:"url" bson:"url" validate:"required,min=3,max=500"`
	LocalUrl  string             `json:"localUrl" bson:"localUrl" validate:"-""`
	Title     string             `json:"title" bson:"title" validate:"required"`
	Type      string             `json:"type" bson:"type" validate:"-"`
	Download  bool               `json:"download" bson:"download" validate:"-"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt" validate:"-"`
}
