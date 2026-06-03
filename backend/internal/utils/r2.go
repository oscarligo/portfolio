package utils

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type R2Uploader struct {
	s3Client   *s3.Client
	bucketName string
	publicURL  string
}

func NewR2Uploader() (*R2Uploader, error) {

	// Env variables
	accountID := os.Getenv("R2_ACCOUNT_ID")
	accessKey := os.Getenv("R2_ACCESS_KEY_ID")
	secretKey := os.Getenv("R2_SECRET_ACCESS_KEY")
	bucket := os.Getenv("R2_BUCKET_NAME")
	publicURL := os.Getenv("R2_PUBLIC_URL")

	if accountID == "" || accessKey == "" || secretKey == "" || bucket == "" {
		return nil, fmt.Errorf("missing critical environment variables for R2 configuration")
	}

	// Load AWS SDK configuration with custom credentials and region for R2
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, fmt.Errorf("error while loading default config: %w", err)
	}
	// Cloudflare R2 endpoint configuration
	r2Endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID)

	// Create an S3 client with the custom endpoint for R2
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = &r2Endpoint
	})

	return &R2Uploader{
		s3Client:   client,
		bucketName: bucket,
		publicURL:  publicURL,
	}, nil
}

// UploadImage uploads an image to R2 and returns the public URL for accessing
func (u *R2Uploader) UploadImage(ctx context.Context, objectKey string, fileReader io.Reader, contentType string) (string, error) {
	_, err := u.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &u.bucketName,
		Key:         &objectKey,
		Body:        fileReader,
		ContentType: &contentType,
	})
	if err != nil {
		return "", fmt.Errorf("error while uploading image to R2: %w", err)
	}

	return fmt.Sprintf("%s/%s", u.publicURL, objectKey), nil
}
