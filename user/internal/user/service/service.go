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
	"go.mongodb.org/mongo-driver/mongo"
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
		if domain.IsDuplicateSQLError(err) {
			s.log.Warn("register user: ", err)
			return nil, &domain.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "fail to register user",
				ErrCode: response.INVALID_PARAMS,
				Errors: map[string][]string{
					"email": {"email already been register"},
				},
			}
		}
		s.log.Error("register user: ", err)
		return nil, err
	}
	s.log.Info("new user registered: ", user.ID)

	return user, nil
}

func (s *service) Login(ctx context.Context, dto *domain.LoginDTO) (*domain.User, error) {
	user, err := s.repository.FindByEmail(ctx, dto.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Warn("login user: ", err)
			return nil, &domain.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "fail to sign in user",
				ErrCode: response.INVALID_PARAMS,
				Errors: map[string][]string{
					"email": {"no user found with given email"},
				},
			}
		}
		s.log.Error("login user: ", err)
		return nil, err
	}

	if err := user.CompareHashedPassword(dto.Password); err != nil {
		s.log.Warn("login user: ", err)
		return nil, &domain.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "fail to sign in user",
			ErrCode: response.INVALID_PARAMS,
			Errors: map[string][]string{
				"password": {"password didn't match"},
			},
		}
	}

	return user, nil
}

func (s *service) GetUserInfo(ctx context.Context, userID string) (*domain.User, error) {
	user, err := s.repository.FindByID(ctx, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Warn("user info: ", err)
			return nil, &domain.ResponseError{
				Message: "missing user data",
				Code:    http.StatusBadRequest,
				ErrCode: response.INVALID_PARAMS,
				Errors:  nil,
			}
		}
		s.log.Error("user info: ", err)
		return nil, err
	}

	return user, nil
}
