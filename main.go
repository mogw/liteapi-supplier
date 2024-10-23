package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"liteapi-supplier/controllers"
	"liteapi-supplier/services"
	"github.com/go-resty/resty/v2"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	// Initialize a real Resty client
	restyClient := resty.New()

	// Initialize Hotelbeds service with the Resty client
	hotelbedsService := services.NewHotelbedsService(restyClient)

	// Initialize the Hotels controller with the service
	hotelsController := controllers.NewHotelsController(hotelbedsService)

	// Set up Gin router
	r := gin.Default()

	// Define routes
	r.GET("/hotels/cheapest", hotelsController.GetCheapestRates)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
