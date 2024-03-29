package models

type UserAddBasic struct {
	Email     string `json:"email" bson:"email" validate:"required"`
	FirstName string `json:"firstName" bson:"firstName" validate:"required"`
	LastName  string `json:"lastName" bson:"lastName" validate:"required"`
	Password  string `json:"password" bson:"password" validate:"required"`
}
