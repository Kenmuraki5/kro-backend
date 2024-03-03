// main.go

package main

import (
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/application/services"
	"github.com/Kenmuraki5/kro-backend.git/application/services/auth"
	"github.com/Kenmuraki5/kro-backend.git/infrastructure/persistence/dynamoDb"
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
	//Customers
	customerRepo := dynamoDb.NewDynamoDBCustomerRepository(dbClient.Client)
	customerservice := services.NewCustomerService(customerRepo, auth.AuthService{})
	customerController := rest.NewCustomerController(customerservice)

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

	router.Use(cors.Default())

	gameController.SetupRoutes(router)
	orderController.SetupRoutes(router)
	customerController.SetupRoutes(router)
	consoleController.SetupRoutes(router)

	err = router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
