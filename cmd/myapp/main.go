// main.go

package main

import (
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/application/services"
	"github.com/Kenmuraki5/kro-backend.git/application/services/auth"
	"github.com/Kenmuraki5/kro-backend.git/infrastructure/persistence/dynamoDb"
	s3 "github.com/Kenmuraki5/kro-backend.git/infrastructure/persistence/s3upload"
	"github.com/Kenmuraki5/kro-backend.git/interface/api/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	dbClient, err := dynamoDb.NewDynamoDBClient()
	if err != nil {
		fmt.Printf("Error initializing DynamoDB client: %v\n", err)
		return
	}

	//s3
	s3Service := s3.NewS3Uploader()
	// endpoint, err := s3Service.UploadFile("path/to/your/file.jpg")
	// if err != nil {
	// 	fmt.Printf("Error uploading to S3: %v\n", err)
	// 	return
	// }

	// fmt.Println("Picture uploaded successfully to:", endpoint)
	// //users
	userRepo := dynamoDb.NewDynamoDBUserRepository(dbClient.Client)
	userservice := services.NewUserService(userRepo, auth.AuthService{}, s3Service)
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

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Next()
	})

	gameController.SetupRoutes(router)
	orderController.SetupRoutes(router)
	userController.SetupRoutes(router)
	consoleController.SetupRoutes(router)

	err = router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
