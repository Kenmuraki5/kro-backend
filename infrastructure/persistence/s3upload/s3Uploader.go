package s3

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func S3uploader(c *gin.Context) {
	awsEndpoint := "http://localhost:4566"
	awsRegion := "us-east-1"

	err := c.Request.ParseMultipartForm(10 << 20)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fmt.Println("2")
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
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

	fmt.Println("1")
	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	bucketName := "kro-gamestore"

	var imageUrls []string

	fmt.Println(c.Request.MultipartForm.File)
	for _, files := range c.Request.MultipartForm.File {

		for _, file := range files {
			objectKey := uuid.NewString() + ".png"
			src, err := file.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			defer src.Close()
			_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
				Bucket: &bucketName,
				Key:    &objectKey,
				Body:   src,
			})
			if err != nil {
				log.Fatalf("Error uploading picture: %v", err)
			}

			imageUrls = append(imageUrls, fmt.Sprintf("%s/%s/%s", awsEndpoint, bucketName, objectKey))
		}
	}

	c.JSON(http.StatusOK, gin.H{"imageUrls": imageUrls})
}
