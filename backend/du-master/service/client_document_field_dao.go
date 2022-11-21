package service

import (
	"context"
	"du-master/config"
	"du-master/model"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const clientMasterColl = "clientMaster"
const docFieldMasterColl = "documentFieldMaster"
const documentMasterColl = "DocumentMaster"

type ClientDocFieldMappingDAO struct {
	ClientDocFieldMappingCollection *mongo.Collection
}

func (e *ClientDocFieldMappingDAO) Insert(request model.ClientDocumentFieldRequest) ([]model.ClientDocumentFieldMapping, error) {
	var saveData model.ClientDocumentFieldMapping
	var response []model.ClientDocumentFieldMapping
	clientId, err := FetchIdsByClientName(request.ClientName)
	if err != nil {
		return response, err
	}
	documentFieldNameId, err := FetchIdsByDocumentFieldName(request.DocumentFieldName)
	if err != nil {
		return response, err
	}
	documentNameId, err := FetchIdsByDocumentName(request.DocumentName)
	if err != nil {
		return response, err
	}
	check := e.CheckIfAlreadyExistsInCDFM(clientId, documentFieldNameId, documentNameId)
	if check {
		return response, errors.New("Already Exists")
	}
	saveData.ClientMaster = clientId
	saveData.DocumentFieldName = documentFieldNameId
	saveData.DocumentName = documentNameId
	saveData.Status = StatusActiveValue
	saveData.CreationDateTime = time.Now()

	data, err := e.ClientDocFieldMappingCollection.InsertOne(context.Background(), saveData)
	if err != nil {
		fmt.Println("Got Error While Saving", err)
		return response, err
	}
	saveData.Id = data.InsertedID.(primitive.ObjectID)
	response = append(response, saveData)
	return response, err
}

func (e *ClientDocFieldMappingDAO) FindByID(idStr string) ([]*model.ClientDocumentFieldMapping, error) {
	var data []*model.ClientDocumentFieldMapping

	id, err := ConvertStringIntoHex(idStr)
	if err != nil {
		log.Println(err)
		return data, err
	}
	filter := bson.D{primitive.E{Key: FieldID, Value: id}}
	filter = append(filter, primitive.E{Key: FieldStatus, Value: StatusActiveValue})
	fetchData, err := e.ClientDocFieldMappingCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Got Error While Saving :", err)
		return data, err
	}
	data, err = ConvertResultIntoClientDocFieldMappingStruct(fetchData)
	if err != nil {
		log.Println("Got Error While converting data :", err)
		return data, err
	}
	if len(data) == 0 {
		return data, mongo.ErrNoDocuments
	}
	return data, err
}

func (e *ClientDocFieldMappingDAO) DeactivatedByID(idStr string) (string, error) {
	var updateDocument bson.M
	id, err := ConvertStringIntoHex(idStr)
	if err != nil {
		log.Println(err)
		return "", err
	}
	filter := bson.D{primitive.E{Key: FieldID, Value: id}}
	filter = append(filter, primitive.E{Key: FieldStatus, Value: StatusActiveValue})
	updateQuery := bson.D{primitive.E{Key: FieldStatus, Value: StatusDeleteValue}}
	query := bson.D{primitive.E{Key: "$set", Value: updateQuery}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	e.ClientDocFieldMappingCollection.FindOneAndUpdate(context.Background(), filter, query, opts).Decode(&updateDocument)
	if updateDocument == nil {
		log.Println("Got Error While Deactivating :", err)
		return "", errors.New("Unable to Deactivate")
	}

	return "Deactivated successfully", err
}

func (e *ClientDocFieldMappingDAO) FindAll() ([]*model.ClientDocumentFieldMapping, error) {
	var data []*model.ClientDocumentFieldMapping

	filter := bson.D{}
	fetchData, err := e.ClientDocFieldMappingCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Got Error While Fetching :", err)
		return data, err
	}
	data, err = ConvertResultIntoClientDocFieldMappingStruct(fetchData)
	if err != nil {
		log.Println("Got Error While converting data :", err)
		return data, err
	}
	return data, err
}

func (e *ClientDocFieldMappingDAO) SearchFilter(request model.ClientDocFieldSearchFilterRequest) ([]*model.ClientDocumentFieldMapping, error) {
	var data []*model.ClientDocumentFieldMapping

	filter := bson.D{}
	if request.Id != "" {
		id, err := ConvertStringIntoHex(request.Id)
		if err != nil {
			log.Println(err)
			return data, err
		}
		filter = append(filter, primitive.E{Key: FieldID, Value: id})
	}
	if request.DocumentFieldName != "" {
		docFMId, err := FetchIdsByDocumentFieldName(request.DocumentFieldName)
		if err != nil {
			log.Println(err)
			return data, err
		}
		filter = append(filter, primitive.E{Key: FieldDocumentFieldName, Value: docFMId})
	}
	if request.DocumentName != "" {
		dmId, err := FetchIdsByDocumentName(request.DocumentName)
		if err != nil {
			log.Println(err)
			return data, err
		}
		filter = append(filter, primitive.E{Key: "documentName", Value: dmId})
	}
	if request.ClientName != "" {
		cmId, err := FetchIdsByClientName(request.ClientName)
		if err != nil {
			log.Println(err)
			return data, err
		}
		filter = append(filter, primitive.E{Key: "clientName", Value: cmId})
	}
	if request.Status != "" {
		filter = append(filter, primitive.E{Key: FieldStatus, Value: request.Status})
	}

	fetchData, err := e.ClientDocFieldMappingCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Got Error While Fetching :", err)
		return data, err
	}
	data, err = ConvertResultIntoClientDocFieldMappingStruct(fetchData)
	if err != nil {
		log.Println("Got Error While converting data :", err)
		return data, err
	}
	if len(data) == 0 {
		return data, mongo.ErrNoDocuments
	}
	return data, err
}

func (e *ClientDocFieldMappingDAO) UpdateByID(idStr string, request model.ClientDocumentFieldRequest) (bson.M, error) {
	var updateDocument bson.M

	id, err := ConvertStringIntoHex(idStr)
	if err != nil {
		log.Println(err)
		return updateDocument, err
	}
	filter := bson.D{primitive.E{Key: FieldID, Value: id}}

	updateQuery := bson.D{}
	if request.DocumentFieldName != "" {
		docFMId, err := FetchIdsByDocumentFieldName(request.DocumentFieldName)
		if err != nil {
			log.Println(err)
			return updateDocument, err
		}
		updateQuery = append(updateQuery, primitive.E{Key: FieldDocumentFieldName, Value: docFMId})
	}
	if request.DocumentName != "" {
		dmId, err := FetchIdsByDocumentName(request.DocumentName)
		if err != nil {
			log.Println(err)
			return updateDocument, err
		}
		updateQuery = append(updateQuery, primitive.E{Key: "documentName", Value: dmId})
	}
	if request.ClientName != "" {
		cmId, err := FetchIdsByClientName(request.ClientName)
		if err != nil {
			log.Println(err)
			return updateDocument, err
		}
		updateQuery = append(updateQuery, primitive.E{Key: "clientName", Value: cmId})
	}
	if request.Status != "" {
		updateQuery = append(updateQuery, primitive.E{Key: FieldStatus, Value: request.Status})
	}

	query := bson.D{primitive.E{Key: "$set", Value: updateQuery}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	e.ClientDocFieldMappingCollection.FindOneAndUpdate(context.Background(), filter, query, opts).Decode(&updateDocument)
	if updateDocument == nil {
		return updateDocument, mongo.ErrNoDocuments
	}

	return updateDocument, err
}

func FetchIdsByClientName(clientName string) (primitive.ObjectID, error) {
	ClientMasterCollection := config.GetCollection(config.DB, clientMasterColl)
	clientCur, err := ClientMasterCollection.Find(context.Background(), bson.D{primitive.E{Key: "clientName", Value: clientName}, primitive.E{Key: FieldStatus, Value: StatusActiveValue}})
	if err != nil {
		return primitive.NilObjectID, err
	}

	var clientModel []model.Client
	clientCur.All(context.Background(), &clientModel)
	fmt.Println(clientModel)
	if len(clientModel) == 0 {
		return primitive.NilObjectID, errors.New("Invalid ClientName")
	}
	return clientModel[0].ID, err
}

func FetchIdsByDocumentFieldName(documentFieldName string) (primitive.ObjectID, error) {
	DocumentFieldMasterCollection := config.GetCollection(config.DB, docFieldMasterColl)
	documentFieldCur, err := DocumentFieldMasterCollection.Find(context.Background(), bson.D{primitive.E{Key: "documentFieldName", Value: documentFieldName}, primitive.E{Key: FieldStatus, Value: StatusActiveValue}})
	if err != nil {
		return primitive.NilObjectID, err
	}
	var documentFieldModel []model.DocumentFieldMaster
	documentFieldCur.All(context.Background(), &documentFieldModel)
	if len(documentFieldModel) == 0 {
		return primitive.NilObjectID, errors.New("Invalid DocumentFieldName")
	}
	return documentFieldModel[0].Id, err
}

func FetchIdsByDocumentName(documentName string) (primitive.ObjectID, error) {
	DocumentMasterCollection := config.GetCollection(config.DB, documentMasterColl)
	documentCur, err := DocumentMasterCollection.Find(context.Background(), bson.D{primitive.E{Key: "documentName", Value: documentName}, primitive.E{Key: FieldStatus, Value: StatusActiveValue}})
	if err != nil {
		return primitive.NilObjectID, err
	}
	var docNameModel []model.Document
	documentCur.All(context.Background(), &docNameModel)
	if len(docNameModel) == 0 {
		return primitive.NilObjectID, errors.New("Invalid DocumentName")
	}
	return docNameModel[0].ID, err
}

func (e *ClientDocFieldMappingDAO) CheckIfAlreadyExistsInCDFM(clientNameId, documentFieldNameId, documentNameId primitive.ObjectID) bool {
	filter := bson.D{primitive.E{Key: FieldStatus, Value: StatusActiveValue}}
	filter = append(filter, primitive.E{Key: "clientMaster", Value: clientNameId})
	filter = append(filter, primitive.E{Key: "documentFieldName", Value: documentFieldNameId})
	filter = append(filter, primitive.E{Key: "documentName", Value: documentNameId})
	cur, err := e.ClientDocFieldMappingCollection.Find(context.Background(), filter)
	if err != nil {
		return true
	}
	if cur.Next(context.Background()) {
		return true
	}
	return false
}

func ConvertResultIntoClientDocFieldMappingStruct(fetchDataCursor *mongo.Cursor) ([]*model.ClientDocumentFieldMapping, error) {
	var finaldata []*model.ClientDocumentFieldMapping
	for fetchDataCursor.Next(context.Background()) {
		var data model.ClientDocumentFieldMapping
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}
