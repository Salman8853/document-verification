package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInsertDocFieldMaster(t *testing.T) {
	payload := strings.NewReader(`{
		"documentFieldName": "test_dummy"
	}`)
	payloadWithoutFieldName := strings.NewReader(`{
		"documentFieldName": ""
	}`)
	payloadWithEmptyJson := strings.NewReader(`{
		}`)
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/add/document_field_master/", nil)
		response := httptest.NewRecorder()

		InsertDocFieldMaster(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if new doc field name is added", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/add/document_field_master/", payload)
		response := httptest.NewRecorder()

		InsertDocFieldMaster(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusAccepted)
	})
	t.Run("check with empty json", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/add/document_field_master/", payloadWithEmptyJson)
		response := httptest.NewRecorder()

		InsertDocFieldMaster(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if new doc field name is added when docuemntFieldName is empty", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/add/document_field_master/", payloadWithoutFieldName)
		response := httptest.NewRecorder()

		InsertDocFieldMaster(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
}

func TestFindDocFieldMasterByID(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/search/document_field_master/", nil)
		response := httptest.NewRecorder()

		FindDocFieldMasterByID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if api able to fetch data", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/search/document_field_master/6374afec52056ae109537f87", nil)
		response := httptest.NewRecorder()

		FindDocFieldMasterByID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if api able to fetch data", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/search/document_field_master/", nil)
		response := httptest.NewRecorder()

		FindDocFieldMasterByID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
}

func TestFindAllDocFieldMaster(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/search-all/document_field_master/", nil)
		response := httptest.NewRecorder()

		FindAllDocFieldMaster(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if api is fetching data", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/search-all/document_field_master/", nil)
		response := httptest.NewRecorder()

		FindAllDocFieldMaster(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusAccepted)
	})
}

func TestDeactivateDocFieldMasterByID(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/delete/document_field_master/", nil)
		response := httptest.NewRecorder()

		DeactivateDocFieldMasterByID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check when id is not passed", func(t *testing.T) {
		request := httptest.NewRequest("DELETE", "/api/delete/document_field_master/", nil)
		response := httptest.NewRecorder()

		DeactivateDocFieldMasterByID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check when id is passed", func(t *testing.T) {
		request := httptest.NewRequest("DELETE", "/api/delete/document_field_master/6374afec52056ae109537f18", nil)
		response := httptest.NewRecorder()

		DeactivateDocFieldMasterByID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
}

func TestSearchFilterDocFieldMaster(t *testing.T) {
	payload := strings.NewReader(`{
		"documentFieldName": "test_dummy"
	}`)
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/search-filter/document_field_master/", nil)
		response := httptest.NewRecorder()

		SearchFilterDocFieldMaster(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if data is being fetched", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/search-filter/document_field_master/", payload)
		response := httptest.NewRecorder()

		SearchFilterDocFieldMaster(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusAccepted)
	})
}

func TestUpdateDocFieldMasterByID(t *testing.T) {
	payload := strings.NewReader(`{
		"documentFieldName": "test96"
	}`)
	payloadWithEmptyJson := strings.NewReader(`{
		"documentFieldName": ""
	}`)
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/update/document_field_master/", nil)
		response := httptest.NewRecorder()

		UpdateDocFieldMasterByID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check response when id is not passed", func(t *testing.T) {
		request := httptest.NewRequest("PUT", "/api/update/document_field_master/", payload)
		response := httptest.NewRecorder()

		UpdateDocFieldMasterByID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check response when id and request body passed", func(t *testing.T) {
		request := httptest.NewRequest("PUT", "/api/update/document_field_master/6374afec52056ae109537f17", payload)
		response := httptest.NewRecorder()

		UpdateDocFieldMasterByID(response, request)
		log.Println(payload)
		log.Println(response.Body)
		checkStatusCodeAndRespond(t, response.Code, http.StatusAccepted)
	})
	t.Run("check response when id and empty request body passed ", func(t *testing.T) {
		request := httptest.NewRequest("PUT", "/api/update/document_field_master/6374afec52056ae109537f17", payloadWithEmptyJson)
		response := httptest.NewRecorder()

		UpdateDocFieldMasterByID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusAccepted)
	})
}

func TestFindAllDocumentFieldName(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("DELETE", "/api/list_all/document_field_master/", nil)
		response := httptest.NewRecorder()

		FindAllDocumentFieldName(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check data is being fetched", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/list_all/document_field_master/", nil)
		response := httptest.NewRecorder()

		FindAllDocumentFieldName(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusAccepted)
	})
}

func checkStatusCodeAndRespond(t *testing.T, got, expected int) {
	t.Helper()

	if got != expected {
		t.Errorf("Failed. Expected %d Got %d\n", expected, got)
	}
}
