package gcloudservices

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type ClientUploader struct {
	client     *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func NewClientUploader() *ClientUploader {
	projectID := os.Getenv("GCP_PROJECT_ID")
	bucketName := os.Getenv("GCP_BUCKET_NAME")
	credentialsDir := os.Getenv("GCP_CREDENTIALS_DIR")
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credentialsDir)

	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &ClientUploader{
		client:     client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "",
	}
}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(fileHeader *multipart.FileHeader) (url string, err error) {
	fileBytes, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("Open file: %v", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	fileName := GenUniqueFilename(fileHeader.Filename)
	// Upload an object with storage.Writer.
	wc := c.client.Bucket(c.bucketName).Object(c.uploadPath + fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, fileBytes); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}
	gcp_domain := os.Getenv("GCP_DOMAIN")
	url = gcp_domain + c.bucketName + "/" + c.uploadPath + fileName
	return url, nil
}

// Set upload dir
func (c *ClientUploader) SetDstPath(dst string) {
	c.uploadPath = dst
}

// create uniq file name
func GenUniqueFilename(name string) (uniqueName string) {
	id := uuid.New()
	uniqueName = id.String() + name
	return
}
