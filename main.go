package main

import (
	"log"
	"net/http"

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

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"response": "Hello World!"})
	})

	server.Run()

	if !models.CheckMongoDBConnection() {
		log.Fatal("Error connecting with MongoDB")
	}
}
