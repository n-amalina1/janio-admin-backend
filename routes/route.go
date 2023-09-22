package routes

import (
	"database/sql"
	"net/http"
	"time"

	api "backend/api"
	"backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetupRoutes(d *sql.DB) {
	db = d
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8008", "http://localhost:8443", "http://localhost:9883"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("client/orders", PostOrdersClient)

	router.GET("admin/orders", GetOrdersAdmin)
	router.PUT("admin/order", UpdateOrderAdmin)
	router.POST("admin/order", PostOrderAdmin)
	router.DELETE("admin/order", DeleteOrderAdmin)

	router.GET("id/orders", GetOrdersIDProvider)
	router.POST("id/order/update", PutStatusIDProvider)

	router.GET("my/orders", GetOrdersMYProvider)

	router.Run("localhost:8080")
}

func PostOrdersClient(c *gin.Context) {
	var newOrders []models.ClientToDBOrder
	if err := c.BindJSON(&newOrders); err != nil {
		return
	}

	orders, err := api.PostOrdersClient(db, &newOrders)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"data": orders})
	}
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"data": order})
	}
}

func DeleteOrderAdmin(c *gin.Context) {
	var orderId models.DeleteAdminOrder

	if err := c.BindJSON(&orderId); err != nil {
		return
	}

	_, err := api.DeleteOrderAdmin(db, orderId.OrderID)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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

func PutStatusIDProvider(c *gin.Context) {
	var status models.IDOrderStatus

	if err := c.BindJSON(&status); err != nil {
		return
	}
	orderStatus, errS := api.PostStatusIDProvider(db, status)

	if errS != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": errS.Error()})
	}
	c.IndentedJSON(http.StatusOK, orderStatus)
}

func GetOrdersMYProvider(c *gin.Context) {
	orders, err := api.GetOrdersMYProvider(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Orders not found"})
	}
	c.IndentedJSON(http.StatusOK, orders)
}
