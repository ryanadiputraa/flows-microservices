package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
	"github.com/stretchr/testify/assert"
)

func NewMockDB(assert *assert.Assertions) (sqlmock.Sqlmock, *sqlx.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.FailNow("fail to create mock db: ", err)
	}
	mockDB := sqlx.NewDb(db, "sqlmock")
	return mock, mockDB
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	user, _ := domain.NewUser(
		uuid.NewString(), gofakeit.FirstName(), gofakeit.LastName(), gofakeit.Email(),
		gofakeit.ImageURL(120, 120), gofakeit.Password(true, true, true, true, false, 20), domain.IDR)

	cases := []struct {
		name   string
		user   *domain.User
		err    error
		mockDB func(mock sqlmock.Sqlmock)
	}{
		{
			name: "should successfully save new user",
			user: user,
			err:  nil,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("^INSERT INTO users").WithArgs(
					user.ID, user.Email, user.Password, user.FirstName, user.LastName,
					user.Currency, user.Picture, user.CreatedAt,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "should return error when fail to save user",
			user: user,
			err:  errors.New("fail to create user"),
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("^INSERT INTO users").WithArgs(
					user.ID, user.Email, user.Password, user.FirstName, user.LastName,
					user.Currency, user.Picture, user.CreatedAt,
				).WillReturnError(errors.New("fail to create user"))
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mock, db := NewMockDB(assert)
			c.mockDB(mock)
			r := NewRepository(db)

			err := r.Save(context.Background(), user)
			assert.Equal(c.err, err)
		})
	}
}

// func TestFindByID(t *testing.T) {
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
// 		mockDB   func(mock sqlmock.Sqlmock)
// 	}{
// 		{
// 			name:     "should return user with given id",
// 			userID:   uid1,
// 			expected: user1,
// 			err:      nil,
// 			mockDB: func(mock sqlmock.Sqlmock) {
// 				mock.ExpectQuery("SELECT (.+) FROM users").
// 					WillReturnRows(
// 						sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "picture", "currency", "verified_email"}).
// 							AddRow(user1.ID, user1.FirstName, user1.LastName, user1.Email, user1.Picture, user1.Currency, user1.VerifiedEmail),
// 					)
// 			},
// 		},
// 		{
// 			name:     "should return empty user when given user id didn't exists",
// 			userID:   uid2,
// 			expected: &domain.User{},
// 			err:      sql.ErrNoRows,
// 			mockDB: func(mock sqlmock.Sqlmock) {
// 				mock.ExpectQuery("SELECT (.+) FROM users").
// 					WillReturnError(sql.ErrNoRows)
// 			},
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			mock, db := NewMockDB(assert)
// 			c.mockDB(mock)
// 			r := NewRepository(db)

// 			user, err := r.FindByID(context.Background(), c.userID)
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
// 			assert.Equal(c.expected.VerifiedEmail, user.VerifiedEmail)
// 		})
// 	}
// }

// func TestFindByEmail(t *testing.T) {
// 	assert := assert.New(t)

// 	email1 := gofakeit.Email()
// 	user1, _ := domain.NewUser(
// 		uuid.NewString(), gofakeit.FirstName(), email1, gofakeit.Email(),
// 		gofakeit.ImageURL(120, 120), gofakeit.Password(true, true, true, true, false, 20), domain.IDR, true)

// 	email2 := gofakeit.Email()

// 	cases := []struct {
// 		name     string
// 		email    string
// 		expected *domain.User
// 		err      error
// 		mockDB   func(mock sqlmock.Sqlmock)
// 	}{
// 		{
// 			name:     "should return user with given email",
// 			email:    email1,
// 			expected: user1,
// 			err:      nil,
// 			mockDB: func(mock sqlmock.Sqlmock) {
// 				mock.ExpectQuery("SELECT (.+) FROM users").
// 					WillReturnRows(
// 						sqlmock.NewRows([]string{"id", "password"}).
// 							AddRow(user1.ID, user1.Password),
// 					)
// 			},
// 		},
// 		{
// 			name:     "should return empty user when given user email didn't exists",
// 			email:    email2,
// 			expected: &domain.User{},
// 			err:      sql.ErrNoRows,
// 			mockDB: func(mock sqlmock.Sqlmock) {
// 				mock.ExpectQuery("SELECT (.+) FROM users").
// 					WillReturnError(sql.ErrNoRows)
// 			},
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			mock, db := NewMockDB(assert)
// 			c.mockDB(mock)
// 			r := NewRepository(db)

// 			user, err := r.FindByEmail(context.Background(), c.email)
// 			assert.Equal(c.err, err)
// 			if err != nil {
// 				assert.Empty(user)
// 				return
// 			}

// 			assert.Equal(c.expected.ID, user.ID)
// 			assert.Equal(c.expected.Password, user.Password)
// 		})
// 	}
// }
