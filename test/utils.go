package test

import (
	"bytes"
	"fmt"
	"github/golang-developer-technical-test/internal/entity"
	"github/golang-developer-technical-test/internal/model"
	"mime/multipart"
	"os"
	"time"
)

func createMultipartUser(data model.UserData, withKtp bool, withSelfie bool) (*bytes.Buffer, string, error) {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	addField(writer, "NIK", data.Nik)
	addField(writer, "birth_date", data.BirthDate.Format(time.RFC3339))
	addField(writer, "birth_place", data.BirthPlace)
	addField(writer, "full_name", data.FullName)
	addField(writer, "legal_name", data.LegalName)
	addField(writer, "salary", fmt.Sprintf("%d", data.Salary))

	if withKtp {
		FilePathToMultipart(KTP_IMG, "image_ktp", writer)
	}
	if withSelfie {
		FilePathToMultipart(SELFIE_IMG, "image_selfie", writer)
	}

	err := writer.Close()
	if err != nil {
		return nil, "", err
	}
	return &b, writer.FormDataContentType(), nil
}

func FilePathToMultipart(filePath string, formName string, writer *multipart.Writer) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	fileInfo, err := file.Stat()

	if err != nil {
		return err
	}

	filePart, err := writer.CreateFormFile(formName, fileInfo.Name())
	if err != nil {
		return err
	}

	_, err = filePart.Write(make([]byte, fileInfo.Size()))
	if err != nil {
		return err
	}
	return nil
}

func StringToTime(strTime string) (*time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, strTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, err
	}

	return &parsedTime, err
}

func addField(w *multipart.Writer, fieldName, value string) {
	part, err := w.CreateFormField(fieldName)
	if err != nil {
		fmt.Println("Error creating form field:", err)
		return
	}
	_, err = part.Write([]byte(value))
	if err != nil {
		fmt.Println("Error writing field value:", err)
		return
	}
}

func ClearUsers() {
	err := db.Where("pk_ms_user is not null").Delete(&entity.MsUser{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}
