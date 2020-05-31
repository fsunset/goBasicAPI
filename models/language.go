package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Language handles language-model-structure
type Language struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name    string             `bson:"name" json:"name"`
	Creator string             `bson:"creator" json:"creator"`
}
