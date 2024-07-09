package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"payment-banking-x/internal/config"
	"time"
)

func CreateTopic(ctx context.Context, cfg *config.Config, topic string) error {
	numPartitions := 1
	replicationFactor := 1

	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": cfg.BootstrapServers})
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create Admin client: %s", err))
	}
	defer adminClient.Close()

	topicsMetadata, err := adminClient.GetMetadata(&topic, false, 5000)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create Admin client: %s", err))

	}

	topicExists := false
	for _, t := range topicsMetadata.Topics {
		if t.Topic == topic {
			topicExists = true
			break
		}
	}

	if !topicExists {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		results, err := adminClient.CreateTopics(ctx, []kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     numPartitions,
			ReplicationFactor: replicationFactor,
		}})
		if err != nil {
			log.Fatalf("Failed to create topic: %s", err)
		}

		for _, result := range results {
			if result.Error.Code() != kafka.ErrNoError {
				log.Fatalf("Failed to create topic %s: %v", result.Topic, result.Error)
			}
			fmt.Printf("Topic %s created\n", result.Topic)
		}
	}
	return nil
}
