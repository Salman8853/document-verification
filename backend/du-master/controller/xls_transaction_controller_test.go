package controller

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"
)

func TestXlsTransactionCreateInDb(t *testing.T) {
	filePath := "D:\\go lang workspace\\document-verification\\backend\\du-master\\excel\\download\\uploadFromExcelTemplate.xlsx"
	fieldName := "file"
	body := new(bytes.Buffer)

	mw := multipart.NewWriter(body)

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	w, err := mw.CreateFormFile(fieldName, filePath)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := io.Copy(w, file); err != nil {
		t.Fatal(err)
	}

	// close the writer before making the request
	mw.Close()

	// router is of type http.Handler
	request := httptest.NewRequest("POST", "/api/xls_transaction_case_creation/", body)
	request.Header.Add("Content-Type", mw.FormDataContentType())

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(XlsTransactionCaseCreation)
	handler.ServeHTTP(response, request)

	fmt.Println(response.Code)
	fmt.Println(response.Body)
	// Check the status code is what we expect.
	if status := response.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// ------------------------negative Testing-------------------
func TestXlsTransactionSaveWithOutMandatoryField(t *testing.T) {
	filePath := "D:\\go lang workspace\\document-verification\\backend\\du-master\\excel\\download\\uploadFromExcelTemplate.xlsx"
	fieldName := "file"
	body := new(bytes.Buffer)

	mw := multipart.NewWriter(body)

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	w, err := mw.CreateFormFile(fieldName, filePath)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := io.Copy(w, file); err != nil {
		t.Fatal(err)
	}

	// close the writer before making the request
	mw.Close()

	// router is of type http.Handler
	request := httptest.NewRequest("POST", "/api/xls_transaction_case_creation/", body)
	request.Header.Add("Content-Type", mw.FormDataContentType())

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(XlsTransactionCaseCreation)
	handler.ServeHTTP(response, request)

	fmt.Println(response.Code)
	fmt.Println(response.Body)
	// Check the status code is what we expect.
	if status := response.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestXlsTransactionInvalidRequest(t *testing.T) {
	filePath := "D:\\go lang workspace\\document-verification\\backend\\du-master\\excel\\download\\uploadFromExcelTemplate.xlsx"
	fieldName := "file"
	body := new(bytes.Buffer)

	mw := multipart.NewWriter(body)

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	w, err := mw.CreateFormFile(fieldName, filePath)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := io.Copy(w, file); err != nil {
		t.Fatal(err)
	}

	// close the writer before making the request
	mw.Close()

	// router is of type http.Handler
	request := httptest.NewRequest("POST", "/api/xls_transaction_case_creation/", nil)
	request.Header.Add("Content-Type", mw.FormDataContentType())

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(XlsTransactionCaseCreation)
	handler.ServeHTTP(response, request)

	fmt.Println(response.Code)
	fmt.Println(response.Body)
	// Check the status code is what we expect.
	if status := response.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
