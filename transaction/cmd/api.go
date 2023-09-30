package main

import (
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/config"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/server"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/db"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/logger"
)

func main() {
	log := logger.NewLogger()

	conf, err := config.LoadConfig("config/config")
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewDB(conf)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(conf, log, db)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
