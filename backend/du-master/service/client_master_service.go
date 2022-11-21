package service

import (
	"context"
	"du-master/model"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientMaster struct {
	ClientCollection *mongo.Collection
}

var ActiveValue = "A"

func (cm *ClientMaster) Insert(clientData model.Client) (interface{}, error) {

	clientData = TrimClientDetails(clientData)

	check, err := cm.ValidateByClientName(clientData.ClientName)
	if err != nil {
		return clientData, err
	}
	if check {
		return clientData, errors.New("Client name already present")
	}

	found, err := cm.ValidateByClientCode(clientData.ClientCode)
	if err != nil {
		return clientData, err
	}
	if found {
		return clientData, errors.New("Client code already present")
	}

	clientData.CreationDateTime = time.Now()
	result, err := cm.ClientCollection.InsertOne(context.Background(), clientData)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	clientData.ID = result.InsertedID.(primitive.ObjectID)
	return clientData, nil
}

func (cm *ClientMaster) ValidateByClientName(clientName string) (bool, error) {

	cur, err := cm.ClientCollection.Find(context.Background(), bson.M{"clientName": clientName})

	if err != nil {
		return true, err
	}
	if cur.Next(context.Background()) {
		return true, err
	}
	return false, err
}

func (cm *ClientMaster) ValidateByClientCode(clientCode string) (bool, error) {
	cur, err := cm.ClientCollection.Find(context.Background(), bson.M{"clientCode": clientCode})

	if err != nil {
		return true, err
	}
	if cur.Next(context.Background()) {
		return true, err
	}
	return false, err
}

func (cm *ClientMaster) FindAll() (interface{}, error) {
	filter := bson.D{{Key: "status", Value: bson.M{"$ne": "I"}}}
	cur, err := cm.ClientCollection.Find(context.Background(), filter)

	var result []model.Client

	if err != nil {
		log.Println(err)
		return nil, err
	}

	cur.All(context.Background(), &result)

	return result, nil
}

func (cm *ClientMaster) FindByClientId(id string) (interface{}, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	cur, err := cm.ClientCollection.Find(context.Background(), bson.M{"_id": objectId})

	var result []model.Client

	if err != nil {
		log.Println(err)
		return nil, err
	}

	cur.All(context.Background(), &result)

	return result, nil
}

func (cm *ClientMaster) FindByClientName(name string) (interface{}, error) {
	cur, err := cm.ClientCollection.Find(context.Background(), bson.M{"clientName": name})

	var result []model.Client

	if err != nil {
		log.Println(err)
		return nil, err
	}

	cur.All(context.Background(), &result)

	return result, nil
}

func (cm *ClientMaster) SearchFilterOnClient(clientReq model.Client) ([]*model.Client, error) {
	var clientData []*model.Client
	clientReq = TrimClientDetails(clientReq)
	query := bson.D{}

	if clientReq.ClientName != "" {
		query = append(query, primitive.E{Key: "clientName", Value: clientReq.ClientName})
	}
	if clientReq.ClientCode != "" {
		query = append(query, primitive.E{Key: "clientCode", Value: clientReq.ClientCode})
	}
	if clientReq.FunctionalEntity != "" {
		query = append(query, primitive.E{Key: "functionalEntity", Value: clientReq.FunctionalEntity})
	}
	if clientReq.LogoPath != "" {
		query = append(query, primitive.E{Key: "logoPath", Value: clientReq.LogoPath})
	}
	if clientReq.Status != "" {
		query = append(query, primitive.E{Key: "status", Value: clientReq.Status})
	}
	if clientReq.Address != "" {
		query = append(query, primitive.E{Key: "address", Value: clientReq.Address})
	}

	cur, err := cm.ClientCollection.Find(context.Background(), query)

	if err != nil {
		return clientData, errors.New("unable to query db")
	}

	cur.All(context.Background(), &clientData)

	if len(clientData) == 0 {
		return clientData, mongo.ErrNoDocuments
	}

	return clientData, nil

}

func (cm *ClientMaster) UpdateByClientId(clientReq model.Client, idStr string) (bson.M, error) {

	var updatedDocument bson.M

	clientReq = TrimClientDetails(clientReq)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return updatedDocument, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	UpdateQuery := bson.D{}
	if clientReq.ClientName != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "clientName", Value: clientReq.ClientName})
	}
	if clientReq.ClientCode != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "clientCode", Value: clientReq.ClientCode})
	}
	if clientReq.FunctionalEntity != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "functionalEntity", Value: clientReq.FunctionalEntity})
	}
	if clientReq.LogoPath != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "logoPath", Value: clientReq.LogoPath})
	}
	if clientReq.Status != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "status", Value: clientReq.Status})
	}
	if clientReq.Address != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "address", Value: clientReq.Address})
	}

	update := bson.D{{Key: "$set", Value: UpdateQuery}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	r := cm.ClientCollection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&updatedDocument)
	if r != nil {
		return updatedDocument, r
	}
	fmt.Println(updatedDocument)
	if updatedDocument == nil {
		return updatedDocument, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

func (cm *ClientMaster) DeactivateClient(idStr string) (bson.M, error) {
	var updatedDocument bson.M
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return updatedDocument, err
	}
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "_id", Value: id}},
				bson.D{{Key: "status", Value: bson.M{"$ne": "I"}}},
			},
		},
	}
	update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "status", Value: "I"}}}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	r := cm.ClientCollection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&updatedDocument)
	if r != nil {
		return updatedDocument, r
	}

	if updatedDocument == nil {
		return updatedDocument, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

func (cm *ClientMaster) DeleteById(id string) (interface{}, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	res, err := cm.ClientCollection.DeleteOne(context.Background(), bson.M{"_id": objectId})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (cm *ClientMaster) ListClients() ([]model.ListClient, error) {
	var clientData []*model.Client
	var finaldata []model.ListClient
	filter := bson.D{{Key: "status", Value: bson.M{"$ne": "I"}}}
	cur, err := cm.ClientCollection.Find(context.Background(), filter)
	if err != nil {
		return finaldata, err
	}

	cur.All(context.Background(), &clientData)

	if len(clientData) == 0 {
		return finaldata, mongo.ErrNoDocuments
	}

	for i := range clientData {

		clientId := clientData[i].ID
		clientNamestr := clientData[i].ClientName
		clientCodestr := clientData[i].ClientCode

		client := model.ListClient{ID: clientId, ClientName: clientNamestr, ClientCode: clientCodestr}

		finaldata = append(finaldata, client)
	}

	return finaldata, nil
}

func TrimClientDetails(clientReq model.Client) model.Client {
	if clientReq.ClientName != "" {
		clientReq.ClientName = strings.TrimSpace(clientReq.ClientName)
	}
	if clientReq.ClientCode != "" {
		clientReq.ClientCode = strings.TrimSpace(clientReq.ClientCode)
	}

	return clientReq
}
