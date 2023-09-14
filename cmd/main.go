package main

import (
	"database/sql"

	api "backend/api"
	routes "backend/routes"
)

var db *sql.DB

func init() {
	db = api.SetupDBConn()
}

func main() {
	routes.SetupRoutes(db)
}
