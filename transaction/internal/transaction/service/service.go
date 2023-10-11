package service

import (
	"context"
	"database/sql"
	"math"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/config"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/transaction"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/logger"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/validator"
)

type service struct {
	config     config.Config
	log        logger.Logger
	validator  validator.Validator
	repository transaction.Repository
}

func NewService(config config.Config, log logger.Logger, validator validator.Validator, repository transaction.Repository) transaction.Usecase {
	return &service{
		config:     config,
		log:        log,
		validator:  validator,
		repository: repository,
	}
}

func (s *service) AddTransaction(ctx context.Context, userID string, dto domain.TransactionDTO) (*domain.Transaction, error) {
	err, errors := s.validator.Validate(dto)
	if err != nil {
		s.log.Warn("add transaction: ", err)
		return nil, &domain.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "fail to add transaction: validation error",
			ErrCode: domain.INVALID_PARAMS,
			Errors:  errors,
		}
	}

	date, _ := time.Parse(time.RFC3339Nano, dto.Date)
	transaction := domain.Transaction{
		ID:          uuid.NewString(),
		UserID:      userID,
		Title:       dto.Title,
		Description: dto.Description,
		Amount:      dto.Amount,
		Date:        date,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if err = s.repository.Save(ctx, transaction); err != nil {
		s.log.Error("add transaction: ", err)
		return nil, &domain.ResponseError{
			Code:    http.StatusInternalServerError,
			Message: "fail to insert transaction",
			ErrCode: domain.INTERNAL_SERVER_ERROR,
			Errors:  nil,
		}
	}

	return &transaction, nil
}

func (s *service) GetTransactionSummary(ctx context.Context, UserID string) (*domain.TransactionSummary, error) {
	c := time.Now().UTC()
	start := time.Date(c.Year(), c.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(c.Year(), c.Month()+1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Microsecond)

	transactions, err := s.repository.List(ctx, UserID, start, end, 5, 1)
	if err != nil {
		s.log.Error(err)
		if err != sql.ErrNoRows {
			return nil, err
		}
	}

	summary := domain.TransactionSummary{
		IncomeInMonth:     0,
		ExpenseInMonth:    0,
		LatestTransaction: transactions,
	}
	for _, t := range transactions {
		if t.Amount > 0 {
			summary.IncomeInMonth += t.Amount
		} else {
			summary.ExpenseInMonth += t.Amount
		}
	}

	return &summary, nil
}

func (s *service) ListTransactions(ctx context.Context, userID, start, end string, size, page int) ([]domain.Transaction, *domain.Meta, error) {
	if start == "" {
		yesterday := time.Now().UTC().Add(-(time.Hour * 24))
		start = yesterday.Format(time.RFC3339Nano)
	}
	if end == "" {
		now := time.Now().UTC()
		end = now.Format(time.RFC3339Nano)
	}
	if size == 0 {
		size = 1
	}
	if page == 0 {
		page = 1
	}

	startDate, err := time.Parse(time.RFC3339Nano, start)
	if err != nil {
		s.log.Warn("invalid date format: ", start)
		return nil, nil, &domain.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "invalid 'start' param, expected ISO861 date string",
			ErrCode: domain.INVALID_PARAMS,
			Errors:  nil,
		}
	}
	endDate, err := time.Parse(time.RFC3339Nano, end)
	if err != nil {
		s.log.Warn("invalid date format: ", end)
		return nil, nil, &domain.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "invalid 'end' param, expected ISO861 date string",
			ErrCode: domain.INVALID_PARAMS,
			Errors:  nil,
		}
	}

	transactions, err := s.repository.List(ctx, userID, startDate, endDate, size, page)
	if err != nil {
		s.log.Error("fail to retrieve transactions: ", err)
		if err != sql.sql.ErrNoRows {
			return nil, nil, err
		}
	}

	total := 0
	totalPages := 0
	if len(transactions) > 0 {
		s.log.Info("hittt")
		total = transactions[0].TotalData
		totalPages = int(math.Ceil(float64(total) / float64(size)))
	}

	return transactions,
		&domain.Meta{
			Size:        size,
			Total:       total,
			TotalPages:  totalPages,
			CurrentPage: page},
		nil
}
