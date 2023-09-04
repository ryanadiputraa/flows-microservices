package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ryanadiputraa/flows/flows-microservices/user/config"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/mocks"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/response"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/validator"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	testConfig    = config.Config{}
	testLog       = logrus.New()
	testValidator = validator.NewValidator()
)

func TestRegister(t *testing.T) {
	dto := &domain.UserDTO{
		Email:     gofakeit.Email(),
		Password:  gofakeit.Password(true, true, true, true, false, 20),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Currency:  "IDR",
		Picture:   gofakeit.ImageURL(80, 80),
	}
	invalidDTO := &domain.UserDTO{
		Email:     "email",
		Password:  "",
		FirstName: "",
		LastName:  "",
		Currency:  "IDR",
		Picture:   "picture",
	}
	invalidDTOCurrency := &domain.UserDTO{
		Email:     gofakeit.Email(),
		Password:  gofakeit.Password(true, true, true, true, false, 20),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Currency:  "id",
		Picture:   gofakeit.ImageURL(80, 80),
	}

	user, _ := domain.NewUser(
		gofakeit.UUID(), dto.FirstName, dto.LastName, dto.Email,
		dto.Picture, dto.Password, dto.Currency)
	user.HashPassowrd()

	cases := []struct {
		name     string
		dto      *domain.UserDTO
		expected *domain.User
		err      error
		mockRepo func(mockRepo *mocks.Repository)
	}{
		{
			name:     "should register user",
			dto:      dto,
			expected: user,
			err:      nil,
			mockRepo: func(mockRepo *mocks.Repository) {
				mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name:     "should fail to register user",
			dto:      dto,
			expected: &domain.User{},
			err:      sql.ErrConnDone,
			mockRepo: func(mockRepo *mocks.Repository) {
				mockRepo.On("Save", mock.Anything, mock.Anything).Return(sql.ErrConnDone)
			},
		},
		{
			name:     "should fail to register user when fail to validate dto",
			dto:      invalidDTO,
			expected: nil,
			err: &domain.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "fail to register user",
				ErrCode: response.INVALID_PARAMS,
				Errors: map[string][]string{
					"email":     {"email should be a valid email address"},
					"password":  {"password is required"},
					"firstname": {"firstname is required"},
					"lastname":  {"lastname is required"},
					"picture":   {"picture should be a valid http url"},
				}},
			mockRepo: func(mockRepo *mocks.Repository) {
				mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name:     "should fail to register user when given currency is invalid",
			dto:      invalidDTOCurrency,
			expected: nil,
			err: &domain.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "fail to register user",
				ErrCode: response.INVALID_PARAMS,
				Errors: map[string][]string{
					"currency": {"currency is not a valid currency"},
				}},
			mockRepo: func(mockRepo *mocks.Repository) {
				mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name:     "should fail to register user when email been register",
			dto:      dto,
			expected: &domain.User{},
			err: &domain.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "fail to register user",
				ErrCode: response.INVALID_PARAMS,
				Errors: map[string][]string{
					"email": {"email already been register"},
				}},
			mockRepo: func(mockRepo *mocks.Repository) {
				mockRepo.On("Save", mock.Anything, mock.Anything).Return(errors.New("sql: duplicate fk"))
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := new(mocks.Repository)
			c.mockRepo(r)
			s := NewService(testConfig, testValidator, testLog, r)

			d, err := s.Register(context.Background(), c.dto)
			assert.Equal(t, c.err, err)
			if err != nil {
				assert.Empty(t, d)
				if domain.IsDuplicateSQLError(err) {
					assert.Equal(t, c.expected, err)
				}
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, d.ID)
			assert.Equal(t, c.expected.FirstName, d.FirstName)
			assert.Equal(t, c.expected.LastName, d.LastName)
			assert.Equal(t, c.expected.Email, d.Email)
			assert.Equal(t, c.expected.Picture, d.Picture)
			assert.Equal(t, c.expected.Currency, d.Currency)
			assert.NotEqual(t, dto.Password, d.Password)
		})
	}
}

func TestLogin(t *testing.T) {
	user, _ := domain.NewUser(
		gofakeit.UUID(), gofakeit.FirstName(), gofakeit.LastName(),
		"test@mail.com", gofakeit.ImageURL(80, 80),
		"testpassword", "IDR")
	user.HashPassowrd()

	cases := []struct {
		name     string
		dto      *domain.LoginDTO
		expected *domain.User
		err      error
		mockRepo func(mockRepo *mocks.Repository)
	}{
		{
			name: "should successfully return logged in user",
			dto: &domain.LoginDTO{
				Email:    user.Email,
				Password: "testpassword",
			},
			expected: user,
			err:      nil,
			mockRepo: func(mockRepo *mocks.Repository) {
				mockRepo.On("FindByEmail", mock.Anything, user.Email).Return(user, nil)
			},
		},
		{
			name: "should fail when user with given email didn't exists",
			dto: &domain.LoginDTO{
				Email:    "rand@mail.com",
				Password: "testpassword",
			},
			expected: nil,
			err: &domain.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "fail to sign in user",
				ErrCode: response.INVALID_PARAMS,
				Errors: map[string][]string{
					"email": {"no user found with given email"},
				},
			},
			mockRepo: func(mockRepo *mocks.Repository) {
				mockRepo.On("FindByEmail", mock.Anything, "rand@mail.com").Return(nil, mongo.ErrNoDocuments)
			},
		},
		{
			name: "should fail when repository fail to fetch",
			dto: &domain.LoginDTO{
				Email:    user.Email,
				Password: "testpassword",
			},
			expected: nil,
			err:      mongo.ErrClientDisconnected,
			mockRepo: func(mockRepo *mocks.Repository) {
				mockRepo.On("FindByEmail", mock.Anything, user.Email).Return(nil, mongo.ErrClientDisconnected)
			},
		},
		{
			name: "should fail when password didn't match",
			dto: &domain.LoginDTO{
				Email:    "rand@mail.com",
				Password: "falsepassword",
			},
			expected: nil,
			err: &domain.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "fail to sign in user",
				ErrCode: response.INVALID_PARAMS,
				Errors: map[string][]string{
					"password": {"password didn't match"},
				},
			},
			mockRepo: func(mockRepo *mocks.Repository) {
				mockRepo.On("FindByEmail", mock.Anything, "rand@mail.com").Return(user, nil)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := new(mocks.Repository)
			c.mockRepo(r)
			s := NewService(testConfig, testValidator, testLog, r)

			d, err := s.Login(context.TODO(), c.dto)
			if err != nil {
				assert.Empty(t, d)
				assert.Equal(t, c.err, err)
				return
			}

			assert.Empty(t, err)
			assert.Equal(t, c.expected, d)
		})
	}
}
