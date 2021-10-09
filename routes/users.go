package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shreyatr03/Instagram/helper"
	"github.com/shreyatr03/Instagram/models"
	"go.mongodb.org/mongo-driver/bson"
)

func usersHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		client := helper.ConnectDB()

		defer cancel()

		instaDatabase := client.Database("instadb")
		usersCollection := instaDatabase.Collection("user")
		// // postsCollection := instaDatabase.Collection("posts")

		userResult, err := usersCollection.InsertOne(ctx, bson.D{
			{Key: "name", Value: "Shreya_1"},
			{Key: "email", Value: "shreyatr@outlook.com"},
			{Key: "password", Value: "pass"},
		})

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(userResult.InsertedID)

	case http.MethodPost:
		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		client := helper.ConnectDB()

		defer cancel()

		instaDatabase := client.Database("instadb")
		usersCollection := instaDatabase.Collection("user")
		// // postsCollection := instaDatabase.Collection("posts")

		userResult, err := usersCollection.InsertOne(ctx, user)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(userResult.InsertedID)
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func userHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// client := helper.ConnectDB()

		// defer cancel()

		// instaDatabase := client.Database("instadb")
		// usersCollection := instaDatabase.Collection("user")
		// user := usersCollection.FindOne(context.Background(), bson.M{"_id": bson.ObjectIdHex(`{id}`)})
		// // // postsCollection := instaDatabase.Collection("posts")

		// userResult, err := usersCollection.InsertOne(ctx, bson.D{
		// 	{Key: "name", Value: "Shreya_1"},
		// 	{Key: "email", Value: "shreyatr@outlook.com"},
		// 	{Key: "password", Value: "pass"},
		// })

		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(userResult.InsertedID)

	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
