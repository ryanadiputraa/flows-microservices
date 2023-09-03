package server

import (
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user/controller"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user/repository"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user/service"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/validator"
)

func (s *Server) MapHandlers() {
	validator := validator.NewValidator()

	userRepository := repository.NewRepository(s.DB, s.Config.DB.DB_Name)
	userService := service.NewService(*s.Config, validator, s.Logger, userRepository)
	controller.NewController(s.Handler, userService)
}
