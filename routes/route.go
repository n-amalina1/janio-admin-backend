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
	router.GET("/orders", getOrders)

	router.Run("localhost:8080")
}

func getOrders(c *gin.Context) {
	orders, err := api.GetAllOrdersDB(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Orders not found"})
	}
	c.IndentedJSON(http.StatusOK, orders)
}
