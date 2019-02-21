package main

import (
	"barnett/api/handlers"
	"barnett/api/database"
	"database/sql"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {

	r := gin.Default() // Set router

	// Routes
	r.GET("/api/products", handlers.GetProducts) // Get all products
	r.POST("/api/products", handlers.AddProduct) // Post product

	database.Connect()
	// database.Migrate()
	defer database.DB.Close() // Closes DB when main() returns TODO: Error handling

	database.Migrate()

	if err := r.Run(":8080"); err != nil { // Run on port 8080
		panic(err)
	}
}
