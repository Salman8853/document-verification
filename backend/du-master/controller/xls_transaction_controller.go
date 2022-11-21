package controller

import (
	"du-master/config"
	"du-master/helper"
	"du-master/service"
	"net/http"
)

var xlsCon service.DocumentTransactionService

const collectionXls = "documentTransactionDetails"

func init() {
	xlsCon = service.DocumentTransactionService{DocumentTransactionCollection: config.GetCollection(config.DB, collectionXls)}
}

func XlsTransactionCaseCreation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 32 MB is the default used by FormFile()
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get a reference to the fileHeaders.
	// They are accessible only after ParseMultipartForm is called
	files := r.MultipartForm.File["file"]
	if len(files) != 1 {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide only one excel file")
		return
	}

	if result, err := xlsCon.XlsTransactionCaseCreation(files); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "excel transaction added successfully", true, result)
	}
}
