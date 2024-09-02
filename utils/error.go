package utils

import "encoding/json"

type APIError struct {
	Code       int    `json:"code"`        // Code 业务状态码
	HttpStatus int    `json:"http_status"` // HttpStatus http响应状态码
	Message    string `json:"message"`     // Message 错误信息
}

func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:       code,
		HttpStatus: code,
		Message:    message,
	}
}

func (e *APIError) SetHttpStatus(code int) *APIError {
	e.HttpStatus = code
	return e
}

func (e *APIError) Error() string {
	return e.Message
}

func (e *APIError) ToJSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}
