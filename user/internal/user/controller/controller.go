package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/response"
)

type controller struct {
	handler *http.ServeMux
	service user.Usecase
}

func NewController(handler *http.ServeMux, service user.Usecase) {
	c := &controller{
		handler: handler,
		service: service,
	}

	handler.HandleFunc("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			c.Register(w, r)
		}
	})
}

func (c *controller) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user domain.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.WriteErrorResponse(w, r, http.StatusBadRequest, "empty request body", response.INVALID_PARAMS, nil)
		return
	}

	data, err := c.service.Register(ctx, &user)
	if err != nil {
		if rErr, ok := err.(*domain.ResponseError); ok {
			response.WriteErrorResponse(w, r, rErr.Code, rErr.Message, rErr.ErrCode, rErr.Errors)
			return
		} else {
			response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", response.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	response.WriteSuccessResponse(w, r, http.StatusOK, "user successfully register", &data)
	return
}
