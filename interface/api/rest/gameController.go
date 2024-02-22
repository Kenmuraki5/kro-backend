package rest

import (
	"fmt"
	"net/http"

	"github.com/Kenmuraki5/kro-backend.git/application/interfaces"
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

func (controller *GameController) GetAllGamesHandler(c *gin.Context) {
	games, err := controller.service.GetAllGames()
	fmt.Println(games)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch games"})
		return
	}

	c.JSON(http.StatusOK, games)
}
