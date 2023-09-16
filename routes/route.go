package routes

import (
	"database/sql"
	"net/http"

	api "backend/api"
	"backend/models"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetupRoutes(d *sql.DB) {
	db = d
	router := gin.Default()
	router.POST("client/orders", PostOrdersClient)

	router.GET("admin/orders", GetOrdersAdmin)
	router.PUT("admin/order", UpdateOrderAdmin)
	router.POST("admin/order", PostOrderAdmin)
	router.DELETE("admin/order", DeleteOrderAdmin)

	router.GET("id/orders", GetOrdersIDProvider)
	router.GET("my/orders", GetOrdersMYProvider)

	router.Run("localhost:8080")
}

func PostOrdersClient(c *gin.Context) {
	orders := api.PostOrdersClient(db)
	c.IndentedJSON(http.StatusOK, orders)
}

func GetOrdersAdmin(c *gin.Context) {
	orders, err := api.GetOrdersAdmin(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}
	c.IndentedJSON(http.StatusOK, orders)
}

func UpdateOrderAdmin(c *gin.Context) {
	var updatedOrder models.UpdateAdminOrder

	if err := c.BindJSON(&updatedOrder); err != nil {
		return
	}

	order, err := api.UpdateOrderAdmin(db, &updatedOrder)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"data": order})
	}

}

func PostOrderAdmin(c *gin.Context) {
	var newOrder models.PostAdminOrder

	if err := c.BindJSON(&newOrder); err != nil {
		return
	}

	order, err := api.PostOrderAdmin(db, &newOrder)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"data": order})
	}
}

func DeleteOrderAdmin(c *gin.Context) {
	var orderId models.DeletAdminOrder

	if err := c.BindJSON(&orderId); err != nil {
		return
	}

	_, err := api.DeleteOrderAdmin(db, orderId.OrderID)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"orderId": orderId.OrderID})
	}
}

func GetOrdersIDProvider(c *gin.Context) {
	orders, err := api.GetOrdersIDProvider(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Orders not found"})
	}
	c.IndentedJSON(http.StatusOK, orders)
}

func GetOrdersMYProvider(c *gin.Context) {
	orders, err := api.GetOrdersMYProvider(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Orders not found"})
	}
	c.IndentedJSON(http.StatusOK, orders)
}
