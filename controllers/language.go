package controllers

import (
	"log"
	"net/http"

	"github.com/fsunset/goBasicAPI/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// FindAll returns all languages in DB
func FindAll(ctx *gin.Context) {
	collection := models.MongoConnection.Database("languagesDB").Collection("language")

	// Initialize slice to store ALL collection's documents returned by query. Must have correct "Language" structure
	var allLanguages []models.Language

	// Quering...
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Prepare data to be returned...
	for cur.Next(ctx) {
		// Initialize variable to store each document. Must have correct "Language" structure
		var language models.Language

		err := cur.Decode(&language)
		if err != nil {
			log.Fatal(err)
		}

		// Appending each document within "allLanguages" slice
		allLanguages = append(allLanguages, language)
	}

	// Returning data
	ctx.JSON(http.StatusOK, gin.H{
		"results": allLanguages,
	})
}
