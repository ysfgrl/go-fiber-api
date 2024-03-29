package user_repository

import (
	"time"
)

type User struct {
	ID         string    `json:"id,omitempty" bson:"_id,omitempty" `
	UserName   string    `json:"userName" bson:"userName" `
	Email      string    `json:"email" bson:"email" `
	FirstName  string    `json:"firstName" bson:"firstName" `
	MiddleName string    `json:"middleName" bson:"middleName" `
	LastName   string    `json:"lastName" bson:"lastName" `
	Password   string    `json:"password" bson:"password" `
	Role       string    `json:"role" bson:"role" `
	Profile    string    `json:"profile" bson:"profile" `
	Active     bool      `json:"active" bson:"active" `
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt" `
	UpdatedAt  time.Time `json:"updatedAt" bson:"UpdatedAt"`
}
