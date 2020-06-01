package main

import (
	"log"

	"github.com/fsunset/goBasicAPI/controllers"
	"github.com/fsunset/goBasicAPI/middlewares"
	"github.com/fsunset/goBasicAPI/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Starts server
	server := gin.Default()

	// Adding Middleware, will be used in all below routes
	server.Use(middlewares.CheckStatusDB)

	// Set all API endpoints
	server.GET("/languages/list", controllers.ListLanguages)
	server.POST("/languages/new", controllers.CreateLanguage)
	server.GET("/languages/list/:id", controllers.ListLanguageByID)
	server.PATCH("/languages/update/:id", controllers.UpdateLanguage)

	// Runs server so it listens for requests...
	server.Run()

	// Checks DB connection at start
	if !models.CheckMongoDBConnection() {
		log.Fatal("Error connecting with MongoDB")
	}
}
