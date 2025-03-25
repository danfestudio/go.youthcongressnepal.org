package connect

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() {
	// MongoDB connection URI
    uri := "mongodb+srv://chetanbudathoki:HeroBudathoki3579@chetanbudathoki.ko2ln.mongodb.net/cb?retryWrites=true&w=majority&appName=chetanbudathoki"

    // Set up MongoDB client options
    clientOptions := options.Client().ApplyURI(uri)

    // Connect to MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Ping the database to verify the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    fmt.Println("Successfully connected to MongoDB!")

    // Close the connection when done
    defer func() {
        if err = client.Disconnect(ctx); err != nil {
            log.Fatalf("Failed to disconnect from MongoDB: %v", err)
        }
        fmt.Println("Disconnected from MongoDB.")
    }()
}