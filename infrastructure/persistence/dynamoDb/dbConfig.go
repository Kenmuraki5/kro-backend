package dynamoDb

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/joho/godotenv"
)

type DynamoDBClient struct {
	Client *dynamodb.Client
}

func NewDynamoDBClient() (*DynamoDBClient, error) {
	godotenv.Load()
	accessKey := os.Getenv("DYNAMO_ACCESSKEY")
	secretAccessKey := os.Getenv("DYNAMO_SECRETACCESSKEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("localhost"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKey,       // ใส่ access key
				SecretAccessKey: secretAccessKey, // ใส่ secret access key
			},
		}),
	)
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBClient{Client: client}, nil
}
