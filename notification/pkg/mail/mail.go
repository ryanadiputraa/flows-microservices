package mail

import (
	"fmt"
	"net/smtp"

	"github.com/ryanadiputraa/flows/flows-microservices/notification/config"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/domain"
)

type SmptMail struct{}

func NewSmptMail() *SmptMail {
	return &SmptMail{}
}

func (s *SmptMail) SendMail(conf config.Mail, payload *domain.Email) (err error) {
	auth := smtp.PlainAuth("", conf.Sender, conf.Pass, "smtp.gmail.com")

	msg := fmt.Sprintf("Subject: %s\n%s", payload.Subject, payload.PlainTextContent)

	err = smtp.SendMail("smtp.gmail.com:587", auth, conf.Sender, []string{payload.To}, []byte(msg))
	return
}
