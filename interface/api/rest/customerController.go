package rest

import (
	"net/http"

	"github.com/Kenmuraki5/kro-backend.git/application/interfaces"
	"github.com/Kenmuraki5/kro-backend.git/application/services/auth"
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/Kenmuraki5/kro-backend.git/interface/middleware"
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
	customerGroup := router.Group("/Customers")
	{
		customerGroup.GET("", middleware.AuthMiddleware(&auth.AuthService{}), gc.GetUserByIdHandler)

		customerGroup.POST("/addCustomer", gc.CreateUserHandler)
		customerGroup.PUT("/updateCustomer", gc.UpdateUserHandler)
	}
}

func (controller *CustomerController) GetUserByIdHandler(c *gin.Context) {
	id, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	user, err := controller.service.GetUserById(id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Customer"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *CustomerController) CreateUserHandler(c *gin.Context) {
	var newCustomer restmodel.Customer
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
	var updateCustomer entity.Customer
	if err := c.ShouldBindJSON(&updateCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedConsole, err := controller.service.UpdateUser(updateCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, addedConsole)
}
