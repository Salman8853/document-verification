package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentTransactionDetails struct {
	Id               primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	TransactionId    string              `bson:"transactionId,omitempty" json:"transactionId,omitempty"`
	ReferenceId      string              `bson:"referenceId,omitempty" json:"referenceId,omitempty"`
	Client           string              `bson:"client,omitempty" json:"client,omitempty"`
	FirstName        string              `bson:"firstName,omitempty" json:"firstName,omitempty"`
	MiddleName       string              `bson:"middleName,omitempty" json:"middleName,omitempty"`
	LastName         string              `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Status           string              `bson:"status,omitempty" json:"status,omitempty"`
	Source           string              `bson:"source,omitempty" json:"source,omitempty"`
	Email            string              `bson:"email,omitempty" json:"email,omitempty"`
	ContactNumber    string              `bson:"contactNumber,omitempty" json:"contactNumber,omitempty"`
	Dob              time.Time           `bson:"dob,omitempty" json:"dob,omitempty"`
	UploadedDocs     []UploadedDocuments `bson:"uploadedDocs" json:"uploadedDocs"`
	CreationDateTime time.Time           `bson:"creationDateTime,omitempty" json:"creationDateTime,omitempty"`
}

type DocumentTransactionRequest struct {
	FirstName     string `json:"firstName,omitempty"`
	MiddleName    string `json:"middleName,omitempty"`
	LastName      string `json:"lastName,omitempty"`
	Dob           string `json:"dob,omitempty"`
	Email         string `bson:"email,omitempty" json:"email,omitempty"`
	DocPath       string `bson:"docPath,omitempty" json:"docPath,omitempty"`
	ContactNumber string `bson:"contactNumber,omitempty" json:"contactNumber,omitempty"`
}

type ExcelTransactionRequest struct {
	FIRSTNAME     string `json:"FIRSTNAME,omitempty"`
	LASTNAME      string `json:"LASTNAME,omitempty"`
	MIDDLENAME    string `json:"MIDDLENAME,omitempty"`
	CONTACTNUMBER string `json:"CONTACTNUMBER,omitempty"`
	DOB           string `json:"DOB,omitempty"`
	DOCUMENTPATH  string `json:"DOCUMENTPATH,omitempty"`
	EMAIL         string `json:"EMAIL,omitempty"`
}

type ExcelerrList struct {
	RowNo    int      `json:"rowNo,omitempty"`
	ExcelErr []string `json:"excelErr,omitempty"`
}
