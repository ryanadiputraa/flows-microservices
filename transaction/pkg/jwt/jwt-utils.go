package jwt

import (
	"net/http"
	"strings"

	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/domain"
)

func (s *service) ExtractJWTTokenHeader(header http.Header) (string, error) {
	headers := header.Get("Authorization")
	if headers == "" {
		return "", &domain.ResponseError{
			Code:    http.StatusUnauthorized,
			Message: "missing authorization header",
			ErrCode: domain.UNAUTHENTICATED,
			Errors:  nil,
		}
	}

	val := strings.Split(headers, " ")
	if len(val) < 2 || val[0] != "Bearer" {
		return "", &domain.ResponseError{
			Code:    http.StatusForbidden,
			Message: "invalid authroziation header",
			ErrCode: domain.FORBIDDEN,
			Errors:  nil,
		}
	}

	return val[1], nil
}
