package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/email"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/pkg/response"
)

type controller struct {
	handler *http.ServeMux
	service email.Usecase
}

func NewController(handler *http.ServeMux, service email.Usecase) {
	c := &controller{
		handler: handler,
		service: service,
	}

	handler.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			c.RegisterNotification(w, r)
		}
	})

}

func (c *controller) RegisterNotification(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var dto domain.EmailDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.WriteErrorResponse(w, r, http.StatusBadRequest, "empty request body", domain.INVALID_PARAMS, nil)
		return
	}

	err := c.service.RegisterNotification(ctx, dto)
	if err != nil {
		if rErr, ok := err.(*domain.ResponseError); ok {
			response.WriteErrorResponse(w, r, rErr.Code, rErr.Message, rErr.ErrCode, rErr.Errors)
			return
		} else {
			response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", domain.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	response.WriteSuccessResponse(w, r, http.StatusOK, "notification sent", nil)
	return
}
