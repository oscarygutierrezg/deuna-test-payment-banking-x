package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"payment-banking-x/internal/config"
	"payment-banking-x/internal/service"
	"payment-banking-x/pkg/dto"
	"payment-banking-x/pkg/dto/enums"
)

type Consumer interface {
	Consume(s *service.Services)
}

func NewPaymentConsumer(cfg *config.Config) (*paymentConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"group.id":          cfg.GroupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	return &paymentConsumer{consumer: c, topic: cfg.ConsumerTopic}, nil
}

type paymentConsumer struct {
	consumer *kafka.Consumer
	topic    string
}

func (c *paymentConsumer) Consume(s *service.Services) {
	defer c.consumer.Close()

	err := c.consumer.Subscribe(c.topic, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}

	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err == nil {
			var message dto.PaymentRequest
			err := json.Unmarshal(msg.Value, &message)
			if err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

			switch message.Type {
			case enums.Payment:
				s.Payment.Create(&message)
			case enums.Refund:
				s.Refund.Create(&message)
			default:
				fmt.Printf("Received message: %+v\n", message)
			}
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
