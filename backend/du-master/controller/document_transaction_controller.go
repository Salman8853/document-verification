package controller

import (
	"du-master/config"
	"du-master/helper"
	"du-master/service"
	"net/http"
)

var dtc service.DocumentTransactionService

const MaxUploadSize = 20 * 1024 * 1024 // 20 mb

func init() {
	const dtcCollection = "documentTransactionDetails"
	dtc = service.DocumentTransactionService{DocumentTransactionCollection: config.GetCollection(config.DB, dtcCollection)}
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	files := r.MultipartForm.File["FileName"]
	firstName := r.MultipartForm.Value["FirstName"][0]
	middleName := r.MultipartForm.Value["MiddleName"][0]
	lastName := r.MultipartForm.Value["LastName"][0]
	email := r.MultipartForm.Value["Email"][0]
	dob := r.MultipartForm.Value["Dob"][0]
	contactNumber := r.MultipartForm.Value["ContactNumber"][0]

	if len(files) == 0 {
		helper.RespondWithError(w, http.StatusBadRequest, "please upload atleast one file")
		return
	}
	if firstName == "" || lastName == "" || email == "" || dob == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "please provide firstName, lastName, email or dob as these are mandatory")
		return
	}

	if result, err := dtc.InsertDT(files, firstName, middleName, lastName, email, dob, contactNumber); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "transaction created successfully", true, result)
	}
}
