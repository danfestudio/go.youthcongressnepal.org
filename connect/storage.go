package connect

import (
	"context"
	"errors"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIO connection configuration
const (
    endpoint        = "94.136.185.141:9000" // Replace with your MinIO server endpoint
    accessKeyID     = "HTkDFKvkz7grU3qIPgDk"     // Replace with your access key
    secretAccessKey = "cJwK85jmHarjY5K78SuVcm5qVIswKRcpubGHqC0z"     // Replace with your secret key
    useSSL          = false            // Set to true if using SSL/TLS
	BucketName      = "youthcongressnepal" // Your bucket name
)

func Storage() (*minio.Client, error) {
    // Initialize MinIO client
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        return nil, err
    }

    // Test the connection by checking if the specified bucket exists
    ctx := context.Background()
    exists, err := minioClient.BucketExists(ctx, BucketName)
    if err != nil {
        log.Printf("Failed to verify MinIO connection: %v", err)
        return nil, errors.New("failed to verify MinIO connection: " + err.Error())
    }
    if !exists {
        log.Printf("Bucket %s does not exist", BucketName)
        return nil, errors.New("specified bucket does not exist")
    }

    log.Printf("Successfully connected to MinIO and verified bucket %s", BucketName)
    
    return minioClient, nil    
}