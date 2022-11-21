package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Document struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DocumentName           string             `bson:"documentName" json:"documentName" binding:"required"`
	DocumentType           string             `bson:"documentType" json:"documentType" binding:"required"`
	DocumentDiscription    string             `bson:"documentDiscription" json:"documentDiscription" binding:"required"`
	Status                 string             `bson:"status" json:"status"`
	DocumentTrainingStatus string             `bson:"documentTrainingStatus" json:"documentTrainingStatus" binding:"required"`
	CreationDateTime       time.Time          `bson:"creationDateTime" json:"creationDateTime"`
}
