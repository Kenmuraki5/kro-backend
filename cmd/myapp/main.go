// main.go

package main

import (
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/application/services"
	"github.com/Kenmuraki5/kro-backend.git/application/services/auth"
	"github.com/Kenmuraki5/kro-backend.git/infrastructure/persistence/dynamoDb"
	s3 "github.com/Kenmuraki5/kro-backend.git/infrastructure/persistence/s3upload"
	"github.com/Kenmuraki5/kro-backend.git/interface/api/rest"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dbClient, err := dynamoDb.NewDynamoDBClient()
	if err != nil {
		fmt.Printf("Error initializing DynamoDB client: %v\n", err)
		return
	}

	// //users
	userRepo := dynamoDb.NewDynamoDBUserRepository(dbClient.Client)
	userservice := services.NewUserService(userRepo, auth.AuthService{})
	userController := rest.NewUserController(userservice)

	//Games
	gameRepo := dynamoDb.NewDynamoDBGameRepository(dbClient.Client)
	gameService := services.NewGameService(gameRepo)
	gameController := rest.NewGameController(gameService)

	//Consoles
	consoleRepo := dynamoDb.NewDynamoDBConsoleRepository(dbClient.Client)
	consoleService := services.NewConsoleService(consoleRepo)
	consoleController := rest.NewConsoleController(consoleService)

	//Orders
	orderRepo := dynamoDb.NewDynamoDBOrderRepository(dbClient.Client)
	orderService := services.NewOrderService(orderRepo, gameRepo, consoleRepo)
	orderController := rest.NewOrderController(orderService)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE", "PUT"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}

	router.Use(cors.New(config))

	//S3
	router.POST("/s3/upload-image", s3.S3uploader)

	gameController.SetupRoutes(router)
	orderController.SetupRoutes(router)
	userController.SetupRoutes(router)
	consoleController.SetupRoutes(router)

	err = router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
