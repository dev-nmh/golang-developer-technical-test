package test

import (
	"encoding/json"
	"github/golang-developer-technical-test/internal/model"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	clearUsers()
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

	request := httptest.NewRequest(http.MethodPost, "/api/v1/users", body)
	request.Header.Set("Content-Type", ContentType)
	request.Header.Set("Accept", "*/*")
	response := httptest.NewRecorder()

	c := App.NewContext(request, response)
	assert.Nil(t, err)
	userController.Register(c)
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.JSONResponseGenerics[model.UserResponseDetail])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	if !assert.Equal(t, http.StatusCreated, responseBody.StatusCode) {
		return
	}
	assert.Equal(t, requestBody.Nik, responseBody.Data.Nik)
	assert.Equal(t, requestBody.BirthDate, responseBody.Data.BirthDate)
	assert.Equal(t, requestBody.BirthPlace, responseBody.Data.BirthPlace)
	assert.Equal(t, requestBody.FullName, responseBody.Data.FullName)
	assert.Equal(t, requestBody.LegalName, responseBody.Data.LegalName)
	assert.Equal(t, requestBody.Salary, responseBody.Data.Salary)
	assert.NotEqual(t, responseBody.Data.ImageKtp, ".")
	assert.NotEqual(t, responseBody.Data.ImageSelfie, ".")
}
