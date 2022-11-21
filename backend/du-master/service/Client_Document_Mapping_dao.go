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

type ClientDocumentMappingDAO struct {
	ClientDocumentMappingCollection *mongo.Collection
}

const clientMasterCollec = "clientMaster"
const documentMasterCollec = "documentMaster"

// ========================Insert_CDM================================
func (e *ClientDocumentMappingDAO) InsertCliDocMapping(request model.CDMRequest) ([]model.ClientDocumentMapping, error) {
	var saveData model.ClientDocumentMapping
	var response []model.ClientDocumentMapping
	clientId, err := FetchIdByClientName(request.ClientName)
	if err != nil {
		return response, err
	}
	documentId, err := FetchIdByDocumentName(request.DocumentName)
	if err != nil {
		return response, err
	}
	check := e.CheckIfAlreadyExistsInCDM(clientId, documentId)
	if check {
		return response, errors.New("Already Exists")
	}
	saveData.ClientId = clientId
	saveData.ClientName = request.ClientName
	saveData.DocumentId = documentId
	saveData.DocumentName = request.DocumentName
	saveData.Status = "A"
	saveData.CreationDateTime = time.Now()

	data, err := e.ClientDocumentMappingCollection.InsertOne(context.Background(), saveData)
	if err != nil {
		fmt.Println("Got Error While Saving", err)
		return response, err
	}
	saveData.ID = data.InsertedID.(primitive.ObjectID)
	response = append(response, saveData)
	return response, err
}

// =============fetchIDByClientName=================
func FetchIdByClientName(clientName string) (primitive.ObjectID, error) {
	ClientMasterCollection := config.GetCollection(config.DB, clientMasterCollec)
	clientCur, err := ClientMasterCollection.Find(context.Background(), bson.D{primitive.E{Key: "clientName", Value: clientName}, primitive.E{Key: "status", Value: "A"}})
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

// =============fetchDocIDByDocName==============
func FetchIdByDocumentName(documentName string) (primitive.ObjectID, error) {
	DocumentMasterCollection := config.GetCollection(config.DB, documentMasterCollec)
	documentCur, err := DocumentMasterCollection.Find(context.Background(), bson.D{primitive.E{Key: "documentName", Value: documentName}, primitive.E{Key: "status", Value: "A"}})
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

// ================checkIfAlreadyExit============
func (e *ClientDocumentMappingDAO) CheckIfAlreadyExistsInCDM(clientId, documentId primitive.ObjectID) bool {
	filter := bson.D{primitive.E{Key: "status", Value: "A"}}
	filter = append(filter, primitive.E{Key: "clientId", Value: clientId})
	filter = append(filter, primitive.E{Key: "documentId", Value: documentId})
	cur, err := e.ClientDocumentMappingCollection.Find(context.Background(), filter)
	if err != nil {
		return true
	}
	if cur.Next(context.Background()) {
		return true
	}
	return false
}

// ========================Find_CDM_By_ID==========================
func (e *ClientDocumentMappingDAO) CliDocMapFindById(id string) ([]*model.ClientDocumentMapping, error) {
	var clientDocMap []*model.ClientDocumentMapping

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return clientDocMap, err
	}
	cur, err := e.ClientDocumentMappingCollection.Find(context.Background(), bson.D{primitive.E{Key: "_id", Value: idHex}})

	if err != nil {
		return clientDocMap, errors.New("unable to query db")
	}

	clientDocMap, err = ConvertDbResultIntoCDMStruct(cur)

	if err != nil {
		return clientDocMap, err
	}

	if len(clientDocMap) == 0 {
		return clientDocMap, mongo.ErrNoDocuments
	}

	return clientDocMap, nil
}

// ==============Search Filter========================
func (e *ClientDocumentMappingDAO) SearchFilterOnDocument(request model.CDMRequest) ([]*model.ClientDocumentMapping, error) {
	var data []*model.ClientDocumentMapping

	filter := bson.D{}
	if request.Id != "" {
		id, err := ConvertStringIntoHex(request.Id)
		if err != nil {
			log.Println(err)
			return data, err
		}
		filter = append(filter, primitive.E{Key: "_id", Value: id})
	}

	if request.DocumentName != "" {
		dmId, err := FetchIdByDocumentName(request.DocumentName)
		if err != nil {
			log.Println(err)
			return data, err
		}
		filter = append(filter, primitive.E{Key: "documentId", Value: dmId})
	}
	if request.ClientName != "" {
		cmId, err := FetchIdByClientName(request.ClientName)
		if err != nil {
			log.Println(err)
			return data, err
		}
		filter = append(filter, primitive.E{Key: "clientId", Value: cmId})
	}

	filter = append(filter, primitive.E{Key: "status", Value: "A"})

	fetchData, err := e.ClientDocumentMappingCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Got Error While Fetching :", err)
		return data, err
	}
	data, err = ConvertDbResultIntoCDMStruct(fetchData)
	if err != nil {
		log.Println("Got Error While converting data :", err)
		return data, err
	}
	if len(data) == 0 {
		return data, mongo.ErrNoDocuments
	}
	return data, err
}

// ==========================Update_CDM_By_ID===============================
func (e *ClientDocumentMappingDAO) Update(cdm model.CDMRequest, id string) (bson.M, error) {
	var updatedCDM bson.M

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return updatedCDM, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: idHex}}

	updateQuery := bson.D{}

	if cdm.ClientName != "" {
		cliId, err := FetchIdsByClientName(cdm.ClientName)
		if err != nil {
			return updatedCDM, err
		}
		updateQuery = append(updateQuery, primitive.E{Key: "cientId", Value: cliId})
	}
	if cdm.DocumentName != "" {
		docId, err := FetchIdsByDocumentName(cdm.DocumentName)
		if err != nil {
			return updatedCDM, err
		}
		updateQuery = append(updateQuery, primitive.E{Key: "documentId", Value: docId})
	}
	if cdm.Status != "" {
		updateQuery = append(updateQuery, primitive.E{Key: "status", Value: cdm.Status})
	}

	query := bson.D{primitive.E{Key: "$set", Value: updateQuery}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = e.ClientDocumentMappingCollection.FindOneAndUpdate(context.Background(), filter, query, opts).Decode(&updatedCDM)
	if err != nil {
		return updatedCDM, err
	}

	if updatedCDM == nil {
		return updatedCDM, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedCDM, nil
}

// ==========================Delete_CDM_By_ID===========================
func (e *ClientDocumentMappingDAO) DeactivateCDM(id string) (bson.M, error) {
	var updatedCDM bson.M
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return updatedCDM, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: idHex}}
	filter = append(filter, primitive.E{Key: "status", Value: "A"})
	updateQuery := bson.D{primitive.E{Key: "status", Value: "I"}}
	query := bson.D{primitive.E{Key: "$set", Value: updateQuery}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = e.ClientDocumentMappingCollection.FindOneAndUpdate(context.Background(), filter, query, opts).Decode(&updatedCDM)
	if err != nil {
		return updatedCDM, err
	}

	if updatedCDM == nil {
		return updatedCDM, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedCDM, nil
}

func (e *ClientDocumentMappingDAO) FindAll() ([]*model.ClientDocumentMapping, error) {
	var data []*model.ClientDocumentMapping

	filter := bson.D{}
	fetchData, err := e.ClientDocumentMappingCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Got Error While Fetching :", err)
		return data, err
	}
	data, err = ConvertDbResultIntoCDMStruct(fetchData)
	if err != nil {
		log.Println("Got Error While converting data :", err)
		return data, err
	}
	return data, err
}

// ======================Struct_Conv..=====================
func ConvertDbResultIntoCDMStruct(fetchDataCursor *mongo.Cursor) ([]*model.ClientDocumentMapping, error) {
	var finaldata []*model.ClientDocumentMapping
	for fetchDataCursor.Next(context.Background()) {
		var data model.ClientDocumentMapping
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}
