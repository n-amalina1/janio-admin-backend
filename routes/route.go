package routes

import (
	"database/sql"
	"net/http"

	api "backend/api"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetupRoutes(d *sql.DB) {
	db = d
	router := gin.Default()
	router.GET("admin/orders", GetOrdersAdmin)

	router.GET("id/orders", GetOrdersIDProvider)

	router.Run("localhost:8080")
}

func GetOrdersAdmin(c *gin.Context) {
	orders, err := api.GetOrdersAdmin(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Orders not found"})
	}
	c.IndentedJSON(http.StatusOK, orders)
}

func GetOrdersIDProvider(c *gin.Context) {
	orders, err := api.GetOrdersIDProvider(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Orders not found"})
	}
	c.IndentedJSON(http.StatusOK, orders)
}
