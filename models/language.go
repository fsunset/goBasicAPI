package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Language handles language-model-structure
type Language struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name    string             `bson:"name" json:"name"`
	Creator string             `bson:"creator" json:"creator"`
}

// CreateLanguageInput handles creation of language-structure
type CreateLanguageInput struct {
	Name    string `bson:"name" json:"name" binding:"required"`
	Creator string `bson:"creator" json:"creator" binding:"required"`
}

// UpdateLanguageInput handles updating of language-structure
type UpdateLanguageInput struct {
	Name    string `bson:"name" json:"name"`
	Creator string `bson:"creator" json:"creator"`
}
