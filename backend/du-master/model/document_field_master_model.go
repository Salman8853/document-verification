package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentFieldMaster struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	DocumentFieldName string             `bson:"documentFieldName,omitempty" json:"documentFieldName,omitempty"`
	Status            string             `bson:"status,omitempty" json:"status,omitempty"`
	CreationDateTime  time.Time          `bson:"creationDateTime,omitempty" json:"creationDateTime,omitempty"`
}

type DocumentFieldMasterRequest struct {
	DocumentFieldName string `json:"documentFieldName"`
}

type DocumentFieldMasterSearchRequest struct {
	Id                string `json:"id,omitempty"`
	DocumentFieldName string `json:"documentFieldName,omitempty"`
	Status            string `json:"status,omitempty"`
}
