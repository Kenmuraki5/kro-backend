// cmd/myapp/main.go
package main

import (
	"net/http"

	"github.com/Kenmuraki5/kro-backend.git/internal/game"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"localhost": "8080",
		})
	})
	router.GET("/kro-games", game.GetAllGame) // Use GetAllGame directly as the handler
	router.Run("localhost:8080")
}
