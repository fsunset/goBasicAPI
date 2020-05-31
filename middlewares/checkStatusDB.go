package middlewares

import (
	"log"

	"github.com/fsunset/goBasicAPI/models"
	"github.com/gin-gonic/gin"
)

// CheckStatusDB checks DB status
func CheckStatusDB(ctx *gin.Context) {
	if !models.CheckMongoDBConnection() {
		log.Fatal("Bad DB connection")
	}

	// If DB connection is OK; pass on to the next-in-chain
	ctx.Next()
}
