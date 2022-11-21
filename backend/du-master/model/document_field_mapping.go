package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentFieldMapping struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	DocumentFieldName string             `bson:"documentFieldName,omitempty" json:"documentFieldName,omitempty"`
	DocumentName      string             `bson:"documentName,omitempty" json:"documentName,omitempty"`
	Status            string             `bson:"status,omitempty" json:"status,omitempty"`
	CreationDateTime  time.Time          `bson:"creationDateTime,omitempty" json:"creationDateTime,omitempty"`
}
