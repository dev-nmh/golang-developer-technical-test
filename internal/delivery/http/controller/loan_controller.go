package controller

import (
	"github/golang-developer-technical-test/internal/constant"
	"github/golang-developer-technical-test/internal/model"
	"github/golang-developer-technical-test/internal/usecase"
	"github/golang-developer-technical-test/internal/util"
	"net/http"

	"braces.dev/errtrace"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type LoanController struct {
	log     *logrus.Logger
	useCase *usecase.LoanUseCase
}

func NewLoanController(log *logrus.Logger, useCase *usecase.LoanUseCase) *LoanController {
	return &LoanController{
		log:     log,
		useCase: useCase,
	}
}

func (loan LoanController) ApprovalUser(e echo.Context) error {
	loan.log.Println("Passing")
	claim, err := util.NewClaimUtil(e)
	if err != nil {
		loan.log.Warnf("Failed For Calim Token %+v", err)
		response := util.CreateResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}
	var AdminId uuid.UUID
	if roleId, err := claim.GetRole(); roleId != constant.USER_ROLES_ADMIN_INT && err != nil {
		loan.log.Warnf("Failed For Autenticate as Admin %+v", err)
		response := util.CreateResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	} else {
		AdminId, err = claim.GetId()
		if err != nil {
			loan.log.Warnf("Failed For Get User Id %+v", err)
			response := util.CreateResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
			return errtrace.Wrap(e.JSON(response.StatusCode, response))
		}
	}
	var req model.UserApproval

	if err := e.Bind(&req); err != nil {
		response := new(model.JSONResponse)
		loan.log.Warnf("Failed to parse request body : %+v", err)
		response.StatusCode = http.StatusBadRequest
		response.Message = "Bad Request"
		response.Data = nil
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}

	req.AdminId = AdminId
	if err := e.Validate(req); err != nil {
		response := new(model.JSONResponse)
		loan.log.Warnf("Failed For Validate %+v", err)
		response.StatusCode = http.StatusBadRequest
		response.Message = "Data Not Valid"
		response.Data = err.Error()
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}
	record, err := loan.useCase.ApprovalUser(e.Request().Context(), &req)
	if err != nil {
		loan.log.Warnf("Failed For Validate %+v", err)
		response := util.CreateResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		if httpError, ok := err.(*echo.HTTPError); ok {
			response.StatusCode = httpError.Code
			response.Message = httpError.Message.(string)
			return errtrace.Wrap(e.JSON(response.StatusCode, response))
		}
		return errtrace.Wrap(e.JSON(response.StatusCode, response))
	}
	loan.log.Println(record)

	response := util.CreateResponse(http.StatusOK, http.StatusText(http.StatusOK), req)
	return errtrace.Wrap(e.JSON(response.StatusCode, response))
}

func (loan LoanController) CreateTenor(e echo.Context) error {
	return errtrace.Wrap(e.JSON(http.StatusAccepted, nil))
}
