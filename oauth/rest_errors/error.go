package rest_errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

//RestErr : "common"
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
}

//NewRestError : generates the new error
func NewRestError(message string, status int, code string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Code:    code,
	}
}

//NewRestErrorFromBytes : creates the resterror domain by taking the paramter as a slice of bytes
func NewRestErrorFromBytes(byte []byte) (*RestErr, error) {
	var apiErr RestErr
	if err := json.Unmarshal(byte, &apiErr); err != nil {
		return nil, errors.New("Invalid json")
	}
	return &apiErr, nil
}

//NewBadRequestError : check the error is for bad request status
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Code:    "bad_request",
	}
}

//NewNotFoundError : check the error is for bad request status
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}

//NewUnauthorizedError : check the error is for unauthorized status
func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusUnauthorized,
		Code:    "unauthorized",
	}
}

//NewInternalServerError : check the error is caused by the server
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Code:    "internal_server_error",
	}
}
