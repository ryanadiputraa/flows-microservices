package notification

import (
	"encoding/json"
	"time"

	"github.com/ryanadiputraa/flows/flows-microservices/user/config"
	"github.com/streadway/amqp"
)

type MailPayload struct {
	Subject   string `json:"subject"`
	To        string `json:"to"`
	FirstName string `json:"first_name"`
	MailType  string `json:"mail_type"`
}

type NotificationService struct {
	broker *amqp.Channel
}

func NewNotificationService(config config.Config) (*NotificationService, error) {
	retryInterval := 5 * time.Second
	maxRetries := 10
	var err error
	var broker *amqp.Connection

	for i := 0; i < maxRetries; i++ {
		broker, err = amqp.Dial(config.Server.Message_Broker_Service)
		if err == nil {
			break
		}
		time.Sleep(retryInterval)
	}

	if broker == nil {
		return nil, amqp.ErrClosed
	}

	ch, err := broker.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"NotificationService", // queue name
		true,                  // durable
		false,                 // auto delete
		false,                 // exclusive
		false,                 // no wait
		nil,                   // arguments
	)
	if err != nil {
		return nil, err
	}

	return &NotificationService{
		broker: ch,
	}, nil
}

func (s *NotificationService) SendMail(dto MailPayload) error {
	body, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}

	return s.broker.Publish(
		"",                    // exchange
		"NotificationService", // queue name
		false,                 // mandatory
		false,                 // immediate
		message,               // message to publish
	)
}
