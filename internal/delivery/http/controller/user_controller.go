package controller

import (
	"github/golang-developer-technical-test/internal/dto"
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

func NewUseController(logger *logrus.Logger) *UserController {
	return &UserController{
		Log: logger,
	}
}

func (c *UserController) Register(e echo.Context) error {
	response := new(dto.JSONResponse)
	req := model.RegisterUserRequest{}

	if err := e.Bind(&req); err != nil {
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
		c.Log.Warnf("Failed to get image_selfie file: %+v", err)
		response.StatusCode = http.StatusBadRequest
		response.Message = "ImageSelfie file upload failed"
		response.Data = nil
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}

	if file, err := e.FormFile("image_ktp"); err == nil {
		req.ImageKtp = file
	} else {
		c.Log.Warnf("Failed to get image_ktp file: %+v", err)
		response.StatusCode = http.StatusBadRequest
		response.Message = "ImageSelfie file upload failed"
		response.Data = nil
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}

	c.Log.Println(req)

	return errtrace.Wrap(e.JSON(response.StatusCode, response))
}
