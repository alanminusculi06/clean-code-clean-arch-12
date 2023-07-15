package domain

import "net/http"

type ApiError struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

func NewBindJsonError(cause string) ApiError {
	return ApiError{Error: "invalid_payload", Message: "Provided body payload is no valid", Cause: cause, Status: http.StatusBadRequest}
}

func NewApiError(error string, message string, cause string, status int) *ApiError {
	return &ApiError{Error: error, Message: message, Cause: cause, Status: status}
}

func NewUnprocessableEntityError(error string, message string, cause string) *ApiError {
	return &ApiError{Error: error, Message: message, Cause: cause, Status: http.StatusUnprocessableEntity}
}
