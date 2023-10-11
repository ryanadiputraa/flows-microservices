package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/config"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/logger"
)

type Server struct {
	Config  *config.Config
	Handler *http.ServeMux
	Logger  logger.Logger
	DB      *sqlx.DB
}

func NewServer(config *config.Config, logger logger.Logger, DB *sqlx.DB) *Server {
	return &Server{
		Config:  config,
		Handler: http.NewServeMux(),
		Logger:  logger,
		DB:      DB,
	}
}

func (s *Server) Run() error {
	s.mapHandler()

	server := &http.Server{
		Addr:         s.Config.Server.Port,
		Handler:      s.Handler,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}

	go func() {
		s.Logger.Info("http server running on port", s.Config.Server.Port)
		if err := server.ListenAndServe(); err != nil {
			s.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)

	sig := <-quit
	s.Logger.Warn("received terminate, graceful shutdown ", sig)

	tc, shutdown := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdown()

	return server.Shutdown(tc)
}
