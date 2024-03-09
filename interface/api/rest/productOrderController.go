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

type OrderController struct {
	service interfaces.OrderService
}

func NewOrderController(service interfaces.OrderService) *OrderController {
	return &OrderController{
		service: service,
	}
}

// set up router
func (gc *OrderController) SetupRoutes(router *gin.Engine) {
	OrderGroup := router.Group("/api/orders")
	{
		OrderGroup.Use(middleware.AuthMiddleware(&auth.AuthService{}))
		OrderGroup.GET("", gc.GetAllOrdersHandler)
		OrderGroup.GET("/userOrders", gc.GetOrdersByEmailHandler)
		OrderGroup.POST("/createPaymentToken", gc.CreatePaymentTokenHandler)
		OrderGroup.POST("/addOrders", gc.AddOrderHandler)
		OrderGroup.PUT("/updateOrder", gc.UpdateOrderHandler)
		OrderGroup.DELETE("/deleteOrder/:orderId/:productId", gc.DeleteOrderHandler)
	}
}

func (controller *OrderController) GetAllOrdersHandler(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if role != "kro-admin" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
		return
	}
	Orders, err := controller.service.GetAllOrders()
	fmt.Println(Orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Orders"})
		return
	}

	c.JSON(http.StatusOK, Orders)
}

func (controller *OrderController) GetOrdersByEmailHandler(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	Orders, err := controller.service.GetOrdersByEmail(email.(string))

	fmt.Println(Orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Orders"})
		return
	}

	c.JSON(http.StatusOK, Orders)
}

func (controller *OrderController) AddOrderHandler(c *gin.Context) {
	var orderData struct {
		NewOrder []restmodel.Order `json:"newOrder"`
		Token    string            `json:"token"`
		Amount   int64             `json:"amount"`
	}

	if err := c.ShouldBindJSON(&orderData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedOrder, err := controller.service.AddOrders(orderData.NewOrder, orderData.Token, orderData.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, addedOrder)
}

func (controller *OrderController) UpdateOrderHandler(c *gin.Context) {
	var updatedOrder entity.Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Orders, err := controller.service.UpdateOrder(updatedOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Orders)
}

func (controller *OrderController) DeleteOrderHandler(c *gin.Context) {
	orderId := c.Param("orderId")
	productId := c.Param("productId")
	err := controller.service.DeleteOrder(orderId, productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order item delete successfully"})
}

func (controller *OrderController) CreatePaymentTokenHandler(c *gin.Context) {
	var payment restmodel.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := controller.service.CreatePaymentToken(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, token)
}
