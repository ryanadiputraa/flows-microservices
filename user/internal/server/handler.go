package server

import (
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user/controller"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user/repository"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user/service"

	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/jwt"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/notification"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/validator"
)

func (s *Server) MapHandlers() {
	validator := validator.NewValidator()

	jwtService := jwt.NewService(s.Logger)
	notificationService, err := notification.NewNotificationService(*s.Config)
	if err != nil {
		s.Logger.Fatal(err)
	}

	userRepository := repository.NewRepository(s.DB, s.Config.DB.DB_Name)
	userService := service.NewService(*s.Config, validator, s.Logger, userRepository)
	controller.NewController(s.Handler, s.Logger, userService, jwtService, *notificationService)
}
