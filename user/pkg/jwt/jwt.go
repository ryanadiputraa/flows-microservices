package jwt

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/logger"
)

type JWTService interface {
	GenerateJWTTokens(ctx context.Context, userID string) (*domain.JWTTokens, error)
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		s.log.Error("read auth service resp: ", err)
		return nil, err
	}

	var serviceResp domain.ServiveResponse[domain.JWTTokens]
	if err = json.Unmarshal(respBody, &serviceResp); err != nil {
		s.log.Error("parse auth service resp: ", err)
		return nil, err
	}

	return &serviceResp.Data, nil
}
