package routes

import (
	"context"
	"log"

	"github.com/danfelab/youthcongressnepal/connect"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

var storageClient *connect.MinioClient // Adjusted to match your Storage return type

func init() {
    var err error
    storageClient, err = connect.Storage()
    if err != nil {
        log.Fatalf("Failed to initialize MinIO client: %v", err)
    }
}

// Assets serves files from MinIO
func Assets(c *fiber.Ctx) error {
    fileName := c.Query("file") // Since route is "/assets/", use query param
    if fileName == "" {
        return c.Status(400).SendString("Missing file parameter")
    }

    log.Printf("Requesting file: %s", fileName)

    object, err := storageClient.GetObject(context.Background(), connect.BucketName, fileName, minio.GetObjectOptions{})
    if err != nil {
        log.Printf("Error fetching %s: %v", fileName, err)
        return c.Status(404).SendString("Image not found")
    }
    defer object.Close()

    c.Set("Content-Type", "image/png")
    c.Set("Content-Disposition", "inline")
    return c.SendStream(object)
}