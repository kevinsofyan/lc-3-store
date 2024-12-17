package main

import (
	"log"
	"os"
	"store/config"
	"store/routes"

	_ "store/docs" // This is required to initialize the Swagger docs

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title FTGO Phase 2 Livecode 3 - Kevin Sofyan
// @livecode 3 - Kevin Sofyan
// @version 1.0
// @contact.name Benedict Kevin Sofyan /kevinsofyan.13@gmail.com
// @host orders-lc3-717e5d36a486.herokuapp.com
// @BasePath /

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	e := echo.New()

	// Initialize the database
	config.InitDB()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize routes
	routes.InitRoutes(e)

	// Serve Swagger UI
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Get the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	log.Printf("Starting server on port %s", port)

	e.Logger.Fatal(e.Start(":" + port))
}
