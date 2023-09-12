package service

import (
	"context"
	"net/http"

	"github.com/ryanadiputraa/flows/flows-microservices/notification/config"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/email"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/pkg/logger"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/pkg/validator"
)

type service struct {
	conf      config.Config
	log       logger.Logger
	validator validator.Validator
}

func NewService(conf config.Config, log logger.Logger, validator validator.Validator) email.Usecase {
	return &service{
		conf:      conf,
		log:       log,
		validator: validator,
	}
}

func (s *service) RegisterNotification(ctx context.Context, dto domain.EmailDTO) error {
	err, errors := s.validator.Validate(dto)
	if err != nil {
		s.log.Warn("register notification: ", err)
		return &domain.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "fail to send notification",
			ErrCode: domain.INVALID_PARAMS,
			Errors:  errors,
		}
	}

	email, err := domain.NewEmail(dto)
	if err != nil {
		s.log.Warn("register notification: ", err)
		return &domain.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "fail to send notification",
			ErrCode: domain.INVALID_PARAMS,
			Errors: map[string][]string{
				"mail_type": {"invalid mail type"},
			},
		}
	}

	s.log.Info(email)

	return nil
}
