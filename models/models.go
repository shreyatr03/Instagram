package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// post structure
type Post struct {
	Id        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string              `json:"caption" bson:"caption"`
	Image_URL string              `json:"image,omitempty" bson:"image,omitempty"`
	TimeStamp primitive.Timestamp `json:"timestamp" bson:"timestamp"`
	UserId    primitive.ObjectID  `json:"userId,omitempty" bson:"userId,omitempty"`
}

// user structure
type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name" `
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type Users []User
type Posts []Post
