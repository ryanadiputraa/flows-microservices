package domain

import (
	"fmt"
	"strings"

	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/response"
)

type ResponseError struct {
	Code    int
	Message string
	ErrCode response.ERR_CODE
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
