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
	"golang.org/x/crypto/bcrypt"
)

type DynamoDBCustomerRepository struct {
	Client *dynamodb.Client
}

func NewDynamoDBCustomerRepository(client *dynamodb.Client) *DynamoDBCustomerRepository {
	return &DynamoDBCustomerRepository{Client: client}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (repo *DynamoDBCustomerRepository) CreateUser(customer entity.Customer) (string, error) {
	hashedPassword, err := hashPassword(customer.Password)
	if err != nil {
		return "", err
	}
	customer.Password = hashedPassword
	item, err := attributevalue.MarshalMap(customer)
	if err != nil {
		return "", err
	}
	input := &dynamodb.PutItemInput{
		TableName:           aws.String("Customers"),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(Email)"),
	}

	_, err = repo.Client.PutItem(context.Background(), input)
	if err != nil {
		fmt.Printf("Couldn't Create User to table. Here's why: %v\n", err)
		return "", err
	}

	return customer.Email, nil
}

func (repo *DynamoDBCustomerRepository) UpdateUser(customer restmodel.Customer, email string) (string, error) {
	item, err := attributevalue.MarshalMap(customer)
	if err != nil {
		return "", err
	}

	key := map[string]types.AttributeValue{
		"Email": &types.AttributeValueMemberS{Value: email},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:        aws.String("Customers"),
		Key:              key,
		UpdateExpression: aws.String("SET FullName = :fullname, PhoneNumber = :phoneNumber, Address = :address"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":fullname":    item["FullName"],
			":phoneNumber": item["PhoneNumber"],
			":address":     item["Address"],
		},
	}

	_, err = repo.Client.UpdateItem(context.Background(), input)
	if err != nil {
		fmt.Printf("Couldn't Update User in table. Here's why: %v\n", err)
		return "", err
	}

	return email, nil
}

func (repo *DynamoDBCustomerRepository) GetUserByEmail(email string) (*entity.Customer, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Customers"),
		Key: map[string]types.AttributeValue{
			"Email": &types.AttributeValueMemberS{Value: email},
		},
	}

	result, err := repo.Client.GetItem(context.TODO(), input)
	if err != nil {
		fmt.Println("Error getting user by email:", err)
		return nil, fmt.Errorf("error getting user by email: %w", err)
	}

	if len(result.Item) == 0 {
		return nil, errors.New("user not found")
	}

	var customer entity.Customer
	err = attributevalue.UnmarshalMap(result.Item, &customer)
	if err != nil {
		fmt.Println("Error unmarshaling DynamoDB result:", err)
		return nil, fmt.Errorf("error unmarshaling DynamoDB result: %w", err)
	}

	return &customer, nil
}

func (repo *DynamoDBCustomerRepository) AuthenticateUser(email, password string) (bool, error) {
	result, err := repo.GetUserByEmail(email)
	if err != nil {
		return false, err
	}

	if result == nil {
		return false, errors.New("user not found")
	}

	hashedPassword := result.Password
	if hashedPassword == "" {
		return false, errors.New("password not found in user record")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
