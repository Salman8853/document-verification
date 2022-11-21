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

var dfm service.DocumentFieldMasterDAO

func init() {
	const collectionName = "documentFieldMaster"
	dfm = service.DocumentFieldMasterDAO{DocumentFieldCollection: config.GetCollection(config.DB, collectionName)}
}

func InsertDocFieldMaster(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	var dataBody model.DocumentFieldMasterRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	if dataBody.DocumentFieldName == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide documentFieldName")
		return
	}
	if result, err := dfm.Insert(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record inserted successfully", true, result)
	}
}

func FindDocFieldMasterByID(w http.ResponseWriter, r *http.Request) {
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
	if result, err := dfm.FindByID(id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record fetch successfully", true, result)
	}
}

func FindAllDocFieldMaster(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := dfm.FindAll(); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record fetch successfully", true, result)
	}
}

func DeactivateDocFieldMasterByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide Id for Deactivate")
	}

	if result, err := dfm.DeactivatedByID(id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record deactivated successfully", true, result)
	}
}

func SearchFilterDocFieldMaster(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	var dataBody model.DocumentFieldMasterSearchRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := dfm.SearchFilter(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record fetch successfully", true, result)
	}
}

func UpdateDocFieldMasterByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
		return
	}

	var dataBody model.DocumentFieldMasterSearchRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := dfm.UpdateByID(id, dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record updated successfully", true, result)
	}
}

func FindAllDocumentFieldName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := dfm.FindAllDocFieldName(); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record fetch successfully", true, result)
	}
}
