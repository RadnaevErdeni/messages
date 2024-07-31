package service

import (
	"context"
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

func (s *MessageService) UpdateStatus(ctx context.Context, id int, status string) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *MessageService) UpdateStatusErr(ctx context.Context, id int, status string) error {
	return s.repo.UpdateStatus(ctx, id, status)
}
