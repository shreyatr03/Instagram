package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type Post struct {
	Id        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string              `json:"caption,omitempty" bson:"caption,omitempty"`
	Image_URL string              `json:"image" bson:"image,omitempty"`
	TimeStamp primitive.Timestamp `json:"timestamp" bson:"timestamp,omitempty"`
	User      *User               `json:"user" bson:"user,omitempty"`
}

type User struct {
	Id       int    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty" `
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}
