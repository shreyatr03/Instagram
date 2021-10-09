package helper

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.hvyie.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

//user := models.User{}
// clientOptions := options.Client().
// 	ApplyURI("mongodb+srv://admin:admin@cluster0.hvyie.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()
// client, err := mongo.Connect(ctx, clientOptions)
// if err != nil {
// 	log.Fatal(err)
// }
