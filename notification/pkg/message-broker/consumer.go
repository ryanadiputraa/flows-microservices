package messagebroker

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ryanadiputraa/flows/flows-microservices/notification/config"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/email"
	"github.com/ryanadiputraa/flows/flows-microservices/notification/pkg/logger"
	"github.com/streadway/amqp"
)

type MessageBrokerConsumer struct {
	broker  *amqp.Channel
	log     logger.Logger
	service email.Usecase
}

func NewMessageBrokerConsumer(config config.Config, log logger.Logger, service email.Usecase) (*MessageBrokerConsumer, error) {
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

	return &MessageBrokerConsumer{
		broker:  ch,
		log:     log,
		service: service,
	}, nil
}

func (m *MessageBrokerConsumer) Listen() error {
	forever := make(chan bool)

	messages, err := m.broker.Consume(
		"NotificationService", // queue name
		"",                    // consumer
		true,                  // auto-ack
		false,                 // exclusive
		false,                 // no local
		false,                 // no wait
		nil,                   // arguments
	)
	if err != nil {
		return err
	}

	go func() {
		for message := range messages {
			var dto domain.EmailDTO
			if err = json.Unmarshal(message.Body, &dto); err != nil {
				m.log.Warn("fail to parse message body", err)
				continue
			}

			m.service.RegisterNotification(context.Background(), dto)
		}
	}()

	<-forever

	return nil
}
