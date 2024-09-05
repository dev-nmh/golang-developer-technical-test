package dto

import "github.com/labstack/echo/v4"

type JSONResponse struct {
	StatusCode int       `json:"status_code"`
	ErrorCode  string    `json:"error_code"`
	Message    string    `json:"message"`
	Data       *echo.Map `json:"data"`
}
