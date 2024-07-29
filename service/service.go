package service

import (
	"context"
	ms "messageService"
	"messageService/repository"
)

type Message interface {
	CreateMessage(mes ms.NewMessage) (int, error)
	StatusMessage() ([]ms.MessageDB, error)
}
type Kafka interface {
	UpdateMessageStatus(ctx context.Context, id string, status string) error
}

type Service struct {
	Message
	Kafka
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Message: NewMessageService(repo.Message),
		Kafka:   NewKafkaService(repo.Kafka),
	}
}
