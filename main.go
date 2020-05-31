package main

import (
	"log"

	"github.com/fsunset/goBasicAPI/controllers"
	"github.com/fsunset/goBasicAPI/middlewares"
	"github.com/fsunset/goBasicAPI/models"
	"github.com/gin-gonic/gin"
)

const (
	dbUser     = "twittAppClusterUser"
	dbPassword = "twittAppClusterPass20"
	dbCluster  = "twittappcluster-6ibps.mongodb.net"
)

func main() {
	server := gin.Default()

	// Adding Middleware, will be used in all below routes
	server.Use(middlewares.CheckStatusDB)

	server.GET("/languages", controllers.FindAll)

	server.Run()

	if !models.CheckMongoDBConnection() {
		log.Fatal("Error connecting with MongoDB")
	}
}
