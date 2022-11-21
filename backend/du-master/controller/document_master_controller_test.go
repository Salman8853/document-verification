package controller

import (
	"du-master/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDocumentInsert(t *testing.T) {
	payload := strings.NewReader(`{
		"documentName" : "Name of document",
		"documentType" : "Type of document",
		"status" : "A",
		"documentTrainingStatus" : "Du model training status of the document"
		}`)

	req := httptest.NewRequest("POST", "/api/add/document_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDocumentFindById(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/search_one_by_id/document_master/636df00176f1777049e1a282", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(SearchOneDocumentById)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDocumentFindByName(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/search_one_by_name/document_master/Name%20of%20document", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(SearchOneDocumentByName)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestSearchMultipleUser(t *testing.T) {

	payload := strings.NewReader(`{
		"documentName" : "Name of document",
		"documentType" : "Type of document",
		"status" : "A",
		"documentTrainingStatus" : "Du model training status of the document"
		}`)

	req := httptest.NewRequest("POST", "/api/search_by_filter/document_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchMultipleDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdateDocument(t *testing.T) {

	payload := strings.NewReader(`{
		"documentName" : "Name of document update",
		"documentType" : "Type of document update",
		"documentTrainingStatus" : "Du model training status of the document"
		}`)

	req := httptest.NewRequest("PUT", "/api/update/document_master/636df00176f1777049e1a282", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDeleteDocument(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/api/delete/document_master/636df00176f1777049e1a282", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// func TestListOfAllDocuments(t *testing.T) {

// 	req := httptest.NewRequest("GET", "/api/list_all/document_master", nil)
// 	req.Header.Add("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(ListOfAllDocuments)

// 	handler.ServeHTTP(w, req)
// 	resp := w.Result()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(resp.StatusCode)
// 	fmt.Println(resp.Header.Get("Content-Type"))
// 	fmt.Println(string(body))

// 	// Check the status code is what we expect.
// 	if status := w.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// }

// ////////////////////////////////////////////////////////// NEGATIVE TESTING //////////////////////////////////////////////////////////
func TestAddDocumentInvalidPayload(t *testing.T) {
	payload := strings.NewReader(`{
		"test" : "testing of document"
		}`)

	req := httptest.NewRequest("POST", "/api/add/document_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestAddDublicateDocumentName(t *testing.T) {
	payload := strings.NewReader(`{
		"documentName" : "Name of document",
		"documentType" : "Type of document",
		"documentTrainingStatus" : "Du model training status of the document"
		}`)

	req := httptest.NewRequest("POST", "/api/add/document_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)
	// Check the status code is what we expect.
	if successMsg := "Document already present"; successMsg != dataBody.SuccessMsg {
		t.Errorf("handler returned wrong status code: got %v want %v",
			dataBody.SuccessMsg, successMsg)
	}
}

func TestAddWithoutDocumentName(t *testing.T) {
	payload := strings.NewReader(`{
		"documentType" : "Type of document",
		"documentTrainingStatus" : "Du model training status of the document"
		}`)

	req := httptest.NewRequest("POST", "/api/add/document_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestAddWithoutDocumentType(t *testing.T) {
	payload := strings.NewReader(`{
		"documentName" : "Name of document",
		"documentTrainingStatus" : "Du model training status of the document"
		}`)

	req := httptest.NewRequest("POST", "/api/add/document_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestAddWithoutDocumentTrainingStatus(t *testing.T) {
	payload := strings.NewReader(`{
		"documentName" : "Name of document",
		"documentType" : "Type of document"
		}`)

	req := httptest.NewRequest("POST", "/api/add/document_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestAddInvalidPayload(t *testing.T) {
	payload := strings.NewReader(`{
		"document" : "Name of document",
		"document" : "Type of document"
		}`)

	req := httptest.NewRequest("POST", "/api/add/document_master", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(AddDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)
	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSearchOneDocumentByIdWithoutId(t *testing.T) {

	req := httptest.NewRequest("GET", "/api/search_one/document_master/", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(SearchOneDocumentById)
	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSearchOneDocumentByIdWithInvalidId(t *testing.T) {

	req := httptest.NewRequest("GET", "/api/search_one/document_master/12345", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(SearchOneDocumentById)
	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSearchOneDocumentByNameWithoutName(t *testing.T) {

	req := httptest.NewRequest("GET", "/api/search_one/document_master/", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(SearchOneDocumentByName)
	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSearchOneDocumentByNameWithInvalidName(t *testing.T) {

	req := httptest.NewRequest("GET", "/api/search_one/document_master/12345", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(SearchOneDocumentByName)
	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// func TestSearchMultipleUserInvalidPayload(t *testing.T) {

// 	payload := strings.NewReader(`{
// 		"test" : "Name of document",
// 		"testType" : "Type of document"
// 		}`)

// 	req := httptest.NewRequest("POST", "/api/search_by_filter/document_master", payload)
// 	req.Header.Add("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SearchMultipleDocument)

// 	handler.ServeHTTP(w, req)
// 	resp := w.Result()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(resp.StatusCode)
// 	fmt.Println(resp.Header.Get("Content-Type"))
// 	fmt.Println(string(body))

// 	// Check the status code is what we expect.
// 	if status := w.Code; status != http.StatusAccepted {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// }

func TestUpdateDocumentWithoutId(t *testing.T) {

	payload := strings.NewReader(`{
		"documentName" : "Name of document update",
		"documentType" : "Type of document update",
		"documentTrainingStatus" : "Du model training status of the document"
		}`)

	req := httptest.NewRequest("PUT", "/api/update/document_master/", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestUpdateDocumentWithInvalidId(t *testing.T) {

	payload := strings.NewReader(`{
		"documentName" : "Name of document update",
		"documentType" : "Type of document update",
		"documentTrainingStatus" : "Du model training status of the document"
		}`)

	req := httptest.NewRequest("PUT", "/api/update/document_master/abc", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestUpdateDocumentWithoutDocumentName(t *testing.T) {

	payload := strings.NewReader(`{
		"documentType" : "Type of document update",
		"documentTrainingStatus" : "Du model training status of the document"
		}`)

	req := httptest.NewRequest("PUT", "/api/update/document_master/636df00176f1777049e1a282", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestUpdateDocumentWithoutDocumentType(t *testing.T) {

	payload := strings.NewReader(`{
		"documentName" : "Name of document update",
		"documentTrainingStatus" : "Du model training status of the document"
		}`)

	req := httptest.NewRequest("PUT", "/api/update/document_master/636df00176f1777049e1a282", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestUpdateDocumentWithoutDocumentTrainingStatus(t *testing.T) {

	payload := strings.NewReader(`{
		"documentName" : "Name of document update",
		"documentType" : "Type of document update"
		}`)

	req := httptest.NewRequest("PUT", "/api/update/document_master/636df00176f1777049e1a282", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateDocument)

	handler.ServeHTTP(w, req)

	var dataBody model.Response
	_ = json.NewDecoder(w.Body).Decode(&dataBody)
	fmt.Println(w.Code)
	fmt.Println(dataBody)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestDeleteDocumentByIdWithoutId(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/api/search_one/document_master/", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(DeleteDocument)
	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestDeleteWithInvalidId(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/api/search_one/document_master/12345", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(DeleteDocument)
	handler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// func TestListOfAllDocumentsNegative(t *testing.T) {

// 	req := httptest.NewRequest("GET", "/api/search_one/document_master", nil)
// 	req.Header.Add("Content-Type", "application/json")

// 	w := httptest.NewRecorder()

// 	handler := http.HandlerFunc(ListOfAllDocuments)
// 	handler.ServeHTTP(w, req)

// 	fmt.Println(w.Code)
// 	fmt.Println(w.Body)

// 	// Check the status code is what we expect.
// 	if status := w.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusBadRequest)
// 	}
// }
