package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientDocumentMapping struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ClientId         primitive.ObjectID `bson:"clientId" json:"clientId"`
	ClientName       string             `bson:"clientName" json:"clientName"`
	DocumentId       primitive.ObjectID `bson:"documentId" json:"documentId"`
	DocumentName     string             `bson:"documentName" json:"documentName"`
	Status           string             `bson:"status" json:"status"`
	CreationDateTime time.Time          `bson:"creationDateTime" json:"creationDateTime"`
}

type CDMSearch struct {
	ClientName   string `json:"clientName"`
	DocumentName string `json:"documentName"`
}

type CDMRequest struct {
	Id           string `json:"id"`
	ClientName   string `json:"clientName"`
	DocumentName string `json:"documentName"`
	Status       string `json:"status"`
}
