package domain

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	cases := []struct {
		name     string
		currency Currency
		err      error
	}{
		{
			name:     "should return new user with valid currency",
			currency: IDR,
			err:      nil,
		},
		{
			name:     "should return empty user and return invalid currency error",
			currency: "bitcoin",
			err:      errors.New("invalid currency"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			id := uuid.NewString()
			firstName := gofakeit.FirstName()
			lastName := gofakeit.LastName()
			email := gofakeit.Email()
			picture := gofakeit.ImageURL(120, 120)
			verified := true
			password := gofakeit.Password(true, true, true, true, false, 20)

			user, err := NewUser(id, firstName, lastName, email, picture, password, c.currency, verified)

			assert.Equal(t, c.err, err)
			if err != nil {
				assert.Empty(t, user)
				return
			}
			assert.NotEmpty(t, user)
			assert.Equal(t, id, user.ID)
			assert.Equal(t, firstName, user.FirstName)
			assert.Equal(t, lastName, user.LastName)
			assert.Equal(t, email, user.Email)
			assert.Equal(t, picture, user.Picture)
			assert.Equal(t, c.currency, user.Currency)
			assert.Equal(t, password, user.Password)
		})
	}
}

func TestPasswordHash(t *testing.T) {
	cases := []struct {
		name string
	}{
		{
			name: "should hash user password",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			id := uuid.NewString()
			firstName := gofakeit.FirstName()
			lastName := gofakeit.LastName()
			email := gofakeit.Email()
			picture := gofakeit.ImageURL(120, 120)
			currency := IDR
			verified := true
			password := gofakeit.Password(true, true, true, true, false, 20)

			user, err := NewUser(id, firstName, lastName, email, picture, password, currency, verified)
			user.HashPassowrd()
			assert.NoError(t, err)

			assert.NotEqual(t, password, user.Password)
			assert.NoError(t, user.CompareHashedPassword(password))
		})
	}
}
