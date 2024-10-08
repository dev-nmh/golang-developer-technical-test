package test

import (
	"bytes"
	"encoding/json"
	"github/golang-developer-technical-test/internal/constant"
	"github/golang-developer-technical-test/internal/model"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAccount(t *testing.T) {
	ClearAllTr()
	ClearMapUserTenor()
	clearUsers()
	clearAccountUser()
	CreateAdmin()
	CreateExtern()

	h := CreateAccountUser()
	d, _ := json.Marshal(h)

	request := httptest.NewRequest(http.MethodPost, "/"+constant.PREFIX_API+"/public/register", bytes.NewBuffer(d))
	request.Header.Set("Content-Type", ContentTypeJSON)
	request.Header.Set("Accept", "*/*")
	request.Header.Set("X-API-KEY", viperConfig.GetString("app.api_key"))
	response := httptest.NewRecorder()

	App.ServeHTTP(response, request)
	if assert.Equal(t, response.Result().StatusCode, http.StatusCreated) {
		bytes, err := io.ReadAll(response.Body)
		assert.Nil(t, err)

		responseBody := new(model.JSONResponseGenerics[model.AccountResponse])
		err = json.Unmarshal(bytes, responseBody)
		assert.Nil(t, err)
		log.Println("======================================================")
		log.Println(responseBody)
		log.Println(*responseBody.Data)
		log.Println("======================================================")

		tokenUser = responseBody.Data.AccessToken

		if !assert.Equal(t, http.StatusCreated, responseBody.StatusCode) {
			return
		}
	}
}
func TestCreateUserProfile(t *testing.T) {

	birthDate, err := StringToTime("2000-12-05T14:41:50+07:00")
	assert.Nil(t, err)
	requestBody := model.UserData{
		Nik:        "1244563345672135",
		BirthDate:  *birthDate,
		BirthPlace: "Jakarta",
		FullName:   "Annisa Data 1",
		LegalName:  "Annisa Data 2",
		Salary:     10000000,
	}

	body, ContentType, err := createMultipartUser(requestBody, true, true)
	assert.Nil(t, err)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/user/profile", body)
	request.Header.Set("Content-Type", ContentType)
	request.Header.Set("Accept", "*/*")
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+tokenUser)
	response := httptest.NewRecorder()

	App.ServeHTTP(response, request)
	// assert.Nil(t, err)
	// userController.CreateProfile(c)
	// assert.Nil(t, err)
	if assert.Equal(t, response.Result().StatusCode, http.StatusCreated) {
		bytes, err := io.ReadAll(response.Body)
		assert.Nil(t, err)
		reqBody := new(model.JSONResponseGenerics[model.UserResponseDetail])
		err = json.Unmarshal(bytes, reqBody)
		assert.Nil(t, err)
		if !assert.Equal(t, http.StatusCreated, reqBody.StatusCode) {
			return
		}
		assert.Equal(t, requestBody.Nik, reqBody.Data.Nik)
		assert.Equal(t, requestBody.BirthDate, reqBody.Data.BirthDate)
		assert.Equal(t, requestBody.BirthPlace, reqBody.Data.BirthPlace)
		assert.Equal(t, requestBody.FullName, reqBody.Data.FullName)
		assert.Equal(t, requestBody.LegalName, reqBody.Data.LegalName)
		assert.Equal(t, requestBody.Salary, reqBody.Data.Salary)
		assert.NotEqual(t, reqBody.Data.ImageKtp, ".")
		assert.NotEqual(t, reqBody.Data.ImageSelfie, ".")
	}

}
