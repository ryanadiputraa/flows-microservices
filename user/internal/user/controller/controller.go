package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/middleware"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/jwt"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/logger"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/notification"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/response"
)

type controller struct {
	handler             *http.ServeMux
	log                 logger.Logger
	service             user.Usecase
	jwtService          jwt.JWTService
	notificationService notification.NotificationService
}

func NewController(
	handler *http.ServeMux, log logger.Logger, service user.Usecase, authMiddleware *middleware.AuthMiddleware,
	jwtService jwt.JWTService, notificationService notification.NotificationService) {
	c := &controller{
		handler:             handler,
		log:                 log,
		service:             service,
		jwtService:          jwtService,
		notificationService: notificationService,
	}

	handler.HandleFunc("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			c.Register(w, r)
		}
	})
	handler.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			c.Login(w, r)
		}
	})
	handler.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			authMiddleware.ClaimJWTToken(c.GetUserInfo)(w, r)
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

	go func() {
		err := c.notificationService.SendMail(notification.MailPayload{
			Subject:   "Welcome to Flows",
			To:        data.Email,
			FirstName: data.FirstName,
			MailType:  "register",
		})
		if err != nil {
			c.log.Error("notification service: ", err)
		}
	}()

	response.WriteSuccessResponse(w, r, http.StatusOK, "user successfully register", &data)
	return
}

func (c *controller) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var dto domain.LoginDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.WriteErrorResponse(w, r, http.StatusBadRequest, "empty request body", response.INVALID_PARAMS, nil)
		return
	}

	user, err := c.service.Login(ctx, &dto)
	if err != nil {
		if rErr, ok := err.(*domain.ResponseError); ok {
			response.WriteErrorResponse(w, r, rErr.Code, rErr.Message, rErr.ErrCode, rErr.Errors)
			return
		} else {
			response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", response.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	tokens, err := c.jwtService.GenerateJWTTokens(ctx, user.ID)
	if err != nil {
		response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", response.INTERNAL_SERVER_ERROR, nil)
		return
	}

	response.WriteSuccessResponse(w, r, http.StatusOK, "jwt tokens generated", &tokens)
	return
}

func (c *controller) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := fmt.Sprintf("%v", r.Context().Value("userID"))

	user, err := c.service.GetUserInfo(ctx, userID)
	if err != nil {
		if rErr, ok := err.(*domain.ResponseError); ok {
			response.WriteErrorResponse(w, r, rErr.Code, rErr.Message, rErr.ErrCode, rErr.Errors)
			return
		} else {
			response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", response.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	response.WriteSuccessResponse(w, r, http.StatusOK, "fetch user info", user)
}
