package util

import (
	"github/golang-developer-technical-test/internal/model"
)

func CreateResponse(sc int, msg string, data interface{}) model.JSONResponse {
	return model.JSONResponse{
		StatusCode: sc,
		Message:    msg,
		Data:       data,
	}
}

func CreateResponseGenerics[T any](sc int, msg string, data T) model.JSONResponseGenerics[T] {
	return model.JSONResponseGenerics[T]{
		StatusCode: sc,
		Message:    msg,
		Data:       &data,
	}
}
