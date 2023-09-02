package main

import (
	"github.com/ryanadiputraa/flows/flows-microservices/user/config"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/server"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/db"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/logger"
)

func main() {
	l := logger.NewLogger()

	c, err := config.LoadConfig("config/config")
	if err != nil {
		l.Fatal(err)
	}

	db, err := db.NewDB(*c.DB)
	if err != nil {
		l.Fatal(err)
	}

	s := server.NewServer(c, l, db)

	if err := s.Run(); err != nil {
		l.Fatal(err)
	}
}
