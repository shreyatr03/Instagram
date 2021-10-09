package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/shreyatr03/Instagram/helper"
	"github.com/shreyatr03/Instagram/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UsersHandlers(response http.ResponseWriter, request *http.Request) {

	switch request.Method {

	case http.MethodPost:
		var user models.User
		json.NewDecoder(request.Body).Decode(&user)

		var instaDatabase = helper.ConnectDB().Database("instadb")
		var usersCollection = instaDatabase.Collection("user")

		userResult, err := usersCollection.InsertOne(context.TODO(), user)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(userResult.InsertedID)
		json.NewEncoder(response).Encode(userResult)

	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UserHandlers(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		var user models.User
		id := path.Base(request.RequestURI)

		var instaDatabase = helper.ConnectDB().Database("instadb")
		var usersCollection = instaDatabase.Collection("user")

		docID, err := primitive.ObjectIDFromHex(id)
		fmt.Println(docID)
		if err != nil {
			log.Println("Invalid id")
		}

		err = usersCollection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&user)
		if err != nil {
			log.Println(err)
		}

		json.NewEncoder(response).Encode(user)
	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
