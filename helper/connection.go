package helper

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.hvyie.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
