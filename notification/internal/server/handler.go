package server

import (
	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/email/controller"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/email/service"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/pkg/mail"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/pkg/validator"
)

func (s *Server) mapHandler() {
	validator := validator.NewValidator()
	smtp := mail.NewSmptMail()

	serivce := service.NewService(*s.Config, s.Logger, validator, *smtp)
	controller.NewController(s.Handler, serivce)
}
