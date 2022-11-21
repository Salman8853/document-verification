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

func TestCreateTransaction(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/create-transaction/", nil)
		response := httptest.NewRecorder()

		CreateTransaction(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if tranx created when passing req", func(t *testing.T) {
		filePath := "C:/Users/RamashankarKumar/Downloads/sample1.pdf"
		fieldName := "FileName"
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
		mw.WriteField("FirstName", "dummy")
		mw.WriteField("MiddleName", "")
		mw.WriteField("LastName", "kumar")
		mw.WriteField("Dob", "10-10-2022")
		mw.WriteField("Email", "dummy@gmail.com")
		mw.WriteField("ContactNumber", "12345")
		// close the writer before making the request
		mw.Close()
		request := httptest.NewRequest("POST", "/api/create-transaction/", body)
		request.Header.Add("Content-Type", mw.FormDataContentType())
		response := httptest.NewRecorder()
		CreateTransaction(response, request)
		fmt.Println(response.Body)
		checkStatusCodeAndRespond(t, response.Code, http.StatusAccepted)
	})
	t.Run("check if tranx created when some fields value are missing", func(t *testing.T) {
		filePath := "C:/Users/RamashankarKumar/Downloads/sample1.pdf"
		fieldName := "FileName"
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
		mw.WriteField("FirstName", "")
		mw.WriteField("MiddleName", "")
		mw.WriteField("LastName", "kumar")
		mw.WriteField("Dob", "10-10-2022")
		mw.WriteField("Email", "rk@gmail.com")
		mw.WriteField("ContactNumber", "12345")
		// close the writer before making the request
		mw.Close()
		request := httptest.NewRequest("POST", "/api/create-transaction/", body)
		request.Header.Add("Content-Type", mw.FormDataContentType())
		response := httptest.NewRecorder()
		CreateTransaction(response, request)
		fmt.Println(response.Body)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if tranx created when file is missing", func(t *testing.T) {
		body := new(bytes.Buffer)
		mw := multipart.NewWriter(body)
		mw.CreateFormFile("FileName", "")

		mw.WriteField("FirstName", "dummy")
		mw.WriteField("MiddleName", "")
		mw.WriteField("LastName", "kumar")
		mw.WriteField("Dob", "10-10-2022")
		mw.WriteField("Email", "rk@gmail.com")
		mw.WriteField("ContactNumber", "12345")
		// close the writer before making the request
		mw.Close()
		request := httptest.NewRequest("POST", "/api/create-transaction/", body)
		request.Header.Add("Content-Type", mw.FormDataContentType())
		response := httptest.NewRecorder()
		CreateTransaction(response, request)
		fmt.Println(response.Body)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
}
