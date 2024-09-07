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
	Log     *logrus.Logger
	UseCase *usecase.AccountUseCase
}

func NewAccountController(logger *logrus.Logger, useCase *usecase.AccountUseCase) *AccountController {
	return &AccountController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *AccountController) Based(e echo.Context) (*model.RegisterUserRequest, error) {
	var req model.RegisterUserRequest

	if err := e.Bind(&req); err != nil {
		response := new(model.JSONResponse)
		c.Log.Warnf("Failed to parse request body : %+v", err)
		log.Println(err.Error())
		response.StatusCode = http.StatusBadRequest
		response.Message = "Bad Request"
		response.Data = nil
		return nil, errtrace.Wrap(e.JSON(response.StatusCode, response))
	}

	if err := e.Validate(req); err != nil {
		response := new(model.JSONResponse)
		c.Log.Warnf("Failed For Validate %+v", err)
		response.StatusCode = http.StatusBadRequest
		response.Message = "Data Not Valid"
		response.Data = nil
		return nil, errtrace.Wrap(e.JSON(response.StatusCode, response))
	}

	return &req, nil
}

func (a *AccountController) LoginUser(e echo.Context) error {

	return nil
}

func (a *AccountController) LoginAdmin(e echo.Context) error {
	return nil
}
