package controller

import (
	"net/http/httptest"
	"testing"
)

func TestInsertClientDocFieldMapping(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/add/document_field_master/", nil)
		response := httptest.NewRecorder()

		InsertClientDocFieldMapping(response, request)
		got := response.Body.String()
		want := `{"success":false,"success_msg":"Invalid Request","data":{}}`

		if got == want {
			t.Logf("Pass. Expected %s, Got %s\n", want, got)
		} else {
			t.Errorf("Failed. Expected %s ,Got %s\n", want, got)
		}

	})
}

func TestFindClientDocFieldMappingByID(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/search/document_field_master/", nil)
		response := httptest.NewRecorder()

		FindClientDocFieldMappingByID(response, request)
		got := response.Body.String()
		want := `{"success":false,"success_msg":"Invalid Request","data":{}}`

		if got == want {
			t.Logf("Pass. Expected %s, Got %s\n", want, got)
		} else {
			t.Errorf("Failed. Expected %s ,Got %s\n", want, got)
		}

	})
}

func TestFindAllClientDocFieldMapping(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/search-all/document_field_master/", nil)
		response := httptest.NewRecorder()

		FindAllClientDocFieldMapping(response, request)
		got := response.Body.String()
		want := `{"success":false,"success_msg":"Invalid Request","data":{}}`

		if got == want {
			t.Logf("Pass. Expected %s, Got %s\n", want, got)
		} else {
			t.Errorf("Failed. Expected %s ,Got %s\n", want, got)
		}

	})
}

func TestDeactivateClientDocFieldMappingByID(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/delete/document_field_master/", nil)
		response := httptest.NewRecorder()

		DeactivateClientDocFieldMappingByID(response, request)
		got := response.Body.String()
		want := `{"success":false,"success_msg":"Invalid Request","data":{}}`

		if got == want {
			t.Logf("Pass. Expected %s, Got %s\n", want, got)
		} else {
			t.Errorf("Failed. Expected %s ,Got %s\n", want, got)
		}

	})
}

func TestSearchFilterClientDocFieldMapping(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/search-filter/document_field_master/", nil)
		response := httptest.NewRecorder()

		SearchFilterClientDocFieldMapping(response, request)
		got := response.Body.String()
		want := `{"success":false,"success_msg":"Invalid Request","data":{}}`

		if got == want {
			t.Logf("Pass. Expected %s, Got %s\n", want, got)
		} else {
			t.Errorf("Failed. Expected %s ,Got %s\n", want, got)
		}

	})
}

func TestUpdateClientDocFieldMappingByID(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/update/document_field_master/", nil)
		response := httptest.NewRecorder()

		UpdateClientDocFieldMappingByID(response, request)
		got := response.Body.String()
		want := `{"success":false,"success_msg":"Invalid Request","data":{}}`

		if got == want {
			t.Logf("Pass. Expected %s, Got %s\n", want, got)
		} else {
			t.Errorf("Failed. Expected %s ,Got %s\n", want, got)
		}

	})
}
