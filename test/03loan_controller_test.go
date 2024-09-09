package test

import (
	"bytes"
	"context"
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

func TestLoanApproval(t *testing.T) {

	user := CreateAccountUser()
	acc, err := accountUseCase.Verify(context.Background(), &user)
	log.Println(err)
	accByte, err := json.Marshal(acc)
	log.Println(err)
	var account model.AccountResponse
	err = json.Unmarshal(accByte, &account)
	log.Println(err)

	user = CreateAccountAdmin()
	acc, err = accountUseCase.Verify(context.Background(), &user)
	log.Println(err)
	accByte, err = json.Marshal(acc)
	log.Println(err)
	var adminAccount model.AccountResponse
	err = json.Unmarshal(accByte, &adminAccount)
	log.Println(err)

	h := CreateLoanApproval(*account.UserId)
	d, _ := json.Marshal(h)
	log.Println("/" + constant.PREFIX_API + "/admin/user/" + account.UserId.String() + "/approval")
	request := httptest.NewRequest(http.MethodPost, "/"+constant.PREFIX_API+"/admin/user/"+account.UserId.String()+"/approval", bytes.NewBuffer(d))
	request.Header.Set("Content-Type", ContentTypeJSON)
	request.Header.Set("Accept", "*/*")
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+adminAccount.AccessToken)
	response := httptest.NewRecorder()
	App.ServeHTTP(response, request)
	if assert.Equal(t, http.StatusOK, response.Result().StatusCode) {
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
		userId = *responseBody.Data.UserId

		if !assert.Equal(t, http.StatusOK, responseBody.StatusCode) {
			return
		}
	}

}
