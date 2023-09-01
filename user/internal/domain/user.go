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
	ID        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Currency  Currency  `json:"currency" db:"currency"`
	Picture   string    `json:"picture" db:"picture"`
	CreatedAt time.Time `json:"-" db:"created_at"`
}

func NewUser(id, firstName, lastName, email, picture, password string, currency Currency, verified bool) (*User, error) {
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
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return nil
}

func (u *User) CompareHashedPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
