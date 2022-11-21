package model

import (
	"time"
)

type UploadedDocuments struct {
	DocumentName string    `bson:"documentName" json:"documentName"`
	Size         int64     `bson:"size" json:"size"`
	DocumentPath string    `bson:"documentPath " json:"documentPath "`
	DateCreated  time.Time `bson:"dateCreated" json:"dateCreated"`
}
