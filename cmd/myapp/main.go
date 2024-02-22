// main.go

package main

import (
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/application/services"
	"github.com/Kenmuraki5/kro-backend.git/infrastructure/persistence/dynamoDb"
	"github.com/Kenmuraki5/kro-backend.git/interface/api/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	dbClient, err := dynamoDb.NewDynamoDBClient()

	if err != nil {
		fmt.Printf("Error initializing DynamoDB client: %v\n", err)
		return
	}

	gameRepo := dynamoDb.NewDynamoDBGameRepository(dbClient.Client)

	gameService := services.NewGameService(gameRepo)

	gameController := rest.NewGameController(gameService)

	router := gin.Default()

	// Register route for GetAllGamesHandler
	router.GET("/games", gameController.GetAllGamesHandler)

	err = router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
