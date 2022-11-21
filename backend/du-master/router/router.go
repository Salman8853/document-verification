package router

import (
	"du-master/controller"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	//////////////////////////////////////////////////////////////// document_field_mapping //////////////////////////////////////////////////////////////////
	router.Handle("/api/add/document_field_mapping", cors(http.HandlerFunc(controller.AddNewDocField)))
	router.Handle("/api/search_all/document_field_mapping", cors(http.HandlerFunc(controller.FindAllDocFieldMapping)))
	router.Handle("/api/search_one_by_id/document_field_mapping/", cors(http.HandlerFunc(controller.FindDocFieldMappingById)))
	router.Handle("/api/delete/document_field_mapping/", cors(http.HandlerFunc(controller.DeleteDocFieldMappingById)))
	router.Handle("/api/update/document_field_mapping/", cors(http.HandlerFunc(controller.UpdateDocFieldMappingById)))
	router.Handle("/api/search_by_filter/document_field_mapping", cors(http.HandlerFunc(controller.SearchDocFieldMappingByFilter)))

	//////////////////////////////////////////////////////////////// client_document_mapping////////////////////////////////////////////////////////////////
	router.Handle("/api/add-client-doc-map", cors(http.HandlerFunc(controller.AddClientDocMap)))
	router.Handle("/api/search-cliDocMap/", cors(http.HandlerFunc(controller.SearchByCliDocMapID)))

	router.Handle("/api/search-multiple-cdm", cors(http.HandlerFunc(controller.SearchMultipleCDM)))
	router.Handle("/api/delete-cdm/", cors(http.HandlerFunc(controller.DeactivateCDM)))
	router.Handle("/api/update-cdm/", cors(http.HandlerFunc(controller.UpdateCDM)))
	router.Handle("/api/get-all-cdm", cors(http.HandlerFunc(controller.GetAllCDM)))

	//////////////////////////////////////////////////////////////// document_field_master////////////////////////////////////////////////////////////////
	router.Handle("/api/add/document_field_master/", cors(http.HandlerFunc(controller.InsertDocFieldMaster)))
	router.Handle("/api/search/document_field_master/", cors(http.HandlerFunc(controller.FindDocFieldMasterByID)))
	router.Handle("/api/search-all/document_field_master/", cors(http.HandlerFunc(controller.FindAllDocFieldMaster)))
	router.Handle("/api/delete/document_field_master/", cors(http.HandlerFunc(controller.DeactivateDocFieldMasterByID)))
	router.Handle("/api/search-filter/document_field_master/", cors(http.HandlerFunc(controller.SearchFilterDocFieldMaster)))
	router.Handle("/api/update/document_field_master/", cors(http.HandlerFunc(controller.UpdateDocFieldMasterByID)))
	router.Handle("/api/list_all/document_field_master/", cors(http.HandlerFunc(controller.FindAllDocumentFieldName)))

	//////////////////////////////////////////////////////////////// client_document_field_mapping////////////////////////////////////////////////////////////////
	router.Handle("/api/add/client_document_field_mapping/", cors(http.HandlerFunc(controller.InsertClientDocFieldMapping)))
	router.Handle("/api/search/client_document_field_mapping/", cors(http.HandlerFunc(controller.FindClientDocFieldMappingByID)))
	router.Handle("/api/search-all/client_document_field_mapping/", cors(http.HandlerFunc(controller.FindAllClientDocFieldMapping)))
	router.Handle("/api/delete/client_document_field_mapping/", cors(http.HandlerFunc(controller.DeactivateClientDocFieldMappingByID)))
	router.Handle("/api/search-filter/client_document_field_mapping/", cors(http.HandlerFunc(controller.SearchFilterClientDocFieldMapping)))
	router.Handle("/api/update/client_document_field_mapping/", cors(http.HandlerFunc(controller.UpdateClientDocFieldMappingByID)))

	//////////////////////////////////////////////////////////////// document_master////////////////////////////////////////////////////////////////

	router.Handle("/api/add/document_master", cors(http.HandlerFunc(controller.AddDocument)))
	router.Handle("/api/search_one_by_id/document_master/", cors(http.HandlerFunc(controller.SearchOneDocumentById)))
	router.Handle("/api/search_one_by_name/document_master/", cors(http.HandlerFunc(controller.SearchOneDocumentByName)))
	router.Handle("/api/search_by_filter/document_master", cors(http.HandlerFunc(controller.SearchMultipleDocument)))
	router.Handle("/api/update/document_master/", cors(http.HandlerFunc(controller.UpdateDocument)))
	router.Handle("/api/delete/document_master/", cors(http.HandlerFunc(controller.DeleteDocument)))
	router.Handle("/api/list_all/document_master", cors(http.HandlerFunc(controller.ListOfAllDocuments)))

	//////////////////////////////////////////////////////////////// client_master////////////////////////////////////////////////////////////////
	router.Handle("/api/add/client_master", cors(http.HandlerFunc(controller.AddClientDetails)))
	router.Handle("/api/find_all/client_master", cors(http.HandlerFunc(controller.FindAllClient)))
	router.Handle("/api/find_client_by_id/client_master/", cors(http.HandlerFunc(controller.FindByClientId)))
	router.Handle("/api/find_client_by_name/client_master/", cors(http.HandlerFunc(controller.FindByClientName)))
	router.Handle("/api/search_multiple_filter/client_master/", cors(http.HandlerFunc(controller.SearchMultipleClient)))
	router.Handle("/api/update_client_by_id/client_master/", cors(http.HandlerFunc(controller.UpdateClient)))
	router.Handle("/api/delete-client_by_id/", cors(http.HandlerFunc(controller.DeleteClientById)))
	router.Handle("/api/deactivate_by_id/client_master/", cors(http.HandlerFunc(controller.DeactivateClientById)))
	router.Handle("/api/list_of_client/client_master/", cors(http.HandlerFunc(controller.ListOfAllClient)))

	// ---------------------Xls Transaction Api-------------------------------------
	router.Handle("/api/xls_transaction_case_creation/", cors(http.HandlerFunc(controller.XlsTransactionCaseCreation)))

	router.Handle("/api/create-transaction/", cors(http.HandlerFunc(controller.CreateTransaction)))
	return router
}

func cors(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCORSHeaders(w, r)
		if r.Method == "OPTIONS" {
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})

}

func setCORSHeaders(w http.ResponseWriter, r *http.Request) {

	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")

}
