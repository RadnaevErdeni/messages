package service

import (
	"context"
	ms "messageService"
	"messageService/repository"

	"github.com/segmentio/kafka-go"
)

type Message interface {
	CreateMessage(mes ms.NewMessage) (int, error)
	StatusMessage() ([]ms.MessageDB, error)
	UpdateStatus(ctx context.Context, id int, status string) error
	UpdateStatusErr(ctx context.Context, id int, status string) error
}
type Kafka interface {
	SendToKafka(ctx context.Context, message ms.NewMessage) error
}

type Service struct {
	Message
	Kafka
}

func NewService(repo *repository.Repository, kafkaWriter *kafka.Writer) *Service {
	return &Service{
		Message: NewMessageService(repo.Message),
		Kafka:   NewKafkaService(kafkaWriter),
	}
}
