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

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUseController(logger *logrus.Logger, useCase *usecase.UserUseCase) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *UserController) Register(e echo.Context) error {

	var req model.RegisterUserRequest

	if err := e.Bind(&req); err != nil {
		response := new(model.JSONResponse)
		c.Log.Warnf("Failed to parse request body : %+v", err)
		log.Println(err.Error())
		response.StatusCode = http.StatusBadRequest
		response.Message = "Bad Request"
		response.Data = nil
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}

	if file, err := e.FormFile("image_selfie"); err == nil {
		req.ImageSelfie = file
	} else {
		response := new(model.JSONResponse)
		c.Log.Warnf("Failed to get image_selfie file: %+v", err)
		response.StatusCode = http.StatusBadRequest
		response.Message = "image_selfie file upload failed"
		response.Data = nil
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}

	if file, err := e.FormFile("image_ktp"); err == nil {
		req.ImageKtp = file
	} else {
		response := new(model.JSONResponse)
		c.Log.Warnf("Failed to get image_ktp file: %+v", err)
		response.StatusCode = http.StatusBadRequest
		response.Message = "image_ktp file upload failed"
		response.Data = nil
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}

	if err := e.Validate(req); err != nil {
		response := new(model.JSONResponse)
		c.Log.Warnf("Failed For Validate %+v", err)
		response.StatusCode = http.StatusBadRequest
		response.Message = "Data Not Valid"
		response.Data = nil
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}
	record, err := c.UseCase.Create(e.Request().Context(), &req)
	if err != nil {
		c.Log.Warnf("Failed to create user: %+v", err)
		response := new(model.JSONResponse)
		response.StatusCode = http.StatusInternalServerError
		response.Message = "Failed to create user"
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}
	response := new(model.JSONResponseGenerics[model.UserResponseDetail])
	response.StatusCode = http.StatusCreated
	response.Message = "User created successfully"
	response.Data = record
	return errtrace.Wrap(e.JSON(response.StatusCode, response))
}
