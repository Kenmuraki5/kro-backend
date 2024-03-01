package rest

import (
	"fmt"
	"net/http"

	"github.com/Kenmuraki5/kro-backend.git/application/interfaces"
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
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
	OrderGroup := router.Group("/orders")
	{
		OrderGroup.GET("", gc.GetAllOrdersHandler)
		OrderGroup.POST("/addOrders", gc.AddOrderHandler)
		OrderGroup.PUT("/updateOrder", gc.UpdateOrderHandler)
		OrderGroup.DELETE("/deleteOrder/:orderId/:productId", gc.DeleteOrder)
	}
}

func (controller *OrderController) GetAllOrdersHandler(c *gin.Context) {
	Orders, err := controller.service.GetAllOrders()
	fmt.Println(Orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Orders"})
		return
	}

	c.JSON(http.StatusOK, Orders)
}

func (controller *OrderController) AddOrderHandler(c *gin.Context) {
	var newOrder []restmodel.Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedOrder, err := controller.service.AddOrders(newOrder)
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

func (controller *OrderController) DeleteOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	productId := c.Param("productId")
	err := controller.service.DeleteOrder(orderId, productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order item delete successfully"})
}
