package controller

import (
	"du-master/config"
	"du-master/helper"
	"du-master/model"
	"du-master/service"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

var ead = service.DocumentMaster{}

// func init() {
// 	ead.Server = "mongodb://localhost:27017/"
// 	ead.Database = "DocumentVerification"
// 	ead.Collection = "DocumentMaster"

// 	ead.DocumentConnect()
// }

// var dfmService service.DocumentFieldMappingDAO

func init() {
	collectionName := "documentMaster"
	ead = service.DocumentMaster{DocumentMasterCollection: config.GetCollection(config.DB, collectionName)}
}

func AddDocument(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dataBody model.Document

	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := checkRequired(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if result, err := ead.DocumentInsert(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record inserted successfully", true, result)
	}
}

func checkRequired(DocumentRequest model.Document) error {
	if DocumentRequest.DocumentName == "" {
		return errors.New("please provide a document name")
	}
	if DocumentRequest.DocumentType == "" {
		return errors.New("please provide a document type")
	}
	if DocumentRequest.DocumentTrainingStatus == "" {
		return errors.New("please provide a document training status")
	}
	return nil
}

func SearchOneDocumentById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		http.Error(w, "Please provide Id for Search", http.StatusBadRequest)
		return
	}

	user, err := ead.DocumentFindById(id)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, user)
	}
}

func SearchOneDocumentByName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	name := segment[len(segment)-1]
	if name == "" {
		http.Error(w, "Please provide Document Name for Search", http.StatusBadRequest)
		return
	}

	user, err := ead.DocumentFindByName(name)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, user)
	}
}

func SearchMultipleDocument(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dataBody model.Document
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	Document, err := ead.SearchFilterOnDocument(dataBody)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, Document)
	}
}

func UpdateDocument(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]
	if id == "" {
		http.Error(w, "Please provide Id for Search", http.StatusBadRequest)
		return
	}

	var dataBody model.Document
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := checkRequired(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if result, err := ead.Update(dataBody, id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Document Updated successfully", true, result)
	}
}

func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]
	if id == "" {
		http.Error(w, "Please provide Id for Search", http.StatusBadRequest)
		return
	}

	if result, err := ead.DeleteDocument(id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, result, true, nil)
	}
}

func ListOfAllDocuments(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user, err := ead.ListDocuments()

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, user)
	}
}
