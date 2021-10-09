package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	clientOptions := options.Client().
		ApplyURI("mongodb+srv://admin:admin@cluster0.hvyie.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	// instaDatabase := client.Database("instadb")
	// usersCollection := instaDatabase.Collection("user")
	// // postsCollection := instaDatabase.Collection("posts")

	// userResult, err := usersCollection.InsertOne(ctx, bson.D{
	// 	{Key: "name", Value: "Shreya"},
	// 	{Key: "email", Value: "shreyatr@outlook.com"},
	// 	{Key: "password", Value: "pass"},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(userResult.InsertedID)

	http.HandleFunc("/users/{id}", usersHandlers)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
