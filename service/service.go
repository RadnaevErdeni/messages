package service

import (
	ms "messageService"
	"messageService/repository"
)

type Message interface {
	CreateMessage(mes ms.NewMessage) (int, error)
	StatusMessage() ([]ms.MessageDB, error)
}

type Service struct {
	Message
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Message: NewMessageService(repo.Message),
	}
}
