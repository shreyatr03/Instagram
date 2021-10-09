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

}
