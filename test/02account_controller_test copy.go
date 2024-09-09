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

	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	h := CreateAccountUser()
	d, _ := json.Marshal(h)

	request := httptest.NewRequest(http.MethodPost, "/"+constant.PREFIX_API+"/public/auth", bytes.NewBuffer(d))
	request.Header.Set("Content-Type", ContentTypeJSON)
	request.Header.Set("Accept", "*/*")
	request.Header.Set("X-API-KEY", viperConfig.GetString("app.api_key"))
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

func TestLoginAdmin(t *testing.T) {
	h := CreateAccountAdmin()
	d, _ := json.Marshal(h)

	request := httptest.NewRequest(http.MethodPost, "/"+constant.PREFIX_API+"/public/auth", bytes.NewBuffer(d))
	request.Header.Set("Content-Type", ContentTypeJSON)
	request.Header.Set("Accept", "*/*")
	request.Header.Set("X-API-KEY", viperConfig.GetString("app.api_key"))
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

		tokenAdmin = responseBody.Data.AccessToken

		if !assert.Equal(t, http.StatusOK, responseBody.StatusCode) {
			return
		}
	}

}

func TestLoginExtern(t *testing.T) {
	h := CreateAccountExtern()
	d, _ := json.Marshal(h)

	request := httptest.NewRequest(http.MethodPost, "/"+constant.PREFIX_API+"/public/auth", bytes.NewBuffer(d))
	request.Header.Set("Content-Type", ContentTypeJSON)
	request.Header.Set("Accept", "*/*")
	request.Header.Set("X-API-KEY", viperConfig.GetString("app.api_key"))
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

		tokenExtern = responseBody.Data.AccessToken

		if !assert.Equal(t, http.StatusOK, responseBody.StatusCode) {
			return
		}
	}

}
