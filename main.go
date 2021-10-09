package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shreyatr03/Instagram/helper"
	"go.mongodb.org/mongo-driver/bson"
)

func usersHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Serve the resource.
	case http.MethodPost:

	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	//user := models.User{}
	// clientOptions := options.Client().
	// 	ApplyURI("mongodb+srv://admin:admin@cluster0.hvyie.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	client, ctx := helper.ConnectDB()
	defer client.Disconnect(ctx)

	instaDatabase := client.Database("instadb")
	usersCollection := instaDatabase.Collection("user")
	// postsCollection := instaDatabase.Collection("posts")

	userResult, err := usersCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Shreya_1"},
		{Key: "email", Value: "shreyatr@outlook.com"},
		{Key: "password", Value: "pass"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userResult.InsertedID)

	http.HandleFunc("/users/{id}", usersHandlers)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
