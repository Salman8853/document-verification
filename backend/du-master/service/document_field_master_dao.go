package service

import (
	"context"
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

type DocumentFieldMasterDAO struct {
	DocumentFieldCollection *mongo.Collection
}

var StatusActiveValue = "A"
var StatusDeleteValue = "I"
var FieldStatus = "status"
var FieldID = "_id"
var FieldDocumentFieldName = "documentFieldName"

func (e *DocumentFieldMasterDAO) Insert(request model.DocumentFieldMasterRequest) ([]model.DocumentFieldMaster, error) {
	var saveData model.DocumentFieldMaster
	var response []model.DocumentFieldMaster

	check := e.CheckIfAlreadyExists(request.DocumentFieldName)
	if check {
		return response, errors.New("already exists")
	}
	saveData.DocumentFieldName = request.DocumentFieldName
	saveData.Status = StatusActiveValue
	saveData.CreationDateTime = time.Now()

	data, err := e.DocumentFieldCollection.InsertOne(context.Background(), saveData)
	if err != nil {
		fmt.Println("got error while saving", err)
		return response, err
	}
	saveData.Id = data.InsertedID.(primitive.ObjectID)
	response = append(response, saveData)
	return response, err
}

func (e *DocumentFieldMasterDAO) FindByID(idStr string) ([]*model.DocumentFieldMaster, error) {
	var data []*model.DocumentFieldMaster

	id, err := ConvertStringIntoHex(idStr)
	if err != nil {
		log.Println(err)
		return data, err
	}
	filter := bson.D{primitive.E{Key: FieldID, Value: id}}
	filter = append(filter, primitive.E{Key: FieldStatus, Value: StatusActiveValue})
	fetchData, err := e.DocumentFieldCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("got error while fetching:", err)
		return data, err
	}
	data, err = ConvertResultIntoDocumentFieldMasterStruct(fetchData)
	if err != nil {
		log.Println("Got Error While converting data :", err)
		return data, err
	}
	if len(data) == 0 {
		return data, mongo.ErrNoDocuments
	}
	return data, err
}

func (e *DocumentFieldMasterDAO) DeactivatedByID(idStr string) (string, error) {
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
	e.DocumentFieldCollection.FindOneAndUpdate(context.Background(), filter, query, opts).Decode(&updateDocument)
	if updateDocument == nil {
		log.Println("got error while deactivating :", err)
		return "", errors.New("unable to deactivate")
	}

	return "deactivated successfully", err
}

func (e *DocumentFieldMasterDAO) FindAll() ([]*model.DocumentFieldMaster, error) {
	var data []*model.DocumentFieldMaster

	filter := bson.D{}
	fetchData, err := e.DocumentFieldCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("got error while fetching :", err)
		return data, err
	}
	data, err = ConvertResultIntoDocumentFieldMasterStruct(fetchData)
	if err != nil {
		log.Println("got error while converting data :", err)
		return data, err
	}
	return data, err
}

func (e *DocumentFieldMasterDAO) SearchFilter(request model.DocumentFieldMasterSearchRequest) ([]*model.DocumentFieldMaster, error) {
	var data []*model.DocumentFieldMaster

	filter := bson.D{}
	if request.DocumentFieldName != "" {
		filter = append(filter, primitive.E{Key: FieldDocumentFieldName, Value: bson.M{"$regex": request.DocumentFieldName}})
	}
	if request.Id != "" {
		id, err := ConvertStringIntoHex(request.Id)
		if err != nil {
			log.Println(err)
			return data, err
		}
		filter = append(filter, primitive.E{Key: FieldID, Value: id})
	}
	if request.Status != "" {
		filter = append(filter, primitive.E{Key: FieldStatus, Value: request.Status})
	}

	fetchData, err := e.DocumentFieldCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("got error while fetching :", err)
		return data, err
	}
	data, err = ConvertResultIntoDocumentFieldMasterStruct(fetchData)
	if err != nil {
		log.Println("got error while converting data :", err)
		return data, err
	}
	if len(data) == 0 {
		return data, mongo.ErrNoDocuments
	}

	return data, err
}

func (e *DocumentFieldMasterDAO) UpdateByID(idStr string, request model.DocumentFieldMasterSearchRequest) (bson.M, error) {
	var updateDocument bson.M

	check := e.CheckIfAlreadyExists(request.DocumentFieldName)
	if check {
		return updateDocument, errors.New("already exists")
	}
	id, err := ConvertStringIntoHex(idStr)
	if err != nil {
		log.Println(err)
		return updateDocument, err
	}
	filter := bson.D{primitive.E{Key: FieldID, Value: id}}

	updateQuery := bson.D{}
	if request.DocumentFieldName != "" {
		updateQuery = append(updateQuery, primitive.E{Key: FieldDocumentFieldName, Value: request.DocumentFieldName})
	}
	if request.Status != "" {
		updateQuery = append(updateQuery, primitive.E{Key: FieldStatus, Value: request.Status})
	}

	query := bson.D{primitive.E{Key: "$set", Value: updateQuery}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	e.DocumentFieldCollection.FindOneAndUpdate(context.Background(), filter, query, opts).Decode(&updateDocument)
	if updateDocument == nil {
		return updateDocument, mongo.ErrNoDocuments
	}

	return updateDocument, err
}

func (e *DocumentFieldMasterDAO) CheckIfAlreadyExists(documentFieldName string) bool {
	filter := bson.D{primitive.E{Key: FieldStatus, Value: StatusActiveValue}}
	filter = append(filter, primitive.E{Key: FieldDocumentFieldName, Value: documentFieldName})
	cur, err := e.DocumentFieldCollection.Find(context.Background(), filter)
	if err != nil {
		return true
	}
	if cur.Next(context.Background()) {
		return true
	}
	return false
}

func ConvertStringIntoHex(idStr string) (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return id, err
	}
	return id, err
}

func ConvertResultIntoDocumentFieldMasterStruct(fetchDataCursor *mongo.Cursor) ([]*model.DocumentFieldMaster, error) {
	var finaldata []*model.DocumentFieldMaster
	for fetchDataCursor.Next(context.Background()) {
		var data model.DocumentFieldMaster
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}

func (e *DocumentFieldMasterDAO) FindAllDocFieldName() ([]string, error) {
	var response []string
	var data []*model.DocumentFieldMaster

	filter := bson.D{primitive.E{Key: FieldStatus, Value: StatusActiveValue}}
	fetchData, err := e.DocumentFieldCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("got error while fetching :", err)
		return response, err
	}
	data, err = ConvertResultIntoDocumentFieldMasterStruct(fetchData)
	if err != nil {
		log.Println("got error while converting data :", err)
		return response, err
	}
	for i := range data {
		response = append(response, data[i].DocumentFieldName)
	}
	return response, err
}
