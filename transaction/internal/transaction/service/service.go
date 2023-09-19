package service

import (
	"context"

	"github.com/ryanadiputraa/flows/flows-microservices/transaction/config"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/transaction"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/logger"
)

type service struct {
	config     config.Config
	log        logger.Logger
	repository transaction.Repository
}

func NewService(config config.Config, log logger.Logger, repository transaction.Repository) transaction.Usecase {
	return &service{
		config:     config,
		log:        log,
		repository: repository,
	}
}

func (s *service) AddTransaction(ctx context.Context, userID string, dto *domain.TransactionDTO) (*domain.Transaction, error) {
	return nil, nil
}

func (s *service) GetTransactionSummary(ctx context.Context, UserID string) (*domain.TransactionSummary, error) {
	return nil, nil
}
