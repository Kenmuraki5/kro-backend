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

type CustomerController struct {
	service interfaces.CustomerService
}

func NewCustomerController(service interfaces.CustomerService) *CustomerController {
	return &CustomerController{
		service: service,
	}
}

// set up router
func (gc *CustomerController) SetupRoutes(router *gin.Engine) {
	customerGroup := router.Group("/customers")
	{
		customerGroup.GET("", middleware.AuthMiddleware(&auth.AuthService{}), gc.GetUserByEmailHandler)
		customerGroup.POST("/authentication", gc.Authentication)
		customerGroup.POST("/addCustomer", gc.CreateUserHandler)
		customerGroup.PUT("/updateCustomer", middleware.AuthMiddleware(&auth.AuthService{}), gc.UpdateUserHandler)
	}
}

func (controller *CustomerController) GetUserByEmailHandler(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}

	user, err := controller.service.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Customer"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *CustomerController) CreateUserHandler(c *gin.Context) {
	var newCustomer entity.Customer
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedConsole, err := controller.service.AddUser(newCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, addedConsole)
}

func (controller *CustomerController) UpdateUserHandler(c *gin.Context) {
	var updateCustomer restmodel.Customer
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if err := c.ShouldBindJSON(&updateCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedConsole, err := controller.service.UpdateUser(updateCustomer, email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, addedConsole)
}

func (controller *CustomerController) Authentication(c *gin.Context) {
	var authRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&authRequest); err != nil {
		fmt.Println("errorrrewgaawg")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := controller.service.AuthenticateUser(authRequest.Email, authRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	c.JSON(http.StatusOK, token)
}
