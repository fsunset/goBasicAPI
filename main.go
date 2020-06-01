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
	server.GET("/languages", controllers.ListLanguages)
	server.POST("/languages", controllers.CreateLanguage)
	server.GET("/languages/:id", controllers.ListLanguageByID)
	server.PATCH("/languages/:id", controllers.UpdateLanguage)
	server.DELETE("/languages/:id", controllers.DeleteLanguage)

	// Runs server so it listens for requests...
	server.Run()

	// Checks DB connection at start
	if !models.CheckMongoDBConnection() {
		log.Fatal("Error connecting with MongoDB")
	}
}
