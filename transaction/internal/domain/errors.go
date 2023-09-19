package domain

import (
	"fmt"
	"strings"
)

type ERR_CODE string

const (
	INVALID_PARAMS        ERR_CODE = "invalid_params"
	UNAUTHENTICATED       ERR_CODE = "unauthenticated"
	FORBIDDEN             ERR_CODE = "forbidden"
	NOT_FOUND             ERR_CODE = "not_found"
	INTERNAL_SERVER_ERROR ERR_CODE = "internal_server_error"
	BAD_GATEWAY           ERR_CODE = "bad_gateway"
)

type ResponseError struct {
	Code    int
	Message string
	ErrCode ERR_CODE
	Errors  map[string][]string
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("Error %s", e.Message)
}

func IsDuplicateSQLError(err error) bool {
	errorMessage := err.Error()
	return strings.Contains(strings.ToLower(errorMessage), "duplicate") ||
		strings.Contains(strings.ToLower(errorMessage), "unique")
}
