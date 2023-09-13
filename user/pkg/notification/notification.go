package notification

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ryanadiputraa/flows/flows-microservices/user/config"
	"github.com/ryanadiputraa/flows/flows-microservices/user/pkg/response"
	"github.com/sirupsen/logrus"
)

type MailPayload struct {
	Subject   string `json:"subject"`
	To        string `json:"to"`
	FirstName string `json:"first_name"`
	MailType  string `json:"mail_type"`
}

type NotificationService struct {
	baseURL string
}

func NewNotificationService(config config.Config) *NotificationService {
	return &NotificationService{
		baseURL: config.Server.Notification_Service,
	}
}

func (s *NotificationService) SendMail(dto MailPayload) (*response.HttpResponsePayload[any], error) {
	body, err := json.Marshal(dto)
	if err != nil {
		return &response.HttpResponsePayload[any]{}, err
	}

	url := s.baseURL + "/api/register"
	logrus.Info(url)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return &response.HttpResponsePayload[any]{}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &response.HttpResponsePayload[any]{}, err
	}

	var data response.HttpResponsePayload[any]
	if err = json.Unmarshal(respBody, &data); err != nil {
		return &response.HttpResponsePayload[any]{}, err
	}

	return &data, err
}
