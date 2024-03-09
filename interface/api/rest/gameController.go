package rest

import (
	"fmt"
	"net/http"

	"github.com/Kenmuraki5/kro-backend.git/application/interfaces"
	"github.com/Kenmuraki5/kro-backend.git/application/services/auth"
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/Kenmuraki5/kro-backend.git/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type GameController struct {
	service interfaces.GameService
}

func NewGameController(service interfaces.GameService) *GameController {
	return &GameController{
		service: service,
	}
}

// set up router
func (gc *GameController) SetupRoutes(router *gin.Engine) {
	gameGroup := router.Group("/api/games")
	{
		gameGroup.GET("", gc.GetAllGamesHandler)
		gameGroup.Use(middleware.AuthMiddleware(&auth.AuthService{}))
		gameGroup.POST("/addGame", gc.AddGameHandler)
		gameGroup.PUT("/updateGame", gc.UpdateGameHandler)
		gameGroup.DELETE("/deleteGame/:id", gc.DeleteGame)
	}
}

func (controller *GameController) GetAllGamesHandler(c *gin.Context) {
	games, err := controller.service.GetAllGames()
	fmt.Println(games)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch games"})
		return
	}

	c.JSON(http.StatusOK, games)
}

func (controller *GameController) AddGameHandler(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if role != "kro-admin" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
		return
	}
	var newGame restmodel.Game
	if err := c.ShouldBindJSON(&newGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedGame, err := controller.service.AddGame(newGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, addedGame)
}

func (controller *GameController) UpdateGameHandler(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if role != "kro-admin" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
		return
	}
	var updatedGame entity.Game
	if err := c.ShouldBindJSON(&updatedGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	games, err := controller.service.UpdateGame(updatedGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update game"})
		return
	}

	c.JSON(http.StatusOK, games)
}

func (controller *GameController) DeleteGame(c *gin.Context) {
	id := c.Param("id")
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if role != "kro-admin" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
		return
	}
	err := controller.service.DeleteGame(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Game item delete successfully"})
}
