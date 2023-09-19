package server

import (
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/validator"
)

func (s *Server) mapHandler() {
	_ = validator.NewValidator()
}
