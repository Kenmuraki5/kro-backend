package dynamoDb

import (
	"context"
	"errors"
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type DynamoDBCustomerRepository struct {
	Client *dynamodb.Client
}

func NewDynamoDBCustomerRepository(client *dynamodb.Client) *DynamoDBCustomerRepository {
	return &DynamoDBCustomerRepository{Client: client}
}

func (repo *DynamoDBCustomerRepository) CreateUser(customer restmodel.Customer) (string, error) {
	var id = uuid.NewString()
	item, err := attributevalue.MarshalMap(customer)
	if err != nil {
		return "", err
	}

	item["Id"] = &types.AttributeValueMemberS{Value: id}

	fmt.Print(item)

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Customers"),
		Item:      item,
	})
	if err != nil {
		fmt.Printf("Couldn't Create User to table. Here's why: %v\n", err)
		return "", err
	}

	return id, nil
}

func (repo *DynamoDBCustomerRepository) UpdateUser(customer entity.Customer) (string, error) {
	item, err := attributevalue.MarshalMap(customer)
	if err != nil {
		return "", err
	}

	fmt.Print(item)

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Customers"),
		Item:      item,
	})
	if err != nil {
		fmt.Printf("Couldn't Create User to table. Here's why: %v\n", err)
		return "", err
	}

	return customer.Id, nil
}

func (repo *DynamoDBCustomerRepository) GetUserById(userId string) (*dynamodb.GetItemOutput, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Customers"),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: userId},
		},
	}

	result, err := repo.Client.GetItem(context.TODO(), input)
	if err != nil {
		fmt.Println("Error getting user by ID:", err)
		return nil, errors.New("error getting user by ID")
	}

	if len(result.Item) == 0 {
		return nil, errors.New("user not found")
	}

	return result, nil
}

func (repo *DynamoDBCustomerRepository) AddToken(userId, token string) (string, error) {
	fmt.Print(userId)
	item := map[string]types.AttributeValue{
		"CustomerId": &types.AttributeValueMemberS{Value: userId},
		"Token":      &types.AttributeValueMemberS{Value: token},
	}

	_, err := repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Tokens_c"),
		Item:      item,
	})
	if err != nil {
		fmt.Printf("Couldn't add token to table. Here's why: %v\n", err)
		return "", err
	}

	return token, nil
}
