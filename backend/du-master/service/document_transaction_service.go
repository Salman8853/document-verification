package service

import (
	"context"
	"crypto/rand"
	"du-master/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocumentTransactionService struct {
	DocumentTransactionCollection *mongo.Collection
}

const directory = "D://"
const maxUploadSize = 10 * 1024 * 1024 // 10 mb
const destination = "excel/download/"

var fileName string

func (dt *DocumentTransactionService) InsertDT(files []*multipart.FileHeader, firstName, middleName, lastName, email, dob, contactNumber string) (interface{}, error) {
	var response model.DocumentTransactionDetails

	check, err := dt.CheckDuplicateInDTS(firstName, middleName, lastName, dob)
	if err != nil {
		return response, err
	}
	if check {
		return response, errors.New("duplicate transaction")
	}
	filesarray, err := UploadFile(files)
	if err != nil {
		return response, err
	}

	response, err = SetValueInModel(filesarray, firstName, middleName, lastName, email, dob, contactNumber)
	if err != nil {
		return response, err
	}

	id, err := dt.InsertInDB(response)
	if err != nil {
		return response, err
	}
	response.Id = id
	return response, nil
}

func (dt *DocumentTransactionService) CheckDuplicateInDTS(firstName, middleName, lastName, dob string) (bool, error) {
	dobDate, err := ConvertDate(dob)
	if err != nil {
		return true, err
	}
	filter := bson.D{primitive.E{Key: FieldStatus, Value: StatusActiveValue}}
	filter = append(filter, primitive.E{Key: "firstName", Value: firstName})
	if middleName != "" {
		filter = append(filter, primitive.E{Key: "middleName", Value: middleName})
	}
	filter = append(filter, primitive.E{Key: "lastName", Value: lastName})
	filter = append(filter, primitive.E{Key: "dob", Value: dobDate})

	cur, err := dt.DocumentTransactionCollection.Find(context.Background(), filter)
	if err != nil {
		return true, err
	}
	var res []model.DocumentTransactionDetails
	cur.All(context.Background(), &res)

	if len(res) == 0 {
		return false, nil
	}
	return true, nil
}

func (dt *DocumentTransactionService) InsertInDB(saveData model.DocumentTransactionDetails) (primitive.ObjectID, error) {
	cur, err := dt.DocumentTransactionCollection.InsertOne(context.Background(), saveData)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return cur.InsertedID.(primitive.ObjectID), err
}

func SetValueInModel(filesarray []model.UploadedDocuments, firstName, middleName, lastName, email, dob, contactNumber string) (model.DocumentTransactionDetails, error) {
	var reqBody model.DocumentTransactionDetails

	dobDate, err := ConvertDate(dob)
	if err != nil {
		return reqBody, err
	}
	transactionId, err := GenerateTransactionId(8)
	if err != nil {
		return reqBody, err
	}
	if transactionId == "" {
		return reqBody, errors.New("unable to generate transactionId")
	}

	reqBody.FirstName = firstName
	reqBody.MiddleName = middleName
	reqBody.LastName = lastName
	reqBody.Email = email
	reqBody.Dob = dobDate
	reqBody.ContactNumber = contactNumber
	reqBody.CreationDateTime = time.Now()
	reqBody.UploadedDocs = filesarray
	reqBody.Status = StatusActiveValue
	reqBody.TransactionId = transactionId

	return reqBody, err
}

func UploadFile(files []*multipart.FileHeader) ([]model.UploadedDocuments, error) {
	var filesDesc []model.UploadedDocuments

	for _, fileHeader := range files {
		var fileModel model.UploadedDocuments

		segments := strings.Split(fileHeader.Filename, ".")
		extension := segments[len(segments)-1]

		if !(extension == "pdf" || extension == "image") {
			return filesDesc, errors.New("invalid file format")
		}

		// if fileHeader.Size > maxUploadSize {
		// 	return filesDesc, errors.New("The uploaded file is too big: %s\n " + fileHeader.Filename)
		// }

		_, err := url.Parse(directory)
		if err != nil {
			log.Panicln(err)
			return filesDesc, errors.New("path not available")

		}
		file, err := fileHeader.Open()
		if err != nil {
			return filesDesc, err
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			return filesDesc, err
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return filesDesc, err
		}

		f, err := os.Create(directory + fileHeader.Filename)
		if err != nil {
			return filesDesc, err
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			return filesDesc, err
		}
		fileName := fileHeader.Filename
		fileModel.DocumentName = fileName
		fileModel.Size = fileHeader.Size
		fileModel.DocumentPath = directory + fileName
		fileModel.DateCreated = time.Now()
		filesDesc = append(filesDesc, fileModel)
	}
	return filesDesc, nil
}

func ConvertDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, errors.New("empty")
	}
	date, err := time.Parse("01-02-2006", dateStr)
	if err != nil {
		log.Println(err)
		return date, err
	}

	return date, nil
}

func GenerateTransactionId(length int) (string, error) {
	if length == 0 {
		return "", errors.New("please enter length to generate id")
	}
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

// -------------------xls Transaction-------------------------------------------

func UploadDocument(documentPath string) (string, []model.UploadedDocuments, error) {
	var uploadedDocumentDetails []model.UploadedDocuments
	var uploadedDocumentDetail model.UploadedDocuments
	fmt.Printf("Directory of documents to be uploaded:%s", documentPath)
	files, err := ioutil.ReadDir(documentPath)
	if err != nil {
		return "Unable to access given directory", uploadedDocumentDetails, err
	}
	if len(files) == 0 {
		return "No documents in the upload directory", uploadedDocumentDetails, nil
	}

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		return "Unable to make upload directory", uploadedDocumentDetails, err
	}

	for _, f := range files {
		fmt.Println(f.Name())
		ioutil.ReadFile(f.Name())
		file, err := os.Open(documentPath + "/" + f.Name())
		if err != nil {
			return "Unable to read file in the given directory", uploadedDocumentDetails, err
		}
		uploadedDocumentDetail.Size = f.Size()
		uploadedDocumentDetail.DocumentName = f.Name()
		uploadedDocumentDetail.DocumentPath = documentPath
		uploadedDocumentDetail.DateCreated = time.Now()
		fmt.Println(uploadedDocumentDetail)
		defer file.Close()

		dst, err := os.Create(fmt.Sprintf("./uploads/%s", filepath.Ext(f.Name())))
		if err != nil {
			return "Unable to access upload directory", uploadedDocumentDetails, err
		}

		defer dst.Close()
		// Copy the uploaded file to the filesystem
		// at the specified destination
		_, err = io.Copy(dst, file)

		if err != nil {
			return "Unable to upload file in the given directory", uploadedDocumentDetails, err
		}

		uploadedDocumentDetails = append(uploadedDocumentDetails, uploadedDocumentDetail)
		fmt.Println("Upload successful")
	}
	return "Upload successful", uploadedDocumentDetails, nil
}

func (xl *DocumentTransactionService) XlsTransactionCaseCreation(files []*multipart.FileHeader) (string, error) {
	var excelJson []*map[string]interface{}
	var err error
	// var jsonStr []byt
	var data []*model.DocumentTransactionRequest
	var exedata []model.ExcelTransactionRequest

	excelJson, err = xl.XlsTransactionConvertToJson(files)

	fmt.Println("excelJson:", excelJson)
	if err != nil {
		return "", errors.New("excel unable to convert into json")
	}
	jsonBytes, err := json.MarshalIndent(excelJson, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("jsonBytes:", jsonBytes)

	if reflect.ValueOf(excelJson).Kind() != reflect.Slice {
		fmt.Print("required parameter LineItems is not an array")

	}

	s := reflect.ValueOf(excelJson)

	if s.Len() == 0 {
		fmt.Print("required parameter LineItems can not be empty")

	}
	fmt.Println("s:", s)
	// json.Unmarshal(jsonBytes, &data)

	jsondata, err := json.Marshal(excelJson)
	if err != nil {
		log.Println(err)
	}

	ss := string(jsondata)
	json.Unmarshal([]byte(ss), &exedata)

	fmt.Println("exedata:", exedata)

	for i := range exedata {

		firstName := exedata[i].FIRSTNAME
		lastName := exedata[i].LASTNAME
		middleName := exedata[i].MIDDLENAME
		email := exedata[i].EMAIL
		dob := exedata[i].DOB
		contactNo := exedata[i].CONTACTNUMBER
		docPath := exedata[i].DOCUMENTPATH
		fmt.Println("cNo:", contactNo)
		// DocumentTransactionRequest :=
		documentTransactionReque := model.DocumentTransactionRequest{FirstName: firstName, LastName: lastName, MiddleName: middleName, Email: email, Dob: dob, ContactNumber: contactNo, DocPath: docPath}
		data = append(data, &documentTransactionReque)
	}
	referenceId, validateField, err := xl.ValidateMandatoryField(data)

	if validateField != "" {
		return validateField, err
	}

	return referenceId, err
}

func (xl *DocumentTransactionService) ValidateMandatoryField(data []*model.DocumentTransactionRequest) (string, string, error) {
	var errList []*model.ExcelerrList
	var response model.DocumentTransactionDetails
	var errstrings []string
	var errListStr string
	var filesarray []model.UploadedDocuments
	var err error

	referenceId, err := GenerateTransactionId(8)

	for i := range data {
		// create a slice for th
		firstName := data[i].FirstName
		lastName := data[i].LastName
		middleName := data[i].MiddleName
		email := data[i].Email
		dob := data[i].Dob
		contactNumber := data[i].ContactNumber
		docPath := data[i].DocPath

		if firstName == "" {
			errstr := "Please provide FirstName"
			errstrings = append(errstrings, errstr)
		}

		if lastName == "" {
			errstr := "Please provide lastName"
			errstrings = append(errstrings, errstr)
		}
		if email == "" {
			errstr := "Please provide Email"
			errstrings = append(errstrings, errstr)
		}
		if dob == "" {
			errstr := "Please provide Date Of Birth"
			errstrings = append(errstrings, errstr)
		}
		if docPath == "" {
			errstr := "Please provide DocumentPath"
			errstrings = append(errstrings, errstr)
		}

		if docPath != "" {
			_, filesarray, err = UploadDocument(docPath)
			if err != nil {
				errstrings = append(errstrings, err.Error())
			}
		}
		s := reflect.ValueOf(errstrings)
		if s.Len() != 0 {
			fmt.Printf("error in row : %v ", i+2)
			fmt.Println("errstrings:", errstrings)
			excelErrL := model.ExcelerrList{RowNo: i, ExcelErr: errstrings}
			errList = append(errList, &excelErrL)
			jsonErrList, err := json.Marshal(errList)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("jata:", jsonErrList)

			errListStr = string(jsonErrList)
			return "", errListStr, err
		}

		// 	filesarray, err := UploadDocument(docPath)
		// if err != nil {
		// 	return "", err
		// }

		check, err := xl.CheckDuplicateInDTS(firstName, middleName, lastName, dob)
		if err != nil {
			return "", "", err
		}
		if check {
			return "", "", errors.New("Duplicate Transaction")
		}
		response, err = SetValueInModel(filesarray, firstName, middleName, lastName, email, dob, contactNumber)
		if err != nil {
			return "", "", err
		}
		response.ReferenceId = referenceId
		id, err := xl.InsertInDB(response)
		if err != nil {
			return "", "", err
		}
		response.Id = id
	}

	return referenceId, "", nil
}

func (xl *DocumentTransactionService) XlsTransactionConvertToJson(files []*multipart.FileHeader) ([]*map[string]interface{}, error) {

	// var data []*model.DocumentTransactionDetails
	var response []*map[string]interface{}

	for _, fileHeader := range files {
		if fileHeader.Size > maxUploadSize {
			// fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename)
			return response, errors.New("The uploaded image is too big")
		}

		// Open the file
		file, err := fileHeader.Open()
		if err != nil {

			return response, errors.New("")
		}
		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			return response, errors.New("Error while read file")
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return response, err
		}

		err = os.MkdirAll(destination, os.ModePerm)
		if err != nil {
			return response, err
		}

		f, err := os.Create(destination + fileHeader.Filename)
		if err != nil {
			return response, err
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			return response, errors.New("Error while copy file")
		}
		fileName = fileHeader.Filename
	}

	// excelToJson(w, fileName)
	response, err := xl.conversion2(fileName)
	if err != nil {
		return response, err
	}
	return response, nil

}

func (xl *DocumentTransactionService) conversion2(file string) ([]*map[string]interface{}, error) {

	var (
		headers []string

		result []*map[string]interface{}

		wb = new(excelize.File)

		err error

		sheetName string
	)

	wb, err = excelize.OpenFile(destination + file)

	if err != nil {

		return result, errors.New("Error while open file")

	}

	i := 0

	for _, sheet := range wb.GetSheetMap() {

		if i == 0 {

			sheetName = sheet

		}

		i++

	}

	rows := wb.GetRows(sheetName)

	headers = rows[0]

	for _, row := range rows[1:] {

		var tmpMap = make(map[string]interface{})

		for j, v := range row {

			tmpMap[strings.Join(strings.Split(headers[j], " "), "")] = v

		}

		result = append(result, &tmpMap)

	}

	// respondWithJson(w, http.StatusAccepted, result)

	return result, nil

}
