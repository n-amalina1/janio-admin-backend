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
	router.Use(CORSMiddleware())

	router.GET("client/orders", PostOrdersClient)

	router.GET("admin/orders", GetOrdersAdmin)
	router.PUT("admin/order", UpdateOrderAdmin)
	router.POST("admin/order", PostOrderAdmin)
	router.DELETE("admin/order", DeleteOrderAdmin)

	router.GET("id/orders", GetOrdersIDProvider)
	router.GET("my/orders", GetOrdersMYProvider)

	router.Run("localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
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

func GetOrdersMYProvider(c *gin.Context) {
	orders, err := api.GetOrdersMYProvider(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Orders not found"})
	}
	c.IndentedJSON(http.StatusOK, orders)
}
