package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type Post struct {
	Id        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string              `json:"caption" bson:"caption"`
	Image_URL string              `json:"image,omitempty" bson:"image,omitempty"`
	TimeStamp primitive.Timestamp `json:"timestamp" bson:"timestamp"`
}

type User struct {
	Id       int    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name" `
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Posts    []Post `json:"posts" bson:"posts"`
}
