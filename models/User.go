package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserListItem struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"-"`
	UserName  string             `json:"userName" bson:"userName" validate:"required,min=3,max=12"`
	Email     string             `json:"email" bson:"email" validate:"required"`
	FirstName string             `json:"firstName" bson:"firstName" validate:"required"`
	LastName  string             `json:"lastName" bson:"lastName" validate:"required"`
	Password  string             `json:"password" bson:"password" validate:"required"`
	Role      string             `json:"role" bson:"role" validate:"required"`
	Profile   string             `json:"profile" bson:"profile" validate:"required"`
}
