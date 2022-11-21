package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ClientName       string             `bson:"clientName" json:"clientName"`
	ClientCode       string             `bson:"clientCode" json:"clientCode"`
	Address          string             `bson:"address" json:"address"`
	LogoPath         string             `bson:"logoPath" json:"logoPath"`
	FunctionalEntity string             `bson:"functionalEntity" json:"functionalEntity"`
	Status           string             `bson:"status" json:"status"`
	CreationDateTime time.Time          `bson:"creationDateTime" json:"creationDateTime"`
}
type ListClient struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ClientName string             `bson:"clientName" json:"clientName"`
	ClientCode string             `bson:"clientCode" json:"clientCode"`
}
