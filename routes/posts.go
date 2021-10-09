package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shreyatr03/Instagram/helper"
	"github.com/shreyatr03/Instagram/models"
)

func PostsHandlers(response http.ResponseWriter, request *http.Request) {

	switch request.Method {

	case http.MethodPost:
		var post models.Post
		json.NewDecoder(request.Body).Decode(&post)

		var instaDatabase = helper.ConnectDB().Database("instadb")
		var postsCollection = instaDatabase.Collection("post")

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

}
