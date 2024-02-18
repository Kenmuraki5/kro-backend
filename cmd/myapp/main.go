// cmd/myapp/main.go
package main

import (
	"context"
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/internal/database"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	// Create a DynamoDB client
	dynamoDB, err := database.NewDynamoDBClient()
	if err != nil {
		panic(err)
	}

	// Now you can use dynamoDB.Client to interact with DynamoDB
	// For example, list tables:
	listTablesInput := &dynamodb.ListTablesInput{}
	resp, err := dynamoDB.Client.ListTables(context.TODO(), listTablesInput)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tables in DynamoDB:")
	for _, tableName := range resp.TableNames {
		fmt.Println(tableName)
	}
}
