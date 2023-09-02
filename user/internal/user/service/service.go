package service

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ryanadiputraa/flows/flows-microservices/user/config"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/logger"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/response"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/validator"
)

type service struct {
	config     *config.Config
	validator  validator.Validator
	log        logger.Logger
	repository user.Repository
}

func NewService(config config.Config, validator validator.Validator, log logger.Logger, repository user.Repository) user.Usecase {
	return &service{
		config:     &config,
		validator:  validator,
		log:        log,
		repository: repository,
	}
}

func (s *service) Register(ctx context.Context, dto *domain.UserDTO) (*domain.User, error) {
	err, errors := s.validator.Validate(dto)
	if err != nil {
		s.log.Warn("register user: ", err)
		return nil, &domain.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "fail to register user",
			ErrCode: response.INVALID_PARAMS,
			Errors:  errors,
		}
	}

	user, err := domain.NewUser(
		uuid.NewString(),
		dto.FirstName, dto.LastName, dto.Email,
		dto.Picture, dto.Password, dto.Currency,
	)
	if err != nil {
		s.log.Warn("register user: ", err)
		return nil, &domain.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "fail to register user",
			ErrCode: response.INVALID_PARAMS,
			Errors: map[string][]string{
				"currency": {"currency is not a valid currency"},
			},
		}
	}

	if err := user.HashPassowrd(); err != nil {
		s.log.Error("register user: ", err)
		return nil, err
	}

	if err := s.repository.Save(ctx, user); err != nil {
		s.log.Warn("register user: ", err)
		if domain.IsDuplicateSQLError(err) {
			return nil, &domain.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "fail to register user",
				ErrCode: response.INVALID_PARAMS,
				Errors: map[string][]string{
					"email": {"email already been register"},
				},
			}
		}
		return nil, err
	}
	s.log.Info("new user registered: ", user.ID)

	return user, nil
}
