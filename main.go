package main

import (
	"log"
	"os"

	"github.com/aalekh12/Blog-Post-Api/database"
	_ "github.com/aalekh12/Blog-Post-Api/docs" // Import generated Swagger docs
	"github.com/aalekh12/Blog-Post-Api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // Import Swagger middleware
	"github.com/joho/godotenv"
)

// @title Blog API
// @version 1.0
// @description Blog Post API is a RESTful API designed to manage blog posts. It provides CRUD (Create, Read, Update, Delete) functionality, allowing users to create, update, fetch, and delete blog posts. This API is built using Go and the Fiber framework and is hosted on Render.
// @host blog-post-api-chmy.onrender.com
func main() {
	// Load environment variables
	godotenv.Load()

	// Initialize Fiber app
	app := fiber.New()

	// Connect to database
	go database.ConnectDataBase()

	// Register API routes
	routes.RegisterRoutes(app)

	// Add Swagger documentation route
	app.Get("/swagger/*", swagger.HandlerDefault) // Access Swagger UI at /swagger/index.html

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("HTTP server started on :%s", port)
	log.Fatal(app.Listen(":" + port))
}
