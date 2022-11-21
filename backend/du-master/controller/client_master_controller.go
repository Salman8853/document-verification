package controller

import (
	"du-master/config"
	"du-master/helper"
	"du-master/model"
	"du-master/service"
	"encoding/json"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

var con service.ClientMaster

const collectionClient = "clientMaster"

func init() {
	con = service.ClientMaster{ClientCollection: config.GetCollection(config.DB, collectionClient)}
}

func AddClientDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody model.Client
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	if reqBody.ClientName == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide ClientName")
		return
	}
	if reqBody.ClientCode == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide ClientCode")
		return
	}

	if result, err := con.Insert(reqBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record inserted successfully", true, result)
	}
}

func FindAllClient(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	result, err := con.FindAll()

	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithJson(w, http.StatusOK, "Records fetched successfully", true, result)
}

func FindByClientId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	if result, err := con.FindByClientId(id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record Found successfully", true, result)
	}
}

func FindByClientName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	name := segment[len(segment)-1]
	if name == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide Name for Search")
	}

	if result, err := con.FindByClientName(name); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record Found successfully", true, result)
	}
}

func SearchMultipleClient(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody model.Client
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if result, err := con.SearchFilterOnClient(reqBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "client fetch successfully", true, result)
	}
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody model.Client
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if result, err := con.UpdateByClientId(dataBody, id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Client Updated successfully", true, result)
	}
}

func DeleteClientById(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	result, err := con.DeleteById(id)

	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, _ := result.((*mongo.DeleteResult))

	if res.DeletedCount == 0 {
		helper.RespondWithError(w, http.StatusBadRequest, "unable to delete the record")
		return
	}

	helper.RespondWithJson(w, http.StatusOK, "Records Deleted successfully", true, result)
}

func DeactivateClientById(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide ClientName or ClientCode")
	}

	if result, err := con.DeactivateClient(id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Client Deactivated successfully", true, result)
	}
}
func ListOfAllClient(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user, err := con.ListClients()

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, user)
	}
}
