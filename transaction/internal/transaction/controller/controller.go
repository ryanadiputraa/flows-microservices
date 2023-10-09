package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/transaction"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/jwt"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/response"
)

type controller struct {
	service    transaction.Usecase
	jwtService jwt.JWTService
}

func NewController(handler *http.ServeMux, service transaction.Usecase, jwtService jwt.JWTService) {
	c := &controller{
		service:    service,
		jwtService: jwtService,
	}

	handler.HandleFunc("/api/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			c.GetTransactionSummary(w, r)
		case http.MethodPost:
			c.AddTransaction(w, r)
		}
	})
}

func (c *controller) AddTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token, err := c.jwtService.ExtractJWTTokenHeader(r.Header)
	if err != nil {
		if rErr, ok := err.(*domain.ResponseError); ok {
			response.WriteErrorResponse(w, r, rErr.Code, rErr.Message, rErr.ErrCode, rErr.Errors)
			return
		} else {
			response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", domain.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	claims, err := c.jwtService.ParseJWTClaims(ctx, token)
	if err != nil {
		if rErr, ok := err.(*domain.ResponseError); ok {
			response.WriteErrorResponse(w, r, rErr.Code, rErr.Message, rErr.ErrCode, rErr.Errors)
			return
		} else {
			response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", domain.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	var transactionDTO domain.TransactionDTO
	if err := json.NewDecoder(r.Body).Decode(&transactionDTO); err != nil {
		response.WriteErrorResponse(w, r, http.StatusBadRequest, "empty request body", domain.INVALID_PARAMS, nil)
		return
	}

	transaction, err := c.service.AddTransaction(ctx, claims.UserID, transactionDTO)
	if err != nil {
		if rErr, ok := err.(*domain.ResponseError); ok {
			response.WriteErrorResponse(w, r, rErr.Code, rErr.Message, rErr.ErrCode, rErr.Errors)
			return
		} else {
			response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", domain.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	response.WriteSuccessResponse(w, r, http.StatusCreated, "transaction added", transaction)
}

func (c *controller) GetTransactionSummary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token, err := c.jwtService.ExtractJWTTokenHeader(r.Header)
	if err != nil {
		if rErr, ok := err.(*domain.ResponseError); ok {
			response.WriteErrorResponse(w, r, rErr.Code, rErr.Message, rErr.ErrCode, rErr.Errors)
			return
		} else {
			response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", domain.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	claims, err := c.jwtService.ParseJWTClaims(ctx, token)
	if err != nil {
		if rErr, ok := err.(*domain.ResponseError); ok {
			response.WriteErrorResponse(w, r, rErr.Code, rErr.Message, rErr.ErrCode, rErr.Errors)
			return
		} else {
			response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", domain.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	transactions, err := c.service.GetTransactionSummary(ctx, claims.UserID)
	if err != nil {
		if rErr, ok := err.(*domain.ResponseError); ok {
			response.WriteErrorResponse(w, r, rErr.Code, rErr.Message, rErr.ErrCode, rErr.Errors)
			return
		} else {
			response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", domain.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	response.WriteSuccessResponse(w, r, http.StatusCreated, "get transactions summary", transactions)
}
