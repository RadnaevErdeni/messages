package service

import (
	ms "messageService"
	"messageService/repository"
)

type MessageService struct {
	repo repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{repo: repo}
}
func (s *MessageService) CreateMessage(mes ms.NewMessage) (int, error) {
	return s.repo.CreateMessage(mes)
}

func (s *MessageService) StatusMessage() ([]ms.MessageDB, error) {
	return s.repo.StatusMessage()
}
