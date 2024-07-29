package kafka

import (
	"context"
	"log"
	"messageService/repository"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
	repo   repository.Kafka
}

func NewConsumer(brokers []string, groupID, topic string, repo repository.Kafka) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: groupID,
		Topic:   topic,
	})
	return &Consumer{reader: reader, repo: repo}
}

func (c *Consumer) StartConsuming(ctx context.Context) {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("failed to read messages: %v", err)
			continue
		}

		log.Printf("received message: %s", string(msg.Value))

		err = c.repo.UpdateMessageStatus(ctx, string(msg.Key), "processed")
		if err != nil {
			log.Printf("failed to update message status: %v", err)
		}
	}
}

func (c *Consumer) Close() {
	if err := c.reader.Close(); err != nil {
		log.Printf("failed to close reader: %v", err)
	}
}
