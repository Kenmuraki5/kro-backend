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

type DynamoDBConsoleRepository struct {
	Client *dynamodb.Client
}

func NewDynamoDBConsoleRepository(client *dynamodb.Client) *DynamoDBConsoleRepository {
	return &DynamoDBConsoleRepository{Client: client}
}

func (repo *DynamoDBConsoleRepository) GetAllConsoles() ([]*entity.Console, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Consoles"),
	}
	result, err := repo.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan DynamoDB table: %v", err)
	}
	var consoles []*entity.Console
	for _, item := range result.Items {
		var console entity.Console
		err := attributevalue.UnmarshalMap(item, &console)
		if err != nil {
			return nil, err
		}
		consoles = append(consoles, &console)
	}
	fmt.Println(consoles)
	return consoles, nil
}

func (repo *DynamoDBConsoleRepository) AddConsole(console restmodel.Console) (*restmodel.Console, error) {
	item, err := attributevalue.MarshalMap(console)
	if err != nil {
		return nil, err
	}

	item["Id"] = &types.AttributeValueMemberS{Value: uuid.NewString()}

	item["ReleaseDate"] = &types.AttributeValueMemberS{
		Value: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}

	fmt.Print(item)

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Consoles"),
		Item:      item,
	})
	if err != nil {
		fmt.Printf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}
	return &console, nil
}

// update
func (repo *DynamoDBConsoleRepository) UpdateConsole(updatedConsole entity.Console) (*entity.Console, error) {
	if updatedConsole.Id == "" {
		return nil, fmt.Errorf("cannot update console without a valid ID")
	}

	item, err := attributevalue.MarshalMap(updatedConsole)
	if err != nil {
		return nil, err
	}

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Consoles"),
		Item:      item,
	})
	if err != nil {
		fmt.Printf("Couldn't update item in table. Here's why: %v\n", err)
		return nil, err
	}

	return &updatedConsole, nil
}

func (repo *DynamoDBConsoleRepository) ReleaseStockConsole(order entity.Order) error {
	key, err := attributevalue.MarshalMap(map[string]string{
		"Id": order.ProductId,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal key: %v", err)
	}

	update := &types.Update{
		TableName:        aws.String("Consoles"),
		Key:              key,
		UpdateExpression: aws.String("SET Stock = Stock + :quantity"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":quantity": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", order.Quantity)},
		},
	}

	_, err = repo.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String("Consoles"),
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
func (repo *DynamoDBConsoleRepository) DeleteConsole(id string) error {
	if id == "" {
		return fmt.Errorf("cannot delete console without a valid ID")
	}

	keyCondition := map[string]types.AttributeValue{
		"Id": &types.AttributeValueMemberS{Value: id},
	}

	_, err := repo.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String("Consoles"),
		Key:       keyCondition,
	})
	if err != nil {
		fmt.Printf("Couldn't delete item from table. Here's why: %v\n", err)
		return fmt.Errorf("failed to delete console: %v", err)
	}

	return err
}
