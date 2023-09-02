package domain

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Currency string

const (
	IDR Currency = "IDR"
	USD Currency = "USD"
	GBP Currency = "GBP"
	EUR Currency = "EUR"
	YEN Currency = "YEN"
)

type User struct {
	ID        string    `json:"id" db:"id" validate:"required,max=100"`
	Email     string    `json:"email" db:"email" validate:"required,email,max=100"`
	Password  string    `json:"-" db:"password" validate:"required,max=100"`
	FirstName string    `json:"first_name" db:"first_name" validate:"required,max=100"`
	LastName  string    `json:"last_name" db:"last_name" validate:"required,max=100"`
	Currency  Currency  `json:"currency" db:"currency" validate:"required,max=20"`
	Picture   string    `json:"picture" db:"picture" validate:"url_encoded"`
	CreatedAt time.Time `json:"-" db:"created_at"`
}

type UserDTO struct {
	Email     string   `json:"email" db:"email" validate:"required,email,max=100"`
	Password  string   `json:"password" db:"password" validate:"required,min=8"`
	FirstName string   `json:"first_name" db:"first_name" validate:"required,max=100"`
	LastName  string   `json:"last_name" db:"last_name" validate:"required,max=100"`
	Currency  Currency `json:"currency" db:"currency" validate:"required,max=20"`
	Picture   string   `json:"picture" db:"picture" validate:"http_url"`
}

func NewUser(id, firstName, lastName, email, picture, password string, currency Currency) (*User, error) {
	if !IsValidCurrency(currency) {
		return &User{}, errors.New("invalid currency")
	}

	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Currency:  currency,
		Picture:   picture,
		Password:  password,
		CreatedAt: time.Now(),
	}, nil
}

func IsValidCurrency(currency Currency) bool {
	return currency == IDR || currency == USD || currency == GBP || currency == EUR || currency == YEN
}

func (u *User) HashPassowrd() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return nil
}

func (u *User) CompareHashedPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}