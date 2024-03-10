package dynamoDb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
	Client *dynamodb.Client
}

func NewDynamoDBClient() (*DynamoDBClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBClient{Client: client}, nil
}
