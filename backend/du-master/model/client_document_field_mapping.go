package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientDocumentFieldMapping struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ClientMaster      primitive.ObjectID `bson:"clientMaster,omitempty" json:"clientMaster,omitempty"`
	DocumentFieldName primitive.ObjectID `bson:"documentFieldName,omitempty" json:"documentFieldName,omitempty"`
	DocumentName      primitive.ObjectID `bson:"documentName,omitempty" json:"documentName,omitempty"`
	Status            string             `bson:"status,omitempty" json:"status,omitempty"`
	CreationDateTime  time.Time          `bson:"creationDateTime,omitempty" json:"creationDateTime,omitempty"`
}

type ClientDocumentFieldRequest struct {
	ClientName        string `json:"clientName,omitempty"`
	DocumentFieldName string `json:"documentFieldName,omitempty"`
	DocumentName      string `json:"documentName,omitempty"`
	Status            string `json:"status,omitempty"`
}

type ClientDocFieldSearchFilterRequest struct {
	Id                string `json:"id,omitempty"`
	ClientName        string `json:"clientName,omitempty"`
	DocumentFieldName string `json:"documentFieldName,omitempty"`
	DocumentName      string `json:"documentName,omitempty"`
	Status            string `json:"status,omitempty"`
}
