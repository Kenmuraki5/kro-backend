package s3

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Service interface {
	UploadFile(reader io.Reader, token string) (string, error)
}

type S3Uploader struct {
	Endpoint string
}

func NewS3Uploader() *S3Uploader {
	return &S3Uploader{Endpoint: "http://localhost:4566"}
}

func (s *S3Uploader) UploadFile(reader io.Reader, objKey string) (string, error) {
	awsRegion := "us-east-1"

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if s.Endpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           s.Endpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	bucketName := "kro-gamestore"

	objectKey := "kro-gameStore-" + objKey + ".png"
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   reader,
	})
	if err != nil {
		log.Fatalf("Error uploading picture: %v", err)
		return "", err
	}

	endpoint := fmt.Sprintf("%s/%s/%s", s.Endpoint, bucketName, objectKey)
	log.Printf("Picture uploaded successfully to %s", endpoint)

	return endpoint, nil
}
