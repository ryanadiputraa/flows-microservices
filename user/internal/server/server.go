package server

import (
	"net/http"

	"github.com/ryanadiputraa/flows/flows-microservices/user/config"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/logger"
)

type Server struct {
	Config  *config.Config
	Handler *http.ServeMux
	Logger  logger.Logger
}

func NewServer(config *config.Config, logger logger.Logger) *Server {
	return &Server{
		Config:  config,
		Handler: http.NewServeMux(),
		Logger:  logger,
	}
}

func (s *Server) Run() error {
	s.Logger.Info("http server running on port", s.Config.Server.Port)

	server := &http.Server{
		Addr:    s.Config.Server.Port,
		Handler: s.Handler,
	}

	return server.ListenAndServe()
}
