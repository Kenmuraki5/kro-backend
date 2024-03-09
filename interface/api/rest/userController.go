package rest

import (
	"net/http"

	"github.com/Kenmuraki5/kro-backend.git/application/interfaces"
	"github.com/Kenmuraki5/kro-backend.git/application/services/auth"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/Kenmuraki5/kro-backend.git/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service interfaces.UserService
}

func NewUserController(service interfaces.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// set up router
func (gc *UserController) SetupRoutes(router *gin.Engine) {
	customerGroup := router.Group("/api/users")
	{
		customerGroup.GET("", middleware.AuthMiddleware(&auth.AuthService{}), gc.GetUserByEmailHandler)
		customerGroup.POST("/authentication", gc.Authentication)
		customerGroup.POST("/addUser", gc.CreateUserHandler)
		customerGroup.PUT("/updateUser", middleware.AuthMiddleware(&auth.AuthService{}), gc.UpdateUserHandler)
		customerGroup.GET("/alluser", middleware.AuthMiddleware(&auth.AuthService{}), gc.GetAllUserHandler)
	}
}

func (controller *UserController) GetAllUserHandler(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if role != "kro-admin" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
		return
	}
	user, err := controller.service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *UserController) GetUserByEmailHandler(c *gin.Context) {
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

func (controller *UserController) CreateUserHandler(c *gin.Context) {
	var newUser restmodel.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedConsole, err := controller.service.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, addedConsole)
}

func (controller *UserController) UpdateUserHandler(c *gin.Context) {
	var updateUser restmodel.User
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedConsole, err := controller.service.UpdateUser(updateUser, email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, addedConsole)
}

func (controller *UserController) Authentication(c *gin.Context) {
	var authRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&authRequest); err != nil {
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
