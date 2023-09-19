package main

import (
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/config"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/server"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/logger"
)

func main() {
	l := logger.NewLogger()

	c, err := config.LoadConfig("config/config")
	if err != nil {
		l.Fatal(err)
	}

	// TODO : create db connection
	s := server.NewServer(c, l, nil)

	if err := s.Run(); err != nil {
		l.Fatal(err)
	}
}
