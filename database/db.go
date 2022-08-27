package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const mongoDbUri = "mongodb://localhost:27017/"

var client *mongo.Client

func GetClient() *mongo.Client {

	var err error

	clientOptions := options.Client().ApplyURI(mongoDbUri)
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	return client
}

func CloseClient() {

	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
