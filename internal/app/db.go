package app

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoDB(ctx context.Context) *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	clientDB, e := mongo.Connect(ctx, clientOptions)
	if e != nil {
		fmt.Println(e)
	}

	// Check the connection
	e = clientDB.Ping(context.TODO(), nil)
	if e != nil {
		fmt.Println(e)
	}
	return clientDB
}
