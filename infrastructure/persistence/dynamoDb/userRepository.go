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

type DynamoDBUserRepository struct {
	Client *dynamodb.Client
}

func NewDynamoDBUserRepository(client *dynamodb.Client) *DynamoDBUserRepository {
	return &DynamoDBUserRepository{Client: client}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (repo *DynamoDBUserRepository) CreateUser(user restmodel.User) (string, error) {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword
	item, err := attributevalue.MarshalMap(user)
	item["role"] = &types.AttributeValueMemberS{Value: "customer"}

	if err != nil {
		return "", err
	}
	input := &dynamodb.PutItemInput{
		TableName:           aws.String("Users"),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(Email)"),
	}

	_, err = repo.Client.PutItem(context.Background(), input)
	if err != nil {
		fmt.Printf("Couldn't Create User to table. Here's why: %v\n", err)
		return "", err
	}

	return user.Email, nil
}

func (repo *DynamoDBUserRepository) UpdateUser(user restmodel.User, email string) (string, error) {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return "", err
	}

	key := map[string]types.AttributeValue{
		"Email": &types.AttributeValueMemberS{Value: email},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:        aws.String("Users"),
		Key:              key,
		UpdateExpression: aws.String("SET FullName = :fullname, PhoneNumber = :phoneNumber, Address = :address, ImageProfile = :imageProfile"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":fullname":     item["FullName"],
			":phoneNumber":  item["PhoneNumber"],
			":address":      item["Address"],
			":imageProfile": item["ImageProfile"],
		},
	}

	_, err = repo.Client.UpdateItem(context.Background(), input)
	if err != nil {
		fmt.Printf("Couldn't Update User in table. Here's why: %v\n", err)
		return "", err
	}

	return email, nil
}

func (repo *DynamoDBUserRepository) GetUserByEmail(email string) (*entity.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Users"),
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

	var user entity.User
	err = attributevalue.UnmarshalMap(result.Item, &user)
	if err != nil {
		fmt.Println("Error unmarshaling DynamoDB result:", err)
		return nil, fmt.Errorf("error unmarshaling DynamoDB result: %w", err)
	}

	return &user, nil
}

func (repo *DynamoDBUserRepository) GetAllUser() ([]*entity.User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Users"),
	}
	result, err := repo.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan DynamoDB table: %v", err)
	}
	var users []*entity.User
	for _, item := range result.Items {
		var user entity.User
		err := attributevalue.UnmarshalMap(item, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	fmt.Println(users)
	return users, nil
}

func (repo *DynamoDBUserRepository) AuthenticateUser(email, password string) (string, error) {
	result, err := repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if result == nil {
		return "", errors.New("user not found")
	}
	fmt.Println("role:", result.Role)
	hashedPassword := result.Password
	if hashedPassword == "" {
		return "", errors.New("password not found in user record")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", err
	}
	return result.Role, nil
}
