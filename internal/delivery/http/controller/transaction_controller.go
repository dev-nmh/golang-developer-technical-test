package controller

import (
	"github/golang-developer-technical-test/internal/model"
	"github/golang-developer-technical-test/internal/usecase"
	"github/golang-developer-technical-test/internal/util"
	"net/http"

	"braces.dev/errtrace"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type TranscationLoanController struct {
	log     *logrus.Logger
	useCase *usecase.TranscationLoanUseCase
}

func NewTranscationLoanController(log *logrus.Logger, useCase *usecase.TranscationLoanUseCase) *TranscationLoanController {
	return &TranscationLoanController{
		log:     log,
		useCase: useCase,
	}
}

func (tlc TranscationLoanController) UserCreateLoanTransaction(e echo.Context) error {
	claim, err := util.NewClaimUtil(e)
	if err != nil {
		tlc.log.Warnf("Failed For Calim Token %+v", err)
		response := util.CreateResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}
	var userId uuid.UUID
	if userId, err = claim.GetUserId(); err != nil {
		tlc.log.Warnf("Failed For Autenticate as Admin %+v", err)
		response := util.CreateResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	} else {
		if err != nil {
			tlc.log.Warnf("Failed For Get User Id %+v", err)
			response := util.CreateResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
			return errtrace.Wrap(e.JSON(response.StatusCode, response))
		}
	}
	var req model.RequestLoan
	if err := e.Bind(&req); err != nil {
		response := new(model.JSONResponse)
		tlc.log.Warnf("Failed to parse request body : %+v", err)
		response.StatusCode = http.StatusBadRequest
		response.Message = "Bad Request"
		response.Data = nil
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}

	req.FkMsUser = userId
	req.CreateBy = userId

	if err := e.Validate(req); err != nil {
		response := new(model.JSONResponse)
		tlc.log.Warnf("Failed For Validate %+v", err)
		response.StatusCode = http.StatusBadRequest
		response.Message = "Data Not Valid"
		response.Data = err.Error()
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}
	data, cd, err := tlc.useCase.CreateLoan(e.Request().Context(), &req)
	if err != nil {
		tlc.log.Warnf("Failed to create Loan: %+v", err)
		response := new(model.JSONResponse)
		response.StatusCode = http.StatusInternalServerError
		if cd == 0 {
			response.StatusCode = cd
		}
		response.Message = err.Error()
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}
	response := util.CreateResponse(http.StatusOK, http.StatusText(http.StatusOK), data)
	return errtrace.Wrap(e.JSON(http.StatusOK, response))
}
