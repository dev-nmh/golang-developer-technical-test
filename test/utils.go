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

var tokenUser string
var tokenAdmin string
var tokenExtern string

var userId uuid.UUID

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
func clearAccountExtern() {
	err := db.Where("pk_ms_account is not null AND fk_ms_role = 3 ").Delete(&entity.MsAccount{}).Error
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
func CreateExtern() (*entity.MsAccount, error) {
	request := model.AccoutRequest{Email: "extern@mail.com", Password: "extern"}
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
		FkMsRole:     constant.USER_ROLES_EXTERN_INT,
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

func CreateAccountUser() model.AccoutRequest {
	return model.AccoutRequest{Email: "naufal@mail.com", Password: "user"}
}
func CreateAccountAdmin() model.AccoutRequest {
	return model.AccoutRequest{Email: "admin@mail.com", Password: "admin"}
}
func CreateAccountExtern() model.AccoutRequest {
	return model.AccoutRequest{Email: "extern@mail.com", Password: "extern"}
}
func CreateLoanApproval(UserId uuid.UUID) model.UserApproval {
	return model.UserApproval{
		UserTenor: []model.UserLimitTenor{
			{
				TenorId: "XYZ-TENOR-1",
				Amount:  100000,
			},
			{
				TenorId: "XYZ-TENOR-2",
				Amount:  100000,
			},
			{
				TenorId: "XYZ-TENOR-3",
				Amount:  1000000,
			},
			{
				TenorId: "XYZ-TENOR-4",
				Amount:  5000000,
			},
		},
		ApprovalId: 2,
		UserId:     userId,
	}
}

func CreateRequestLoanUser() model.RequestLoan {
	return model.RequestLoan{
		FkMsItemType:    uuid.MustParse("e8a8c8e5-6e27-11ef-b2e4-0242ac110002"),
		TenorId:         "XYZ-TENOR-4",
		ContractNumber:  "contract-number-1",
		AssetName:       "HONDA BRV",
		OtrAmount:       1000000.00,
		TransactionDate: time.Now(),
	}
}

func CreateRequestLoanExtern(sufix string, otrAmount float64) model.RequestLoan {
	return model.RequestLoan{
		FkMsItemType:    uuid.MustParse("e8a8c8e5-6e27-11ef-b2e4-0242ac110002"), // Example UUID, replace with actual value as needed
		TenorId:         "XYZ-TENOR-1",                                          // Example tenor ID, replace with actual value
		ContractNumber:  "contract-number-" + sufix,                             // Example contract number, replace with actual value
		AssetName:       "HONDA BRV",                                            // Example asset name, replace with actual value
		OtrAmount:       otrAmount,
		FkMsSource:      "CARMUD-SERVICE", // Example amount, replace with actual value
		TransactionDate: time.Now(),       // Example transaction date, replace with actual value
	}
}
func ClearAllTr() {
	ClearTrLoanDetail()
	ClearTrLoanHeader()

}
func ClearTrLoanDetail() {
	err := db.Where("pk_tr_loan_detail is not null").Delete(&entity.TrLoanDetail{}).Error
	if err != nil {
		log.Fatalf("Failed clear address data : %+v", err)
	}
}

func ClearTrLoanHeader() {
	err := db.Where("pk_tr_loan_header is not null").Delete(&entity.TrLoanHeader{}).Error
	if err != nil {
		log.Fatalf("Failed clear address data : %+v", err)
	}
}

func ClearMapUserTenor() {
	err := db.Where("pk_map_user_tenor is not null").Delete(&entity.MapUserTenor{}).Error
	if err != nil {
		log.Fatalf("Failed clear address data : %+v", err)
	}
}
