package controllers

import (
	"net/http"

	"github.com/fsunset/goBasicAPI/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var newLanguageFromRequest models.CreateLanguageInput

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

// ListLanguageByID shows info for single language
func ListLanguageByID(ctx *gin.Context) {
	// Get parameter from URL
	languageID := ctx.Param("id")

	// Set correct structure for decoding query-result
	var languageFound models.Language

	objID, _ := primitive.ObjectIDFromHex(languageID)

	// Querying...
	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&languageFound)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error: ": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": languageFound,
	})
}

// UpdateLanguage alters document info
func UpdateLanguage(ctx *gin.Context) {
	// Search & find the document to edit
	// Get parameter from URL
	languageID := ctx.Param("id")

	// Set correct structure for decoding query-result
	var languageFound models.Language
	objID, _ := primitive.ObjectIDFromHex(languageID)

	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&languageFound)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error: ": err.Error(),
		})

		return
	}

	// Validate data from user
	var updateLanguageFromRequest models.UpdateLanguageInput

	err = ctx.ShouldBindJSON(&updateLanguageFromRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error: ": err.Error(),
		})

		return
	}

	// Create structure from user-input-data
	updateLanguageData := make(map[string]string)

	if len(updateLanguageFromRequest.Name) > 0 {
		updateLanguageData["name"] = updateLanguageFromRequest.Name
	}
	if len(updateLanguageFromRequest.Creator) > 0 {
		updateLanguageData["creator"] = updateLanguageFromRequest.Creator
	}

	updateData := bson.M{
		"$set": updateLanguageData,
	}

	// Filter for "UpdateOne" function
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	// Updating document
	documentUpdated, err := collection.UpdateOne(ctx, filter, updateData)

	ctx.JSON(http.StatusOK, gin.H{
		"result": documentUpdated,
	})
}
