package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/shreyatr03/Instagram/helper"
	"github.com/shreyatr03/Instagram/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostsHandlers(response http.ResponseWriter, request *http.Request) {

	switch request.Method {

	case http.MethodPost:
		var post models.Post
		var user models.User

		json.NewDecoder(request.Body).Decode(&post)

		uid := post.UserId
		log.Println(uid)

		var instaDatabase = helper.ConnectDB().Database("instadb")
		var usersCollection = instaDatabase.Collection("user")
		err := usersCollection.FindOne(context.TODO(), bson.M{"_id": uid}).Decode(&user)

		if err != nil {
			log.Fatal("hi")
		}

		var postsCollection = instaDatabase.Collection("post")
		post.TimeStamp = primitive.Timestamp{T: uint32(time.Now().Unix())}

		postResult, err1 := postsCollection.InsertOne(context.TODO(), post)

		if err1 != nil {
			log.Fatal(err)
		}

		fmt.Println(postResult.InsertedID)
		json.NewEncoder(response).Encode(postResult)

	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostHandlers(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		var post models.Post
		id := path.Base(request.RequestURI)

		var instaDatabase = helper.ConnectDB().Database("instadb")
		var postsCollection = instaDatabase.Collection("post")

		docID, err := primitive.ObjectIDFromHex(id)
		fmt.Println(docID)
		if err != nil {
			log.Println("Invalid id")
		}

		err = postsCollection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&post)
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(response).Encode(post)

	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
