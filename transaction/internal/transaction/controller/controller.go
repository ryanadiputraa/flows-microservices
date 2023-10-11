package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/middleware"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/transaction"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/response"
	"github.com/sirupsen/logrus"
)

type controller struct {
	service transaction.Usecase
}

func NewController(handler *http.ServeMux, service transaction.Usecase, authMiddleware *middleware.AuthMiddleware) {
	c := &controller{
		service: service,
	}

	handler.HandleFunc("/api/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			authMiddleware.ClaimJWTToken(c.ListTransactions)(w, r)
		case http.MethodPost:
			authMiddleware.ClaimJWTToken(c.AddTransaction)(w, r)
		}
	})
	handler.HandleFunc("/api/transactions/summary", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			authMiddleware.ClaimJWTToken(c.GetTransactionSummary)(w, r)
		}
	})
}

func (c *controller) AddTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := fmt.Sprintf("%v", r.Context().Value("userID"))

	var transactionDTO domain.TransactionDTO
	if err := json.NewDecoder(r.Body).Decode(&transactionDTO); err != nil {
		response.WriteErrorResponse(w, r, http.StatusBadRequest, "empty request body", domain.INVALID_PARAMS, nil)
		return
	}

	transaction, err := c.service.AddTransaction(ctx, userID, transactionDTO)
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
	userID := fmt.Sprintf("%v", r.Context().Value("userID"))

	transactions, err := c.service.GetTransactionSummary(ctx, userID)
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

func (c *controller) ListTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := fmt.Sprintf("%v", r.Context().Value("userID"))
	q := r.URL.Query()
	start := q.Get("start")
	end := q.Get("end")
	size, _ := strconv.Atoi(q.Get("size"))
	page, _ := strconv.Atoi(q.Get("page"))

	transactions, meta, err := c.service.ListTransactions(ctx, userID, start, end, size, page)
	logrus.Info(meta)
	if err != nil {
		if rErr, ok := err.(*domain.ResponseError); ok {
			response.WriteErrorResponse(w, r, rErr.Code, rErr.Message, rErr.ErrCode, rErr.Errors)
			return
		} else {
			response.WriteErrorResponse(w, r, http.StatusInternalServerError, "internal server error", domain.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	response.WriteSuccessResponse(w, r, http.StatusCreated, "get list transactions", transactions)
}
