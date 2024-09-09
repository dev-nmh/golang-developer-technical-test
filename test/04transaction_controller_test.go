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

func TestTransaction(t *testing.T) {
	ClearAllTr()
	user := CreateAccountUser()
	acc, err := accountUseCase.Verify(context.Background(), &user)
	log.Println(err)
	accByte, err := json.Marshal(acc)
	log.Println(err)
	var account model.AccountResponse
	err = json.Unmarshal(accByte, &account)
	log.Println(err)

	user = CreateAccountExtern()
	acc, err = accountUseCase.Verify(context.Background(), &user)
	log.Println(err)
	accByte, err = json.Marshal(acc)
	log.Println(err)
	var externAccount model.AccountResponse
	err = json.Unmarshal(accByte, &externAccount)
	log.Println(err)

	t.Run("Extern 0", func(t *testing.T) {
		t.Parallel()
		requestStruct := CreateRequestLoanExtern("0", 2000000.00)
		requestByte, err := json.Marshal(requestStruct)
		log.Println(err)
		request := httptest.NewRequest(http.MethodPost, "/"+constant.PREFIX_API+"/admin/user/loan/"+account.UserId.String(), bytes.NewBuffer(requestByte))
		request.Header.Set("Content-Type", ContentTypeJSON)
		request.Header.Set("Accept", "*/*")
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+externAccount.AccessToken)
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
	})

	t.Run("Web Services", func(t *testing.T) {
		t.Parallel()
		requestStruct := CreateRequestLoanUser()
		requestByte, err := json.Marshal(requestStruct)
		log.Println(err)
		request := httptest.NewRequest(http.MethodPost, "/"+constant.PREFIX_API+"/user/loan", bytes.NewBuffer(requestByte))
		request.Header.Set("Content-Type", ContentTypeJSON)
		request.Header.Set("Accept", "*/*")
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+account.AccessToken)
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
	})
	t.Run("Extern 1", func(t *testing.T) {
		t.Parallel()
		requestStruct := CreateRequestLoanExtern("2", 3000000.00)
		requestByte, err := json.Marshal(requestStruct)
		log.Println(err)
		request := httptest.NewRequest(http.MethodPost, "/"+constant.PREFIX_API+"/admin/user/loan/"+account.UserId.String(), bytes.NewBuffer(requestByte))
		request.Header.Set("Content-Type", ContentTypeJSON)
		request.Header.Set("Accept", "*/*")
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+externAccount.AccessToken)
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
	})
	t.Run("Extern 2", func(t *testing.T) {
		t.Parallel()
		requestStruct := CreateRequestLoanExtern("3", 2000000.00)
		requestByte, err := json.Marshal(requestStruct)
		log.Println(err)
		request := httptest.NewRequest(http.MethodPost, "/"+constant.PREFIX_API+"/admin/user/loan/"+account.UserId.String(), bytes.NewBuffer(requestByte))
		request.Header.Set("Content-Type", ContentTypeJSON)
		request.Header.Set("Accept", "*/*")
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+externAccount.AccessToken)
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
	})

	t.Run("Extern 3", func(t *testing.T) {
		t.Parallel()
		requestStruct := CreateRequestLoanExtern("4", 1500000.00)
		requestByte, err := json.Marshal(requestStruct)
		log.Println(err)
		request := httptest.NewRequest(http.MethodPost, "/"+constant.PREFIX_API+"/admin/user/loan/"+account.UserId.String(), bytes.NewBuffer(requestByte))
		request.Header.Set("Content-Type", ContentTypeJSON)
		request.Header.Set("Accept", "*/*")
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+externAccount.AccessToken)
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
	})

}
