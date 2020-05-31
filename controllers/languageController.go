package controllers

import (
	"net/http"

	"github.com/fsunset/goBasicAPI/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// "Global" var for the whole package
var collection = models.MongoConnection.Database("languagesDB").Collection("language")

// ListLanguages returns all languages in DB
func ListLanguages(ctx *gin.Context) {
	// Initialize slice to store ALL collection's documents returned by query. Must have correct "Language" structure
	var allLanguages []models.Language

	// Quering...
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error: ": err.Error(),
		})

		return
	}

	// Prepare data to be returned...
	for cur.Next(ctx) {
		// Initialize variable to store each document. Must have correct "Language" structure
		var language models.Language

		err := cur.Decode(&language)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error: ": err.Error(),
			})

			return
		}

		// Appending each document within "allLanguages" slice
		allLanguages = append(allLanguages, language)
	}

	// Returning data
	ctx.JSON(http.StatusOK, gin.H{
		"results": allLanguages,
	})
}

// CreateLanguage inserts a new Language-document into DB
func CreateLanguage(ctx *gin.Context) {
	// Initialize correct structure
	var newLanguageFromRequest models.CreateLanguage

	// Validates user's input data
	err := ctx.ShouldBindJSON(&newLanguageFromRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error: ": err.Error(),
		})

		return
	}

	var existentLanguage models.Language

	// Checks if Language.Name already exists. It must be unique
	err = collection.FindOne(ctx, bson.M{"name": newLanguageFromRequest.Name}).Decode(&existentLanguage)
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error:": "Language already exists",
		})

		return
	}

	// Get data & insert into DB
	newLanguage := models.Language{
		Name:    newLanguageFromRequest.Name,
		Creator: newLanguageFromRequest.Creator,
	}

	insertedLanguage, err := collection.InsertOne(ctx, newLanguage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error: ": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": insertedLanguage,
	})
}
