package main

import (
	"github.com/ryanadiputraa/flows/flows-microservices/notification/config"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/server"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/pkg/logger"
)

func main() {
	l := logger.NewLogger()

	c, err := config.LoadConfig("config/config")
	if err != nil {
		l.Fatal(err)
	}

	s := server.NewServer(c, l)

	if err := s.Run(); err != nil {
		l.Fatal(err)
	}
}
