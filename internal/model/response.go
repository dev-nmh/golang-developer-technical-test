package model

type JSONResponse struct {
	StatusCode int         `json:"status_code"`
	ErrorCode  string      `json:"error_code,omitempty"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type JSONResponseGenerics[T any] struct {
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code,omitempty"`
	Message    string `json:"message"`
	Data       *T     `json:"data"`
}
