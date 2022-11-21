package controller

import (
	"du-master/config"
	"du-master/helper"
	"du-master/model"
	"du-master/service"
	"encoding/json"
	"net/http"
	"strings"
)

var CDMDao service.ClientDocumentMappingDAO

func init() {
	collectionName := "clientDocumentMapping"
	CDMDao = service.ClientDocumentMappingDAO{ClientDocumentMappingCollection: config.GetCollection(config.DB, collectionName)}
}

func AddClientDocMap(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	var dataBody model.CDMRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if dataBody.ClientName == "" || dataBody.DocumentName == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide clientName and documenName")
		return
	}

	if result, err := CDMDao.InsertCliDocMapping(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record inserted successfully", true, result)
	}
}

func SearchByCliDocMapID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	user, err := CDMDao.CliDocMapFindById(id)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, user)
	}
}

func SearchMultipleCDM(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	var dataBody model.CDMRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	Document, err := CDMDao.SearchFilterOnDocument(dataBody)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, Document)
	}
}

func DeactivateCDM(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if result, err := CDMDao.DeactivateCDM(id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Document Status changed successfully", true, result)
	}
}

func UpdateCDM(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody model.CDMRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if result, err := CDMDao.Update(dataBody, id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Document Updated successfully", true, result)
	}
}

func GetAllCDM(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	user, err := CDMDao.FindAll()

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, user)
	}
}
