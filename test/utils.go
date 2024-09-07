package test

import (
	"bytes"
	"fmt"
	"github/golang-developer-technical-test/internal/constant"
	"github/golang-developer-technical-test/internal/entity"
	"github/golang-developer-technical-test/internal/model"
	"github/golang-developer-technical-test/internal/util"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/google/uuid"
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
		filePathToMultipart(KTP_IMG, "image_ktp", writer)
	}
	if withSelfie {
		filePathToMultipart(SELFIE_IMG, "image_selfie", writer)
	}

	err := writer.Close()
	if err != nil {
		return nil, "", err
	}
	return &b, writer.FormDataContentType(), nil
}

func filePathToMultipart(filePath string, formName string, writer *multipart.Writer) error {
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

	_, err = io.Copy(filePart, file)
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

func clearUsers() {
	err := db.Where("pk_ms_user is not null").Delete(&entity.MsUser{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func clearAccountUser() {
	err := db.Where("pk_ms_account is not null AND fk_ms_role = 1 ").Delete(&entity.MsAccount{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}
func clearAccountAdmin() {
	err := db.Where("pk_ms_account is not null AND fk_ms_role = 2 ").Delete(&entity.MsAccount{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func CreateAdmin() (*entity.MsAccount, error) {
	request := model.AccoutRequest{Email: "admin@mail.com", Password: "admin"}
	salt := uuid.New().String()
	if hashPassword, err := util.HashPassword(viperConfig.GetString("app.app_key") + request.Password + salt); err != nil {
		return nil, err
	} else {
		request.Password = hashPassword
	}
	accountId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	entity := &entity.MsAccount{
		PkMsAccount:  accountId,
		FkMsRole:     constant.USER_ROLES_ADMIN_INT,
		Password:     request.Password,
		PasswordSalt: salt,
		Email:        request.Email,
		Stamp: entity.Stamp{
			CreatedBy: accountId.String(),
			UpdatedBy: accountId.String(),
		},
	}
	if err := accountRepository.Create(db, entity); err != nil {
		return nil, err
	}

	return entity, nil
}
