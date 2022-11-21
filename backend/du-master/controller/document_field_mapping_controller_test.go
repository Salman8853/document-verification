package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddNewDocField(t *testing.T) {

	payload := strings.NewReader(`
		{
			"documentName": "Pan card",
			"documentFieldName": "Full Name",
			"status": "A"
		}
	`)

	t.Run("check if http methods other than POST are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/add/document_field_mapping", nil)
		response := httptest.NewRecorder()

		AddNewDocField(response, request)
		got := response.Body.String()
		want := `{"success":false,"success_msg":"method not allowed","data":{}}`

		assertBody(t, got, want)
		assertStatusCode(t, http.StatusMethodNotAllowed, response.Code)
	})

	t.Run("check if new doc field mapping is added", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/add/document_field_mapping", payload)
		response := httptest.NewRecorder()

		AddNewDocField(response, request)

		assertStatusCode(t, http.StatusOK, response.Code)

		assertBodyContent(t, response.Body)
	})
}

func TestFindAllDocFieldMapping(t *testing.T) {
	t.Run("check if http methods other than GET are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/search_all/document_field_mapping", nil)
		response := httptest.NewRecorder()

		FindAllDocFieldMapping(response, request)
		got := response.Body.String()
		want := `{"success":false,"success_msg":"method not allowed","data":{}}`

		assertBody(t, got, want)
		assertStatusCode(t, http.StatusMethodNotAllowed, response.Code)
	})

	t.Run("check if all doc field mapping are returned", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/search_all/document_field_mapping", nil)
		response := httptest.NewRecorder()

		FindAllDocFieldMapping(response, request)

		assertStatusCode(t, http.StatusOK, response.Code)

		assertBodyContent(t, response.Body)
	})
}

func assertStatusCode(t *testing.T, got_status int, want_status int) {
	t.Helper()

	if got_status != want_status {
		t.Errorf("got %d want %d", got_status, want_status)
	}
}

func assertBody(t *testing.T, body_got string, body_want string) {
	t.Helper()

	if body_got != body_want {
		t.Errorf("got %s want %s", body_got, body_want)
	}
}

func assertBodyContent(t *testing.T, body *bytes.Buffer) {
	t.Helper()

	var bodyMap map[string]interface{}
	_ = json.Unmarshal(body.Bytes(), &bodyMap)

	if success, _ := bodyMap["success"].(bool); !success {
		t.Errorf("got success: %t but want success: %s", success, "true")
	}

	if bodyMap["data"] == nil {
		t.Error("data should be found in response")
	}
}
