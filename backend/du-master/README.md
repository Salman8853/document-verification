# document-verification

		Schema Structure 
		Database Name: DocumentVerification

## Client Master Table

_id
clientName
clientCode
address
logoPath
functionalEntity
status
creationDateTime

## API Reference

#### Add new client

```http
  POST localhost:8080/api/add/client_master
```

```curl
curl --location --request POST 'http://localhost:8080/api/add/client_master' \
--header 'Content-Type: application/json' \
--data-raw '{
    "clientName":"",
    "clientCode":"",
    "address":"",
    "logoPath":"",
    "functionalEntity":"",
    "status":"A/I"
}'
```

#### Search one client by id

```http
  GET localhost:8080/api/find_client_by_id/client_master/{id}
```
```curl
  curl --location --request GET 'localhost:8080/api/find_client_by_id/client_master/{id}' \
--header 'Content-Type: application/json' \
```

#### Search one client by name

```http
  GET localhost:8080/api/find_client_by_name/client_master/{clientName}
```
```curl
  curl --location --request GET 'localhost:8080/api/find_client_by_name/client_master/{clientName}' \
--header 'Content-Type: application/json' \
```

#### Search multiple clients
```http
  POST localhost:8080/api/search_multiple_filter/client_master/
```

```curl
curl --location --request POST 'localhost:8080/api/search_multiple_filter/client_master/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "clientName":"client name",
    "clientCode":"clientcode",
    "address":"",
    "logoPath":"",
    "functionalEntity":"",
    "status":"A/I"
}'
```

#### Update client
```http
  PUT localhost:8080/api/update_client_by_id/client_master/{id}
```

```curl
curl --location --request PUT 'localhost:8080/api/update_client_by_id/client_master/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "clientName":"",
    "clientCode":"",
    "address":"",
    "logoPath":"",
    "functionalEntity":"",
    "status":"A/I"
}'
```

#### Delete client

```http
  DELETE localhost:8080/api/delete-client_by_id/{id}
```
```curl
  curl --location --request DELETE 'localhost:8080/api/delete-client_by_id/{id}' \
--header 'Content-Type: application/json' \
```

#### Deactivate Document

```http
  PUT localhost:8080/api/deactivate_by_id/client_master/{id}
```

```curl
curl --location --request PUT 'localhost:8080/api/deactivate_by_id/client_master/{id} \
--header 'Content-Type: application/json' \
--data-raw '{
"status" : "active/ deactive"
}'
```

#### find all client

```http
  GET localhost:8080/api/find_all/client_master
```
```curl
  curl --location --request GET 'localhost:8080/api/find_all/client_master \
--header 'Content-Type: application/json' \
```
#### list of all client

```http
  GET localhost:8080/api/list_of_client/client_master/
```curl
  curl --location --request GET 'localhost:8080/api/list_of_client/client_master/ \
--header 'Content-Type: application/json' \
```
## Authors

- [Aarti Kumari](https://www.github.com/gic-aartikr)


## Document Master Table

_id
documentName
documentType
status
documentTrainingStatus
creationDateTime


## API Reference

#### Add new document

```http
  POST localhost:9000/api/add/document_master
```

```curl
curl --location --request POST 'localhost:9000/api/add/document_master' \
--header 'Content-Type: application/json' \
--data-raw '{
    "documentName": " Name of document test",
    "documentType": "Type of document test",
    "documentDiscription": "short discription about the uploaded document",
    "documentTrainingStatus": "Du model training status of the document"
}'
```

#### Search one document by id

```http
  GET localhost:9000/api/search_one_by_id/document_master{id}
```
```curl
curl --location --request GET 'localhost:9000/api/search_one_by_id/document_master/{id}' \
--header 'Content-Type: application/json'
```

#### Search one document by name

```http
  GET localhost:9000/api/search_one_by_name/document_master/{documentName}
```
```curl
curl --location --request GET 'localhost:9000/api/search_one_by_name/document_master/{documentName}' \
--header 'Content-Type: application/json'
```

#### Search multiple Documents
```http
  POST localhost:9000/api/search_by_filter/document_master
```

```curl
curl --location --request POST 'localhost:9000/api/search_by_filter/document_master' \
--header 'Content-Type: application/json' \
--data-raw '{
    "documentName": "Name of document test",
    "documentType": "Type of document",
    "status": "A",
    "documentDiscription": "short discription about the uploaded document",
    "documentTrainingStatus": "Du model training status of the document"
}'
```

#### Update document
```http
  PUT localhost:9000/api/update/document_master/{id}
```

```curl
curl --location --request PUT 'localhost:9000/api/update/document_master/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "documentName": " Name of document test",
    "documentType": "Type of document test",
    "documentDiscription": "short discription about the uploaded document",
    "documentTrainingStatus": "Du model training status of the document",
    "status": "A/I"
}'

#### Delete Document

```http
  DELETE localhost:9000/api/delete/document_master/{id}
```
```curl
curl --location --request DELETE 'localhost:9000/api/delete/document_master/{id}' \
--header 'Content-Type: application/json'
```

#### Deactivate Document

```http
  PUT localhost:9000/api/update/document_master/{id}
```

```curl
curl --location --request PUT 'localhost:9000/api/update/document_master/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{
"status" : "I"
}'
```

#### Update Document Training Status

```http
  PUT localhost:9000/api/update/document_master/{id}
```

```curl
curl --location --request PUT 'localhost:9000/api/update/document_master/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "documentTrainingStatus": "Du model training status of the document"
}'
```

#### List of all active document

```http
  GET localhost:9000/api/list_active/document_master
```
```curl
curl --location --request GET 'localhost:9000/api/list_all/document_master'
```
## Authors

- [Vidhi Goel](https://www.github.com/gic-vidhi)

3. documentFieldMaster

 >_id, documentFieldName, status, creationDateTime

## API

Add new document

    curl --location --request POST 'http://localhost:8080/api/add/document_field_master/' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "documentFieldName": "string"
    }'

Fetch document by id

    curl --location -g --request GET 'http://localhost:8080/api/search/document_field_master/{id}'

Fetch all document having status 'A'

    curl --location --request GET 'http://localhost:8080/api/search-all/document_field_master/'

Search filter

    curl --location --request POST 'http://localhost:8080/api/search-filter/document_field_master/' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "id": "string",
        "documentFieldName": "string",
        "status": "string"
    }'
Update document by id

    curl --location -g --request PUT 'http://localhost:8080/api/update/document_field_master/{id}' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "documentFieldName": "string",
		"status": "string"
    }'

Deactivate document by id

    curl --location -g --request DELETE 'http://localhost:8080/api/delete/document_field_master/{id}'

    fetch list of document field names having status 'A'

    curl --location --request GET 'http://localhost:8080/api/list_all/document_field_master/'

## Authors

- [Ramashankar ](https://www.github.com/gic-ramashankar)

4. documentFieldMapping

_id
documentFieldName
documentName
status
creationDateTime


5. Client Document Mapping
_id
clientMaster
documentName
status
creationDateTime


Add ClientDocumentMapping

    curl --location --request POST 'http://localhost:8080/api/add-client-doc-map' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "clientMaster": "37637464",
        "documentName": "education",
        "status": "A"
    }'

Search ClientDocumentMapping by Id

    curl --location --request GET 'http://localhost:8080/api/search-cliDocMap/636e0ce4d355f89ca4c7165c'


Search filter ClientDocumentMapping

    curl --location --request GET 'http://localhost:8080/api/search-multiple-cdm' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "status": "A"
    }'

Deactivate ClientDocumentMapping

    curl --location --request PUT 'http://localhost:8080/api/delete-cdm/636e0ce4d355f89ca4c7165c'

Update ClientDocumentMapping

    curl --location --request PUT 'http://localhost:8080/api/update-cdm/636e0ce4d355f89ca4c7165c' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "status": "A"
    }'

Get All ClientDocumentMapping

       curl --location --request GET 'http://localhost:9000/api/get-all-cdm'

##    Authors
    Anurag

6. Client Document Field Mapping

> _id, clientMaster, documentFieldName, documentName, status, creationDateTime

## API

Add new document

     curl --location --request POST 'http://localhost:8080/api/add/client_document_field_mapping/' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "documentName": "String- ObjectID",
        "clientName": "String- ObjectID",
        "documentFieldName": "String- ObjectID"
    }'

Fetch document by id

    curl --location --request GET 'http://localhost:8080/api/search/client_document_field_mapping/{id}'
    
Fetch all document

    curl --location --request GET 'http://localhost:8080/api/search-all/client_document_field_mapping/'

Search filter

    curl --location --request POST 'http://localhost:8080/api/search-filter/client_document_field_mapping/' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "documentName": "String- ObjectID",
        "clientName": "String- ObjectID",
        "documentFieldName": "String- ObjectID",
        "status": "String"
    }'
    
Update document by id

     curl --location --request PUT 'http://localhost:8080/api/update/client_document_field_mapping/636e12a99b8daf1289165c91' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "documentName": "String- ObjectID",
        "clientName": "String- ObjectID",
        "documentFieldName": "String- ObjectID",
        "status": "A"
    }'
    
Deactivate document by id

    curl --location --request DELETE 'http://localhost:8080/api/delete/client_document_field_mapping/{id}'

    ## Authors

- [Ramashankar ](https://www.github.com/gic-ramashankar)

## Xls uplaod Api

## API Reference

#### Add xls data in DB

```http
  POST http://localhost:9000/api/xls_transaction_case_creation/
```

```curl
curl --location --request POST 'http://localhost:9000/api/xls_transaction_case_creation/'
```

## Create transaction from UI

    curl --location --request POST 'http://localhost:8080/api/create-transaction/' \
    --form 'FirstName="test"' \
    --form 'MiddleName="dummy"' \
    --form 'LastName="kumar2"' \
    --form 'Email="d@gmail.com"' \
    --form 'Dob="10-10-2022"' \
    --form 'ContactNumber="234567876"' \
    --form 'FileName=@"/path/to/file"'