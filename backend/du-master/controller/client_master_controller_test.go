package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddClient(t *testing.T) {
	payload := strings.NewReader(`
	{
		"clientName":"RPA Testingw12",
		"clientCode":"dwed12",
		"address":"delhi",
		"logoPath":"",
		"functionalEntity":"IT",
		"status":"Active"
	}
	`)

	req := httptest.NewRequest("POST", "/api/add/client_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddClientDetails)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFindByClientId(t *testing.T) {

	req := httptest.NewRequest("GET", "/api/find_client_by_id/client_master/636d3ede4ac9d8e4849bd94b", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(FindByClientId)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFindByClientName(t *testing.T) {

	req := httptest.NewRequest("GET", "/api/find_client_by_name/client_master/{name}", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(FindByClientName)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFindAllClient(t *testing.T) {

	req := httptest.NewRequest("GET", "/api/find_all/client_master", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(FindAllClient)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestSearchMultipleClient(t *testing.T) {

	payload := strings.NewReader(`
	{
		"clientName":"RPA Testing",
		"clientCode":"assw456",
		"address":"delhi",
		"logoPath":"",
		"functionalEntity":"IT",
		"status":"Active"
	}
	`)
	req := httptest.NewRequest("POST", "/api/search_multiple_filter/client_master/", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(SearchMultipleClient)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateClient(t *testing.T) {

	payload := strings.NewReader(`
	{
		"clientName":"RPA Testing",
		"clientCode":"annu456",
		"address":" New delhi",
		"logoPath":"",
		"functionalEntity":"IT",
		"status":"Active"
	}
	`)
	req := httptest.NewRequest("PUT", "/api/update_client_by_id/client_master/636d3ede4ac9d8e4849bd94b", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(UpdateClient)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeactivateById(t *testing.T) {

	req := httptest.NewRequest("PUT", "/api/deactivate_by_id/client_master/636d3ede4ac9d8e4849bd94b", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(DeleteClientById)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteById(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/api/delete-client_by_id/636d3ede4ac9d8e4849bd94b", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(DeleteClientById)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestListOfAllClient(t *testing.T) {

	req := httptest.NewRequest("GET", "/api/list_of_client/client_master/", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(ListOfAllClient)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// --------------------Negative Testing---------------------------------

func TestAddDublicateClientName(t *testing.T) {
	payload := strings.NewReader(`
	{
		"clientName":"RPA Testing",
		"clientCode":"assw456",
		"address":"delhi",
		"logoPath":"",
		"functionalEntity":"IT"
	}
	`)

	req := httptest.NewRequest("POST", "/api/add/client_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddClientDetails)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestAddDublicateClientCode(t *testing.T) {
	payload := strings.NewReader(`
	{
		"clientName":"RPA Testings",
		"clientCode":"assw456",
		"address":"delhi",
		"logoPath":"",
		"functionalEntity":"IT"
	}
	`)

	req := httptest.NewRequest("POST", "/api/add/client_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddClientDetails)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestAddWithoutClientName(t *testing.T) {
	payload := strings.NewReader(`
	{
		"clientName":"",
		"clientCode":"assw456",
		"address":"delhi",
		"logoPath":"",
		"functionalEntity":"IT"
	}
	`)

	req := httptest.NewRequest("POST", "/api/add/client_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddClientDetails)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestAddWithoutClientCode(t *testing.T) {
	payload := strings.NewReader(`
	{
		"clientName":"RPA Testings",
		"clientCode":"",
		"address":"delhi",
		"logoPath":"",
		"functionalEntity":"IT"
	}
	`)

	req := httptest.NewRequest("POST", "/api/add/client_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddClientDetails)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFindByClientwithoutId(t *testing.T) {

	req := httptest.NewRequest("GET", "/api/find_client_by_id/client_master/", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(FindByClientId)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFindByClientInvalid(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/api/find_client_by_id/client_master/52364", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(FindByClientId)
	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestFindByClientWithName(t *testing.T) {

	req := httptest.NewRequest("GET", "/api/find_client_by_name/client_master/", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(FindByClientName)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateClientWithOutId(t *testing.T) {

	payload := strings.NewReader(`
	{
		"clientName":"RPA Testing",
		"clientCode":"annu456",
		"address":" New delhi",
		"logoPath":"",
		"functionalEntity":"IT",
		"status":"Active"
	}
	`)
	req := httptest.NewRequest("PUT", "/api/update_client_by_id/client_master/", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(UpdateClient)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateClientInvalidId(t *testing.T) {

	payload := strings.NewReader(`
	{
		"clientName":"RPA Testing",
		"clientCode":"annu456",
		"address":" New delhi",
		"logoPath":"",
		"functionalEntity":"IT",
		"status":"Active"
	}
	`)
	req := httptest.NewRequest("PUT", "/api/update_client_by_id/client_master/253645", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(UpdateClient)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteByWithoutId(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/api/delete-client_by_id/", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(DeleteClientById)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteInvalidId(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/api/delete-client_by_id/142578", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(DeleteClientById)

	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
