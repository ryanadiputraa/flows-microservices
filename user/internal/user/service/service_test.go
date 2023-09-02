package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ryanadiputraa/flows/flows-microservices/user/config"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/mocks"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/response"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/validator"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		uuid.NewString(), dto.FirstName, dto.LastName, dto.Email,
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

// func TestGetUserInfo(t *testing.T) {
// 	assert := assert.New(t)

// 	uid1 := uuid.NewString()
// 	user1, _ := domain.NewUser(
// 		uid1, gofakeit.FirstName(), gofakeit.LastName(), gofakeit.Email(),
// 		gofakeit.ImageURL(120, 120), gofakeit.Password(true, true, true, true, false, 20), domain.IDR, true)

// 	uid2 := uuid.NewString()

// 	cases := []struct {
// 		name     string
// 		userID   string
// 		expected *domain.User
// 		err      error
// 		mockRepo func(mockRepo *mocks.Repository)
// 	}{
// 		{
// 			name:     "should return user with given id",
// 			userID:   uid1,
// 			expected: user1,
// 			err:      nil,
// 			mockRepo: func(mockRepo *mocks.Repository) {
// 				mockRepo.On("FindByID", mock.Anything, uid1).Return(user1, nil)
// 			},
// 		},
// 		{
// 			name:     "should return empty user when given user id didn't exists",
// 			userID:   uid2,
// 			expected: &domain.User{},
// 			err:      sql.ErrNoRows,
// 			mockRepo: func(mockRepo *mocks.Repository) {
// 				mockRepo.On("FindByID", mock.Anything, uid2).Return(&domain.User{}, sql.ErrNoRows)
// 			},
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			r := new(mocks.Repository)
// 			c.mockRepo(r)
// 			u := NewUseCase(r)

// 			user, err := u.GetUserData(context.Background(), c.userID)
// 			assert.Equal(c.err, err)
// 			if err != nil {
// 				assert.Empty(user)
// 				return
// 			}

// 			assert.Equal(c.expected.ID, user.ID)
// 			assert.Equal(c.expected.FirstName, user.FirstName)
// 			assert.Equal(c.expected.LastName, user.LastName)
// 			assert.Equal(c.expected.Email, user.Email)
// 			assert.Equal(c.expected.Picture, user.Picture)
// 			assert.Equal(c.expected.Currency, user.Currency)
// 			assert.Equal(c.expected.VerifiedEmail, false)
// 		})
// 	}
// }

// func TestSignIn(t *testing.T) {
// 	user, _ := domain.NewUser(
// 		uuid.NewString(), gofakeit.FirstName(), gofakeit.LastName(), gofakeit.Email(),
// 		gofakeit.ImageURL(120, 120), gofakeit.Password(true, true, true, true, false, 20), domain.IDR, true)

// 	mockUser, _ := domain.NewUser(
// 		user.ID, user.FirstName, user.LastName, user.Email, user.Picture,
// 		user.Password, user.Currency, false)
// 	mockUser.HashPassowrd()

// 	cases := []struct {
// 		name     string
// 		user     *domain.User
// 		err      error
// 		mockRepo func(mockRepo *mocks.Repository)
// 	}{
// 		{
// 			name: "should found user with given email",
// 			user: user,
// 			err:  nil,
// 			mockRepo: func(mockRepo *mocks.Repository) {
// 				mockRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(mockUser, nil)
// 			},
// 		},
// 		{
// 			name: "should fail to retrieve user",
// 			user: user,
// 			err:  sql.ErrNoRows,
// 			mockRepo: func(mockRepo *mocks.Repository) {
// 				mockRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(&domain.User{}, sql.ErrNoRows)
// 			},
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			r := new(mocks.Repository)
// 			c.mockRepo(r)
// 			u := NewUseCase(r)

// 			userID, err := u.SignIn(context.Background(), user.Email, user.Password)
// 			assert.Equal(t, c.err, err)
// 			if err != nil {
// 				assert.Empty(t, userID)
// 				return
// 			}
// 			assert.Equal(t, user.ID, userID)
// 		})
// 	}
// }
