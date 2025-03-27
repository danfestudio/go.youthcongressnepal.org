package connect

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var members *mongo.Collection
var err error

// DB initializes the MongoDB connection and returns the client and collection
func DB() (*mongo.Client, *mongo.Collection) {
	if client == nil {
		// MongoDB connection URI
		uri := "mongodb+srv://chetanbudathoki:t4l0EkrUoWHCs03X@youthcongressnepal.h0q3w.mongodb.net/main?retryWrites=true&w=majority&appName=youthcongressnepal"

		// Set up MongoDB client options
		clientOptions := options.Client().ApplyURI(uri)

		// Connect to MongoDB
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		// Ping the database to verify the connection
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatalf("Failed to ping MongoDB: %v", err)
		}

		fmt.Println("Successfully connected to MongoDB!")

		// Set the MongoDB collection
		members = client.Database("cb").Collection("members")
	}

	// Return the client and collection for use in other parts of the app
	return client, members
}
