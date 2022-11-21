package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddClientDocMap(t *testing.T) {
	payload := strings.NewReader(`{
		"clientName": "Amazon",
		"documentName": "aadharcard"
	}`)
	payloadWithoutFieldName := strings.NewReader(`{
		"clientName": ""
	}`)
	payloadWithEmptyJson := strings.NewReader(`{
		}`)
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/add-client-doc-map", nil)
		response := httptest.NewRecorder()

		AddClientDocMap(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if new doc field name is added", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/add-client-doc-map", payload)
		response := httptest.NewRecorder()

		AddClientDocMap(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusAccepted)
	})
	t.Run("check with empty json", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/add-client-doc-map", payloadWithEmptyJson)
		response := httptest.NewRecorder()

		AddClientDocMap(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if new doc field name is added when docuemntFieldName is empty", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/add-client-doc-map", payloadWithoutFieldName)
		response := httptest.NewRecorder()

		AddClientDocMap(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
}

func TestSearchByCliDocMapID(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("", "/api/search-cliDocMap/", nil)
		response := httptest.NewRecorder()

		SearchByCliDocMapID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if api able to fetch data", func(t *testing.T) {
		request := httptest.NewRequest("", "/api/search-cliDocMap/636e0cd1d355f89ca4c7165a", nil)
		response := httptest.NewRecorder()

		SearchByCliDocMapID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if api able to fetch data", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/search-cliDocMap/", nil)
		response := httptest.NewRecorder()

		SearchByCliDocMapID(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
}

func TestSearchMultipleCDM(t *testing.T) {
	payload := strings.NewReader(`{
		"documentFieldName": "test_dummy"
	}`)
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/search-multiple-cdm", nil)
		response := httptest.NewRecorder()

		SearchMultipleCDM(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if data is being fetched", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/search-multiple-cdm", payload)
		response := httptest.NewRecorder()

		SearchMultipleCDM(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusAccepted)
	})
}

func TestUpdateCDM(t *testing.T) {
	payload := strings.NewReader(`{
		"status": "I"
	}`)
	payloadWithEmptyJson := strings.NewReader(`{
		"status": ""
	}`)
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/update-cdm/", nil)
		response := httptest.NewRecorder()

		UpdateCDM(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check response when id is not passed", func(t *testing.T) {
		request := httptest.NewRequest("PUT", "/api/update-cdm/", payload)
		response := httptest.NewRecorder()

		UpdateCDM(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check response when id and request body passed", func(t *testing.T) {
		request := httptest.NewRequest("PUT", "/api/update-cdm/6375d26ead48211eb67e6a00", payload)
		response := httptest.NewRecorder()

		UpdateCDM(response, request)
		log.Println(payload)
		log.Println(response.Body)
		checkStatusCodeAndRespond(t, response.Code, http.StatusAccepted)
	})
	t.Run("check response when id and empty request body passed ", func(t *testing.T) {
		request := httptest.NewRequest("PUT", "/api/update-cdm/6375d26ead48211eb67e6a00", payloadWithEmptyJson)
		response := httptest.NewRecorder()

		UpdateCDM(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
}

func TestDeactivateCDM(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/delete-cdm/", nil)
		response := httptest.NewRecorder()

		DeactivateCDM(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check when id is not passed", func(t *testing.T) {
		request := httptest.NewRequest("DELETE", "/api/delete-cdm/", nil)
		response := httptest.NewRecorder()

		DeactivateCDM(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check when id is passed", func(t *testing.T) {
		request := httptest.NewRequest("DELETE", "/api/delete-cdm/6375d26ead48211eb67e6a00", nil)
		response := httptest.NewRecorder()

		DeactivateCDM(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
}

func TestGetAllCDM(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/get-all-cdm", nil)
		response := httptest.NewRecorder()

		GetAllCDM(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusBadRequest)
	})
	t.Run("check if api is fetching data", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/get-all-cdm", nil)
		response := httptest.NewRecorder()

		GetAllCDM(response, request)
		checkStatusCodeAndRespond(t, response.Code, http.StatusAccepted)
	})
}

func checkStatusCodeAndRespond(t *testing.T, got, expected int) {
	t.Helper()

	if got != expected {
		t.Errorf("Failed. Expected %d Got %d\n", expected, got)
	}
}
