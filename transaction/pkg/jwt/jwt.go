package jwt

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/logger"
)

type JWTService interface {
	GenerateJWTTokens(ctx context.Context, userID string) (*domain.JWTTokens, error)
	ParseJWTClaims(ctx context.Context, token string) (*domain.JWTClaims, error)
	ExtractJWTTokenHeader(header http.Header) (string, error)
}

type service struct {
	baseURL string
	log     logger.Logger
}

func NewService(log logger.Logger) JWTService {
	return &service{
		baseURL: "http://auth",
		log:     log,
	}
}

func (s *service) GenerateJWTTokens(ctx context.Context, userID string) (*domain.JWTTokens, error) {
	url := s.baseURL + "/api/tokens"
	payload := struct {
		UserID string `json:"user_id"`
	}{
		UserID: userID,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		s.log.Error("parse generate tokens body: ", err)
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		s.log.Error("call auth service: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.log.Error("auth service resp: ", resp.StatusCode)
		return nil, errors.New("unexpected auth service response")
	}

	var serviceResp domain.ServiceResponse[domain.JWTTokens]
	if err := json.NewDecoder(resp.Body).Decode(&serviceResp); err != nil {
		s.log.Error("parse auth service resp: ", err)
		return nil, err
	}

	return &serviceResp.Data, nil
}

func (s *service) ParseJWTClaims(ctx context.Context, token string) (*domain.JWTClaims, error) {
	url := s.baseURL + "/api/claims?token=" + token
	resp, err := http.Get(url)
	if err != nil {
		s.log.Error("call auth service: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	var serviceResp domain.ServiceResponse[domain.JWTClaims]
	if err := json.NewDecoder(resp.Body).Decode(&serviceResp); err != nil {
		s.log.Error("parse auth service resp: ", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		s.log.Warn("auth service resp: ", resp.StatusCode)
		return nil, &domain.ResponseError{
			Code:    resp.StatusCode,
			Message: serviceResp.Message,
			ErrCode: "invalid_params",
			Errors:  serviceResp.Errors,
		}
	}

	return &serviceResp.Data, nil
}
