package domain

import (
	"errors"
	"fmt"
)

type MailType string

const (
	Sender   string   = "ryannadiputraa@gmail.com"
	Register MailType = "register"
)

var (
	mailContents = map[string]string{
		string(Register): "Hi %v, thank you for register on Flows!\nFlows is an open source finance tracking application, feel free to contribute.",
	}

	htmlContents = map[string]string{
		string(Register): `
			<h1>Hi %v,</h1><br/>
			<h2>Thank you for register on Flows!</h2><br/><hr/>
			<span>Flows is an open source finance tracking application, feel free to contribute.</span>`,
	}
)

type Email struct {
	From             string
	Subject          string
	To               string
	PlainTextContent string
	HTMLContent      string
}

type EmailDTO struct {
	Subject   string   `json:"subject" validate:"required,max=100"`
	To        string   `json:"to" validate:"required,email"`
	FirstName string   `json:"first_name" validate:"required"`
	MailType  MailType `json:"mail_type" validate:"required"`
}

func NewEmail(dto EmailDTO) (*Email, error) {
	if !IsValidaMailType(dto.MailType) {
		return nil, errors.New("invalid mail type")
	}

	plainTextContent := mailContents[string(dto.MailType)]
	htmlContent := htmlContents[string(dto.MailType)]

	return &Email{
		From:             Sender,
		Subject:          dto.Subject,
		To:               dto.To,
		PlainTextContent: fmt.Sprintf(plainTextContent, dto.FirstName),
		HTMLContent:      fmt.Sprintf(htmlContent, dto.FirstName),
	}, nil
}

func IsValidaMailType(t MailType) bool {
	return t == Register
}
