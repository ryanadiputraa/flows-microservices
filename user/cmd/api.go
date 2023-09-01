package main

import (
	"log"

	"github.com/ryanadiputraa/flows/flows-microservices/user/config"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/server"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/logger"
)

func main() {
	c, err := config.LoadConfig("config/config")
	if err != nil {
		log.Fatal(err)
	}

	l := logger.NewLogger()

	s := server.NewServer(c, l)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
