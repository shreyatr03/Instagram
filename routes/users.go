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

// to create a new user
func UsersHandler(response http.ResponseWriter, request *http.Request) {

	switch request.Method {

	// if method is post
	case http.MethodPost:
		var user models.User

		// read request body as user
		json.NewDecoder(request.Body).Decode(&user)

		// fetch user collection
		var instaDatabase = helper.ConnectDB().Database("instadb")
		var usersCollection = instaDatabase.Collection("user")

		// create new user
		userResult, err := usersCollection.InsertOne(context.TODO(), user)

		// if error in creation
		if err != nil {
			log.Fatal("Cannot create user: ", err)
		}

		// send back
		fmt.Println(userResult.InsertedID)
		json.NewEncoder(response).Encode(userResult)

	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// for fetching user by id
func UserHandler(response http.ResponseWriter, request *http.Request) {

	switch request.Method {

	// if method is get
	case http.MethodGet:
		var user models.User

		// extracting id from url
		id := path.Base(request.RequestURI)

		// fetching user collection
		var instaDatabase = helper.ConnectDB().Database("instadb")
		var usersCollection = instaDatabase.Collection("user")

		// reading user id as objectid
		docID, err := primitive.ObjectIDFromHex(id)

		fmt.Println("UserId of object to be fetced: ", docID)

		// if error in conversion
		if err != nil {
			log.Println("Invalid id: ", err)
		}

		// finding the user by id
		err = usersCollection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&user)

		// if user not found
		if err != nil {
			log.Fatal("User not found: ", err)
		}

		// send back user details as response
		json.NewEncoder(response).Encode(user)

	default:
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
