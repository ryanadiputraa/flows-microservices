package server

import (
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/transaction/controller"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/transaction/repository"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/transaction/service"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/validator"
)

func (s *Server) mapHandler() {
	validator := validator.NewValidator()
	repository := repository.NewRepository(s.DB)
	service := service.NewService(*s.Config, s.Logger, validator, repository)
	controller.NewController(s.Handler, service)
}
