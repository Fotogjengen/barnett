package main

import (
	"barnett/api/database"
	"barnett/api/handlers"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default() // Set router

	store := cookie.NewStore([]byte("sessionSuperSecret"))
	r.Use(sessions.Sessions("mySession", store))


	r.LoadHTMLGlob("templates/*")

	// r.GET("/", handlers.GetRoot)

	// Setup route group for the API
	api := r.Group("/api")

	// Routes
	api.GET("/", handlers.GetRoot) // Get root structure with links to all paths
	api.GET("/products", handlers.GetProducts) // Get all products
	api.POST("/products", handlers.AddProduct) // Post product
	api.GET("/categories", handlers.GetCategories) // Get all categories

	// TODO: Set permissions for this group
	users := r.Group("/users")
	users.POST("/signup", handlers.Signup)
	users.POST("/login", handlers.Login)
	users.GET("/logout", handlers.Logout)

	database.Connect()

	defer database.DB.Close() // Closes DB when main() returns TODO: Error handling

	err := database.DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to DB!")

	// database.Migrate()

	if err := r.Run(":8080"); err != nil { // Run on port 8080
		panic(err)
	}
}
