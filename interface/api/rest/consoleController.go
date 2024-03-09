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

type ConsoleController struct {
	service interfaces.ConsoleService
}

func NewConsoleController(service interfaces.ConsoleService) *ConsoleController {
	return &ConsoleController{
		service: service,
	}
}

// set up router
func (gc *ConsoleController) SetupRoutes(router *gin.Engine) {
	consoleGroup := router.Group("/api/consoles")
	{
		consoleGroup.GET("", gc.GetAllConsolesHandler)
		consoleGroup.Use(middleware.AuthMiddleware(&auth.AuthService{}))
		consoleGroup.POST("/addConsole", gc.AddConsoleHandler)
		consoleGroup.PUT("/updateConsole", gc.UpdateConsoleHandler)
		consoleGroup.DELETE("/deleteConsole/:id", gc.DeleteConsole)
	}
}

func (controller *ConsoleController) GetAllConsolesHandler(c *gin.Context) {
	consoles, err := controller.service.GetAllConsoles()
	fmt.Println(consoles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch consoles"})
		return
	}

	c.JSON(http.StatusOK, consoles)
}

func (controller *ConsoleController) AddConsoleHandler(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if role != "kro-admin" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
		return
	}

	var newConsole restmodel.Console
	if err := c.ShouldBindJSON(&newConsole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedConsole, err := controller.service.AddConsole(newConsole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, addedConsole)
}

func (controller *ConsoleController) UpdateConsoleHandler(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if role != "kro-admin" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
		return
	}
	var updatedConsole entity.Console
	if err := c.ShouldBindJSON(&updatedConsole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consoles, err := controller.service.UpdateConsole(updatedConsole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update console"})
		return
	}

	c.JSON(http.StatusOK, consoles)
}

func (controller *ConsoleController) DeleteConsole(c *gin.Context) {
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
	err := controller.service.DeleteConsole(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete console"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Console item delete successfully"})
}
