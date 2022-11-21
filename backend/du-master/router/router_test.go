package router

import (
	"du-master/controller"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddDocument(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("PUT", "/api/add/document_master", nil)
		response := httptest.NewRecorder()

		controller.AddDocument(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}

func TestSearchOneDocumentById(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/search_one_by_id/document_master/", nil)
		response := httptest.NewRecorder()

		controller.SearchOneDocumentById(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}

func TestSearchOneDocumentByName(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/search_one_by_name/document_master/", nil)
		response := httptest.NewRecorder()

		controller.SearchOneDocumentByName(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}

func TestSearchMultipleDocument(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/search_by_filter/document_master", nil)
		response := httptest.NewRecorder()

		controller.SearchMultipleDocument(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}

func TestUpdateDocument(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/update/document_master/", nil)
		response := httptest.NewRecorder()

		controller.UpdateDocument(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}
	})
}

func TestDeleteDocument(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/delete/document_master/", nil)
		response := httptest.NewRecorder()

		controller.DeleteDocument(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}
	})
}

func TestListOfAllDocuments(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/list_all/document_master", nil)
		response := httptest.NewRecorder()

		controller.ListOfAllDocuments(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}
	})
}

// ---------------------client Master---------------------------

func TestAddClientDetails(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/add/client_master", nil)
		response := httptest.NewRecorder()

		controller.AddClientDetails(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}

func TestFindAllClient(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/find_all/client_master", nil)
		response := httptest.NewRecorder()

		controller.AddClientDetails(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}

func TestFindByClientId(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/find_client_by_id/client_master/", nil)
		response := httptest.NewRecorder()

		controller.FindByClientId(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}

func TestFindByClientName(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/find_client_by_name/client_master/", nil)
		response := httptest.NewRecorder()

		controller.FindByClientName(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}

func TestSearchMultipleClient(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/search_multiple_filter/client_master/", nil)
		response := httptest.NewRecorder()

		controller.SearchMultipleClient(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}
func TestUpdateClient(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/update_client_by_id/client_master/", nil)
		response := httptest.NewRecorder()

		controller.UpdateClient(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}
	})
}
func TestDeleteClientById(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("Get", "/api/delete-client_by_id/", nil)
		response := httptest.NewRecorder()

		controller.DeleteClientById(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}

func TestDeactivateClientById(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("Get", "/api/deactivate_by_id/client_master/", nil)
		response := httptest.NewRecorder()

		controller.DeactivateClientById(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}

func TestListOfAllClient(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/list_of_client/client_master/", nil)
		response := httptest.NewRecorder()

		controller.ListOfAllClient(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}

// ------------------------xls transaction api---------------------

func TestXlsTransactionCaseCreation(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/xls_transaction_case_creation/", nil)
		response := httptest.NewRecorder()

		controller.XlsTransactionCaseCreation(response, request)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

	})
}
