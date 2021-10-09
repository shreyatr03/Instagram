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
		json.NewDecoder(request.Body).Decode(&post)

		var instaDatabase = helper.ConnectDB().Database("instadb")
		var postsCollection = instaDatabase.Collection("post")
		post.TimeStamp = primitive.Timestamp{T: uint32(time.Now().Unix())}

		postResult, err := postsCollection.InsertOne(context.TODO(), post)

		if err != nil {
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
