package service

import (
	"context"
	"du-master/model"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocumentMaster struct {
	DocumentMasterCollection *mongo.Collection
}

// var DocumentCollection *mongo.Collection
var Documentctx = context.Background()

// func (e *DocumentMaster) DocumentConnect() {
// 	clientOptions := options.Client().ApplyURI(e.Server)
// 	client, err := mongo.Connect(Documentctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(Documentctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	DocumentCollection = client.Database(e.Database).Collection(e.Collection)
// }

func (e *DocumentMaster) DocumentInsert(DocumentRequest model.Document) (model.Document, error) {
	DocumentRequest = TrimDocumentDetails(DocumentRequest)
  
	isUnique, err := e.CheckDuplicateDocument(DocumentRequest.DocumentName, DocumentRequest.DocumentType)
	if err != nil {
		return DocumentRequest, err
	}
	if !isUnique {
		return DocumentRequest, errors.New("Document already present")
	}
	if err != nil {
		return DocumentRequest, err
	}

	DocumentRequest.Status = "A"
	DocumentRequest.CreationDateTime = time.Now()
	data, err := e.DocumentMasterCollection.InsertOne(Documentctx, DocumentRequest)

	if err != nil {
		fmt.Print(err.Error())
		return DocumentRequest, err
	}

	if oid, ok := data.InsertedID.(primitive.ObjectID); ok {

		DocumentRequest.ID = oid

	}

	return DocumentRequest, nil
}

func (e *DocumentMaster) DocumentFindById(id string) ([]*model.Document, error) {
	var Documents []*model.Document

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Documents, err
	}
	cur, err := e.DocumentMasterCollection.Find(Documentctx, bson.D{primitive.E{Key: "_id", Value: idHex}})

	if err != nil {
		return Documents, err
	}

	Documents, err = convertDbResultIntoDocumentStruct(cur)

	if err != nil {
		return Documents, err
	}

	if len(Documents) == 0 {
		return Documents, mongo.ErrNoDocuments
	}

	return Documents, nil
}

func (e *DocumentMaster) DocumentFindByName(name string) ([]*model.Document, error) {
	var Documents []*model.Document

	cur, err := e.DocumentMasterCollection.Find(Documentctx, bson.D{primitive.E{Key: "documentName", Value: name}})

	if err != nil {
		return Documents, err
	}

	Documents, err = convertDbResultIntoDocumentStruct(cur)

	if err != nil {
		return Documents, err
	}

	if len(Documents) == 0 {
		return Documents, mongo.ErrNoDocuments
	}

	return Documents, nil
}

func (e *DocumentMaster) SearchFilterOnDocument(DocumentRequest model.Document) ([]*model.Document, error) {
	var Documents []*model.Document
	DocumentRequest = TrimDocumentDetails(DocumentRequest)
	query := bson.D{}

	if DocumentRequest.DocumentName != "" {
		query = append(query, primitive.E{Key: "documentName", Value: DocumentRequest.DocumentName})
	}
	if DocumentRequest.DocumentType != "" {
		query = append(query, primitive.E{Key: "documentType", Value: DocumentRequest.DocumentType})
	}
	if DocumentRequest.Status != "" {
		query = append(query, primitive.E{Key: "status", Value: DocumentRequest.Status})
	}
	if DocumentRequest.DocumentTrainingStatus != "" {
		query = append(query, primitive.E{Key: "documentTrainingStatus", Value: DocumentRequest.DocumentTrainingStatus})
	}
	if DocumentRequest.DocumentDiscription != "" {
		query = append(query, primitive.E{Key: "documentDiscription", Value: DocumentRequest.DocumentDiscription})
	}

	cur, err := e.DocumentMasterCollection.Find(Documentctx, query)

	if err != nil {
		return Documents, err
	}

	Documents, err = convertDbResultIntoDocumentStruct(cur)

	if err != nil {
		return Documents, err
	}

	if len(Documents) == 0 {
		return Documents, mongo.ErrNoDocuments
	}

	return Documents, nil

}

func (e *DocumentMaster) Update(Document model.Document, id string) (bson.M, error) {
	var updatedDocument bson.M
	Document = TrimDocumentDetails(Document)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return updatedDocument, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: idHex}}

	update := bson.D{primitive.E{Key: "$set", Value: Document}}

	err = e.DocumentMasterCollection.FindOneAndUpdate(Documentctx, filter, update).Decode(&updatedDocument)
	if err != nil {
		return updatedDocument, err
	}

	if updatedDocument == nil {
		return updatedDocument, errors.New("data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

func (e *DocumentMaster) DeleteDocument(id string) (string, error) {

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: idHex}}

	res, err := e.DocumentMasterCollection.DeleteOne(Documentctx, filter)
	if err != nil {
		return "Document deletion unsuccessfu!", err
	}

	if res.DeletedCount == 0 {
		return "Document deletion unsuccessfu!", errors.New("No Document were deleted")
	}

	return "Document deletion successfull", err
}

func (e *DocumentMaster) ListDocuments() ([]string, error) {
	var finaldata []string

	filter := bson.D{}

	filter = append(filter, primitive.E{Key: "status", Value: "A"})
	cur, err := e.DocumentMasterCollection.Find(Documentctx, filter)
	if err != nil {
		return finaldata, err
	}

	Documents, err := convertDbResultIntoDocumentStruct(cur)

	if err != nil {
		return finaldata, err
	}

	if len(Documents) == 0 {
		return finaldata, mongo.ErrNoDocuments
	}

	for i := range Documents {
		DocumentName := Documents[i].DocumentName
		finaldata = append(finaldata, DocumentName)
	}

	return finaldata, nil
}

func convertDbResultIntoDocumentStruct(fetchDataCursor *mongo.Cursor) ([]*model.Document, error) {
	var finaldata []*model.Document
	for fetchDataCursor.Next(Documentctx) {
		var data model.Document
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}

func TrimDocumentDetails(DocumentRequest model.Document) model.Document {
	if DocumentRequest.DocumentName != "" {
		DocumentRequest.DocumentName = strings.TrimSpace(DocumentRequest.DocumentName)
	}
	if DocumentRequest.DocumentType != "" {
		DocumentRequest.DocumentType = strings.TrimSpace(DocumentRequest.DocumentType)
	}
	if DocumentRequest.DocumentTrainingStatus != "" {
		DocumentRequest.DocumentTrainingStatus = strings.TrimSpace(DocumentRequest.DocumentTrainingStatus)
	}
	return DocumentRequest
}

func (e *DocumentMaster) CheckDuplicateDocument(DocumentName string, DocumentType string) (bool, error) {
	cur, err := e.DocumentMasterCollection.Find(Documentctx, bson.D{{Key: "documentName", Value: DocumentName}, {Key: "documentType", Value: DocumentType}})
	Documents, err := convertDbResultIntoDocumentStruct(cur)

	if err != nil {
		return true, err
	}

	if len(Documents) == 0 {
		return true, nil
	}
	return false, err
}
