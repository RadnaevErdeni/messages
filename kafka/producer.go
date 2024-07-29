package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})
	return &Producer{writer: writer}
}

func (p *Producer) SendMessage(ctx context.Context, key, value []byte) error {
	err := p.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
	if err != nil {
		log.Printf("failed to write messages: %v", err)
		return err
	}
	return nil
}

func (p *Producer) Close() {
	if err := p.writer.Close(); err != nil {
		log.Printf("failed to close writer: %v", err)
	}
}
