package domain

type ApiError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

func NewBindJsonError(cause string) ApiError {
	return ApiError{Error: "invalid_payload", Message: "Provided body payload is no valid", Cause: cause}
}

func NewApiError(error string, message string, cause string) ApiError {
	return ApiError{Error: error, Message: message, Cause: cause}
}
