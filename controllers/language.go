package controllers

import (
	"net/http"

	"github.com/fsunset/goBasicAPI/models"
	"github.com/gin-gonic/gin"
)

// FindAll returns all languages in DB
func FindAll(ctx *gin.Context) {
	// Initialize variable to store query-results. Must have correct structure
	var languages []models.Language

	// Quering...
	models.MongoConnection.Database("languagesDB").Collection("languages").Find(ctx, &languages)

	// Returning data
	ctx.JSON(http.StatusOK, gin.H{
		"results": languages,
	})
}
