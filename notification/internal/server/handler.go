package server

import (
	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/email/service"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/pkg/mail"
	messagebroker "github.com/ryanadiputraa/flows/flows-microservices/notification/pkg/message-broker"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/pkg/validator"
)

func (s *Server) mapHandler() {
	validator := validator.NewValidator()
	smtp := mail.NewSmptMail()

	serivce := service.NewService(*s.Config, s.Logger, validator, *smtp)
	messageBroker, err := messagebroker.NewMessageBrokerConsumer(*s.Config, s.Logger, serivce)
	if err != nil {
		s.Logger.Fatal(err)
	}
	if err = messageBroker.Listen(); err != nil {
		s.Logger.Fatal(err)
	}
}
