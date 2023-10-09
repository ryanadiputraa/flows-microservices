package service

import (
	"context"
	"database/sql"
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
	end := time.Date(c.Year(), c.Month()+1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Second)

	transactions, err := s.repository.List(ctx, UserID, start, end, 5, 1)
	if err != nil {
		s.log.Error(err)
		if err != sql.ErrNoRows {
			return nil, &domain.ResponseError{
				Code:    http.StatusInternalServerError,
				Message: "transaction summary: fail to retrieve transactions",
				ErrCode: domain.INTERNAL_SERVER_ERROR,
				Errors:  nil,
			}
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
