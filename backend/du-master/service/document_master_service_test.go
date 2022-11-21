package service

import (
	"du-master/model"
	"testing"
)

func TestDocumentInsert(t *testing.T) {
	var document model.Document
	document.DocumentName = " Name of document "
	document.DocumentType = " Type of document "
	document.DocumentTrainingStatus = " Du model training status of the document "

	document = TrimDocumentDetails(document)

	// Check the status code is what we expect.
	if !(document.DocumentName == "Name of document") {
		t.Errorf("handler returned wrong status code: got \"%v\" want \"Name of document\"", document.DocumentName)
	}

	if !(document.DocumentType == "Type of document") {
		t.Errorf("handler returned wrong status code: got \"%v\" want \"Type of document\"", document.DocumentType)
	}

	if !(document.DocumentTrainingStatus == "Du model training status of the document") {
		t.Errorf("handler returned wrong status code: got \"%v\" want \"Du model training status of the document\"", document.DocumentTrainingStatus)
	}

}
