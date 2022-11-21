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

var cdfmd service.ClientDocFieldMappingDAO

func init() {
	const collectionName2 = "clientDocumentFieldMapping"
	cdfmd = service.ClientDocFieldMappingDAO{ClientDocFieldMappingCollection: config.GetCollection(config.DB, collectionName2)}
}

func InsertClientDocFieldMapping(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	var dataBody model.ClientDocumentFieldRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if dataBody.DocumentFieldName == "" || dataBody.ClientName == "" || dataBody.DocumentName == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide documentFieldName, clientName or documenName")
		return
	}

	if result, err := cdfmd.Insert(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record inserted successfully", true, result)
	}
}

func FindClientDocFieldMappingByID(w http.ResponseWriter, r *http.Request) {
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
	if result, err := cdfmd.FindByID(id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record fetch successfully", true, result)
	}
}

func FindAllClientDocFieldMapping(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := cdfmd.FindAll(); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record fetch successfully", true, result)
	}
}

func DeactivateClientDocFieldMappingByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide Id for Deactivate")
		return
	}

	if result, err := cdfmd.DeactivatedByID(id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record deactivated successfully", true, result)
	}
}

func SearchFilterClientDocFieldMapping(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	var dataBody model.ClientDocFieldSearchFilterRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := cdfmd.SearchFilter(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record fetch successfully", true, result)
	}
}

func UpdateClientDocFieldMappingByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	var dataBody model.ClientDocumentFieldRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := cdfmd.UpdateByID(id, dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record updated successfully", true, result)
	}
}
