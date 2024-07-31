package service

import (
	"context"
	"fmt"
	ms "messageService"

	"github.com/segmentio/kafka-go"
)

type KafkaService struct {
	writer *kafka.Writer
}

func NewKafkaService(writer *kafka.Writer) *KafkaService {
	return &KafkaService{writer: writer}
}

func (s *KafkaService) SendToKafka(ctx context.Context, message ms.NewMessage) error {
	kafkaMessage := kafka.Message{
		Key:   []byte(message.Key),
		Value: []byte(message.Payload),
	}
	fmt.Println(ctx, message, "KafkaMessage", kafkaMessage)
	err := s.writer.WriteMessages(ctx, kafkaMessage)
	if err != nil {
		return err
	}

	return nil
}
