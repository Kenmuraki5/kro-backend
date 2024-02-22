package dynamoDb

import (
	"context"
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	// "github.com/Kenmuraki5/kro-backend.git/domain/repository"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBGameRepository struct {
	Client *dynamodb.Client
}

func NewDynamoDBGameRepository(client *dynamodb.Client) *DynamoDBGameRepository {
	return &DynamoDBGameRepository{Client: client}
}

func (repo *DynamoDBGameRepository) GetAllGames() ([]*entity.Game, error) {
	fmt.Println("Getting all games")
	// Create a scan input to get all items from the table
	input := &dynamodb.ScanInput{
		TableName: aws.String("game"),
	}

	// Perform the Scan operation
	result, err := repo.Client.Scan(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan DynamoDB table: %v", err)
	}

	// Unmarshal the items into Game structs
	games := make([]*entity.Game, len(result.Items))
	for i, item := range result.Items {
		fmt.Println(item)
		var game entity.Game
		if err := attributevalue.UnmarshalMap(item, &game); err != nil {
			return nil, fmt.Errorf("failed to unmarshal DynamoDB item: %v", err)
		}
		games[i] = &game
	}

	return games, nil
}
