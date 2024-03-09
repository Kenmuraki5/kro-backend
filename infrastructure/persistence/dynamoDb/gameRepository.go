package dynamoDb

import (
	"context"
	"fmt"
	"time"

	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/google/uuid"

	// "github.com/Kenmuraki5/kro-backend.git/domain/repository"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBGameRepository struct {
	Client *dynamodb.Client
}

func NewDynamoDBGameRepository(client *dynamodb.Client) *DynamoDBGameRepository {
	return &DynamoDBGameRepository{Client: client}
}

func (repo *DynamoDBGameRepository) GetAllGames() ([]*entity.Game, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Games"),
	}
	result, err := repo.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan DynamoDB table: %v", err)
	}
	var games []*entity.Game
	for _, item := range result.Items {
		var game entity.Game
		err := attributevalue.UnmarshalMap(item, &game)
		if err != nil {
			return nil, err
		}
		games = append(games, &game)
	}
	fmt.Println(games)
	return games, nil
}

func (repo *DynamoDBGameRepository) AddGame(game restmodel.Game) (*restmodel.Game, error) {
	item, err := attributevalue.MarshalMap(game)
	item["Id"] = &types.AttributeValueMemberS{Value: uuid.NewString()}
	item["ReleaseDate"] = &types.AttributeValueMemberS{
		Value: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
	fmt.Print(item)
	if err != nil {
		return nil, err
	}

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Games"),
		Item:      item,
	})
	if err != nil {
		fmt.Printf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}
	return &game, nil
}

// update
func (repo *DynamoDBGameRepository) UpdateGame(updatedGame entity.Game) (*entity.Game, error) {
	if updatedGame.Id == "" {
		return nil, fmt.Errorf("cannot update game without a valid ID")
	}

	item, err := attributevalue.MarshalMap(updatedGame)
	if err != nil {
		return nil, err
	}

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Games"),
		Item:      item,
	})
	if err != nil {
		fmt.Printf("Couldn't update item in table. Here's why: %v\n", err)
		return nil, err
	}

	return &updatedGame, nil
}

func (repo *DynamoDBGameRepository) ReleaseStockGame(order entity.Order) error {
	key, err := attributevalue.MarshalMap(map[string]string{
		"Id": order.ProductId,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal key: %v", err)
	}

	update := &types.Update{
		TableName:        aws.String("Games"),
		Key:              key,
		UpdateExpression: aws.String("SET Stock = Stock + :quantity"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":quantity": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", order.Quantity)},
		},
	}

	_, err = repo.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String("Games"),
		Key:                       key,
		UpdateExpression:          update.UpdateExpression,
		ExpressionAttributeValues: update.ExpressionAttributeValues,
		ConditionExpression:       update.ConditionExpression,
	})
	if err != nil {
		fmt.Printf("Failed to update stock in table: %v\n", err)
		return fmt.Errorf("failed to update stock: %v", err)
	}

	return nil
}

// delete
func (repo *DynamoDBGameRepository) DeleteGame(id string) error {
	if id == "" {
		return fmt.Errorf("cannot delete game without a valid ID")
	}

	keyCondition := map[string]types.AttributeValue{
		"Id": &types.AttributeValueMemberS{Value: id},
	}

	_, err := repo.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String("Games"),
		Key:       keyCondition,
	})
	if err != nil {
		fmt.Printf("Couldn't delete item from table. Here's why: %v\n", err)
		return fmt.Errorf("failed to delete game: %v", err)
	}

	return err
}
