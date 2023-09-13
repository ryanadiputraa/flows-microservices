package response

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/domain"
)

type HttpResponsePayload[T any] struct {
	Message string `json:"message"`
	ErrCode T      `json:"err_code"`
	Erros   any    `json:"errors"`
	Data    any    `json:"data"`
}

func WriteSuccessResponse(w http.ResponseWriter, r *http.Request, code int, msg string, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	payload := HttpResponsePayload[any]{
		Message: msg,
		ErrCode: nil,
		Erros:   nil,
		Data:    data,
	}
	json.NewEncoder(w).Encode(&payload)
}

func WriteErrorResponse(w http.ResponseWriter, r *http.Request, code int, msg string, errCode domain.ERR_CODE, errors any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	payload := HttpResponsePayload[domain.ERR_CODE]{
		Message: msg,
		ErrCode: errCode,
		Erros:   errors,
		Data:    nil,
	}
	json.NewEncoder(w).Encode(&payload)
}
