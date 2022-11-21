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

var dfmService service.DocumentFieldMappingDAO

func init() {
	collectionName := "documentFieldMapping"
	dfmService = service.DocumentFieldMappingDAO{DocumentFieldMappingCollection: config.GetCollection(config.DB, collectionName)}
}

func AddNewDocField(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var dataBody model.DocumentFieldMapping
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	isPresent, err := dfmService.IsAlreadyPresent(dataBody)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isPresent {
		helper.RespondWithError(w, http.StatusBadRequest, "mapping already exists in the DB")
		return
	}

	result, err := dfmService.Insert(dataBody)

	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithJson(w, http.StatusOK, "Record inserted successfully", true, result)
}

func FindAllDocFieldMapping(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	result, err := dfmService.FindAll()

	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithJson(w, http.StatusOK, "Records fetched successfully", true, result)
}

func FindDocFieldMappingById(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	path := r.URL.Path
	pathElements := strings.Split(path, "/")

	id := pathElements[len(pathElements)-1]

	if len(pathElements) != 5 || id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "id is not provided in the url")
		return
	}

	result, err := dfmService.FindById(id)

	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithJson(w, http.StatusOK, "Records fetched successfully", true, result)
}

func DeleteDocFieldMappingById(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	path := r.URL.Path
	pathElements := strings.Split(path, "/")

	id := pathElements[len(pathElements)-1]

	if len(pathElements) != 5 || id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "id is not provided in the url")
		return
	}

	result, err := dfmService.DeleteById(id)

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

func UpdateDocFieldMappingById(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	path := r.URL.Path
	pathElements := strings.Split(path, "/")

	id := pathElements[len(pathElements)-1]

	if len(pathElements) != 5 || id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "id is not provided in the url")
		return
	}

	var dataBody model.DocumentFieldMapping
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	result, err := dfmService.UpdateById(id, dataBody)

	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithJson(w, http.StatusOK, "Record Updated successfully", true, result)
}

func SearchDocFieldMappingByFilter(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var dataBody model.DocumentFieldMapping

	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	result, err := dfmService.SearchFilterOnDocument(dataBody)

	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithJson(w, http.StatusOK, "Record inserted successfully", true, result)
}
