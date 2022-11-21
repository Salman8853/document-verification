package service

import (
	"context"
	"du-master/model"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var document_field_mapping = config.GetCollection(config.DB, "document_field_mapping")

type DocumentFieldMappingDAO struct {
	DocumentFieldMappingCollection *mongo.Collection
}

func (dfm *DocumentFieldMappingDAO) Insert(dfMapping model.DocumentFieldMapping) (interface{}, error) {
	dfMapping.Status = "A"
	dfMapping.CreationDateTime = time.Now()

	result, err := dfm.DocumentFieldMappingCollection.InsertOne(context.Background(), dfMapping)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result.InsertedID, nil
}

func (dfm *DocumentFieldMappingDAO) FindAll() (interface{}, error) {
	cur, err := dfm.DocumentFieldMappingCollection.Find(context.Background(), bson.D{})

	var result []model.DocumentFieldMapping

	if err != nil {
		log.Println(err)
		return nil, err
	}

	cur.All(context.Background(), &result)

	return result, nil
}

func (dfm *DocumentFieldMappingDAO) FindById(id string) (interface{}, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	cur, err := dfm.DocumentFieldMappingCollection.Find(context.Background(), bson.M{"_id": objectId})

	var result []model.DocumentFieldMapping

	if err != nil {
		log.Println(err)
		return nil, err
	}

	cur.All(context.Background(), &result)

	return result, nil
}

func (dfm *DocumentFieldMappingDAO) DeleteById(id string) (interface{}, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	res, err := dfm.DocumentFieldMappingCollection.DeleteOne(context.Background(), bson.M{"_id": objectId})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (dfm *DocumentFieldMappingDAO) UpdateById(id string, dfmc model.DocumentFieldMapping) (interface{}, error) {
	var updatedDocument bson.M

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	filter := bson.M{"_id": objectId}

	update := bson.D{primitive.E{Key: "$set", Value: dfmc}}

	err = dfm.DocumentFieldMappingCollection.FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedDocument)

	if err != nil {
		return updatedDocument, err
	}

	if updatedDocument == nil {
		return updatedDocument, errors.New("data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

func (dfm *DocumentFieldMappingDAO) SearchFilterOnDocument(dfmc model.DocumentFieldMapping) (interface{}, error) {
	var result []*model.DocumentFieldMapping
	query := bson.D{}

	if dfmc.DocumentName != "" {
		query = append(query, primitive.E{Key: "documentName", Value: dfmc.DocumentName})
	}

	if dfmc.Status != "" {
		query = append(query, primitive.E{Key: "status", Value: dfmc.Status})
	}

	if dfmc.DocumentFieldName != "" {
		query = append(query, primitive.E{Key: "documentFieldName", Value: dfmc.DocumentFieldName})
	}

	cur, err := dfm.DocumentFieldMappingCollection.Find(context.Background(), query)

	if err != nil {
		return result, err
	}

	cur.All(context.Background(), &result)

	if len(result) == 0 {
		return result, mongo.ErrNoDocuments
	}

	return result, nil
}

func (dfm *DocumentFieldMappingDAO) IsAlreadyPresent(dfmc model.DocumentFieldMapping) (bool, error) {
	var result []model.DocumentFieldMapping

	filter := bson.D{}

	filter = append(filter, primitive.E{Key: "documentName", Value: dfmc.DocumentName})
	filter = append(filter, primitive.E{Key: "documentFieldName", Value: dfmc.DocumentFieldName})

	cur, err := dfm.DocumentFieldMappingCollection.Find(context.Background(), filter)

	if err != nil {
		return false, err
	}

	cur.All(context.Background(), &result)

	if len(result) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
