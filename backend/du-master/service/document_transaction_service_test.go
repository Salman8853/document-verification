package service

import (
	"fmt"
	"testing"
)

func TestConvertDate(t *testing.T) {
	request := "2022-10-10"
	date, err := ConvertDate(request)
	got := date
	want := "2022-10-10 00:00:00 +0000 UTC"
	if err != nil {
		t.Errorf("Failed. Error Occurred %s\n", err)
		return
	}
	if got.String() == want {
		t.Logf("Pass. Expected %s, Got %s\n", want, got)
	} else {
		t.Errorf("Failed. Expected %s ,Got %s\n", want, got)
	}
}

func TestInsertDocFieldMaster2(t *testing.T) {
	request := ""
	date, err := ConvertDate(request)
	got := date
	want := "Empty"
	if err != nil {
		t.Errorf("Failed. Error Occurred, Got :%s\n", err.Error())
		return
	}
	if got.String() == want {
		t.Logf("Pass. Expected %s, Got %s\n", want, got)
	} else {
		t.Errorf("Failed. Expected %s ,Got %s\n", want, got)
	}
}

func TestGenerateTransactionId(t *testing.T) {
	request := 8
	id, err := GenerateTransactionId(request)
	got := id
	want := 16
	fmt.Println(got)
	if err != nil {
		t.Errorf("Failed. Error Occurred, Got :%s\n", err.Error())
		return
	}
	if len([]rune(got)) == want {
		t.Logf("Pass. Expected %d, Got %d\n", want, len(got))
	} else {
		t.Errorf("Failed. Expected %d ,Got %d\n", want, len(got))
	}
}

func TestUploadUserDocuments(t *testing.T) {
	// var uploadPath = "upload/userDocuments"
	filePath := "D:/GIC/document-verification/backend/du-master/test/upload"
	msg, _, err := UploadDocument(filePath)

	if err == nil {
		got := msg
		want := "Upload successful"
		fmt.Println("got:", got)
		fmt.Println("want:", want)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	if err != nil {
		got := msg
		want := ""
		fmt.Println("got:", got)
		fmt.Println("want:", want)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

}

func TestUploadUserDocumentsInvalidUploadDirectioy(t *testing.T) {
	// var uploadPath = "upload/userDocuments"
	filePath := "E:/GIC/document-verification/backend/du-master/test/upload"
	msg, _, err := UploadDocument(filePath)

	if err == nil {
		got := msg
		want := "Unable to access given directory"
		fmt.Println("got:", got)
		fmt.Println("want:", want)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	if err != nil {
		got := msg
		want := ""
		fmt.Println("got:", got)
		fmt.Println("want:", want)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

}

func TestUploadUserDocumentWithNoDocuments(t *testing.T) {
	// var uploadPath = "upload/userDocuments"
	filePath := "D:/GIC/document-verification/backend/du-master/test/test-empty"
	msg, _, err := UploadDocument(filePath)

	if err == nil {
		got := msg
		want := "No documents in the upload directory"
		fmt.Println("got:", got)
		fmt.Println("want:", want)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	if err != nil {
		got := msg
		want := ""
		fmt.Println("got:", got)
		fmt.Println("want:", want)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

}

func TestUploadUse(t *testing.T) {
	// var uploadPath = "upload/userDocuments"
	filePath := "D:/GIC/document-verification/backend/du-master/test/test-empty"
	msg, _, err := UploadDocument(filePath)

	if err == nil {
		got := msg
		want := "No documents in the upload directory"
		fmt.Println("got:", got)
		fmt.Println("want:", want)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	if err != nil {
		got := msg
		want := ""
		fmt.Println("got:", got)
		fmt.Println("want:", want)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

}
