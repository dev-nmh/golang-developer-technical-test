package controller

import (
	"github/golang-developer-technical-test/internal/model"
	"github/golang-developer-technical-test/internal/usecase"
	"log"
	"net/http"

	"braces.dev/errtrace"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AccountController struct {
	log     *logrus.Logger
	useCase *usecase.AccountUseCase
}

func NewAccountController(logger *logrus.Logger, useCase *usecase.AccountUseCase) *AccountController {
	return &AccountController{
		log:     logger,
		useCase: useCase,
	}
}

func (c *AccountController) Based(e echo.Context) (*model.AccoutRequest, error) {
	var req model.AccoutRequest

	if err := e.Bind(&req); err != nil {
		response := new(model.JSONResponse)
		c.log.Warnf("Failed to parse request body : %+v", err)
		log.Println(err.Error())
		response.StatusCode = http.StatusBadRequest
		response.Message = "Bad Request"
		response.Data = nil
		return nil, errtrace.Wrap(e.JSON(response.StatusCode, response))
	}

	if err := e.Validate(req); err != nil {
		response := new(model.JSONResponse)
		c.log.Warnf("Failed For Validate %+v", err)
		response.StatusCode = http.StatusBadRequest
		response.Message = "Data Not Valid"
		response.Data = nil
		return nil, errtrace.Wrap(e.JSON(response.StatusCode, response))
	}

	return &req, nil
}

func (c *AccountController) Register(e echo.Context) error {
	req, err := c.Based(e)
	if err != nil {
		return errtrace.Wrap(err)
	}
	record, err := c.useCase.Register(e.Request().Context(), req)
	if err != nil {
		c.log.Warnf("Failed to create user: %+v", err)
		response := new(model.JSONResponse)
		response.StatusCode = http.StatusInternalServerError
		response.Message = "Failed to create user"
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}
	response := new(model.JSONResponseGenerics[model.AccountResponse])
	response.StatusCode = http.StatusCreated
	response.Message = "Account created successfully"
	response.Data = record
	return errtrace.Wrap(e.JSON(response.StatusCode, response))
}

func (c *AccountController) Auth(e echo.Context) error {
	req, err := c.Based(e)
	if err != nil {
		return errtrace.Wrap(err)
	}
	record, err := c.useCase.Verify(e.Request().Context(), req)
	if err != nil {
		response := new(model.JSONResponse)
		response.StatusCode = http.StatusInternalServerError
		response.Message = "Failed to create user"
		c.log.Warnf("Failed to Verify user: %+v", err)
		if httpError, ok := err.(*echo.HTTPError); ok {
			response.StatusCode = httpError.Code
			response.Message = httpError.Message.(string)
			return errtrace.Wrap(e.JSON(response.StatusCode, response))
		}

		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}
	response := new(model.JSONResponse)
	response.StatusCode = http.StatusCreated
	response.Message = "Login success"
	response.Data = record
	return errtrace.Wrap(e.JSON(response.StatusCode, response))
}

func (c *AccountController) LoginAdmin(e echo.Context) error {
	return nil
}
