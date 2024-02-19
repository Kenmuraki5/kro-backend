// internal/gameservice/gameservice.go
package gameservice

import (
	"net/http"

	"github.com/Kenmuraki5/kro-backend.git/internal/database"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func GetAllGame(c *gin.Context) {
	dbClient, err := database.NewDynamoDBClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating DynamoDB client"})
		return
	}

	// Use the Scan operation to get all items in the "game" table
	result, err := dbClient.Client.Scan(c.Request.Context(), &dynamodb.ScanInput{
		TableName: aws.String("game"),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying DynamoDB"})
		return
	}

	// Check if there are no items in the result
	if len(result.Items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No games found"})
		return
	}

	// Iterate through the items and create a response
	c.JSON(http.StatusOK, result.Items)
}
