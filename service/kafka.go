package service

import (
	"context"
	"messageService/repository"
)

type KafkaService struct {
	repo repository.Kafka
}

func NewKafkaService(repo repository.Kafka) *KafkaService {
	return &KafkaService{
		repo: repo,
	}
}

/*
	func (s *KafkaService) ReadMessageFromKafka(ctx context.Context) (string, error) {
		return s.repo.ReadMessageFromKafka(ctx)
	}
*/
func (s *KafkaService) UpdateMessageStatus(ctx context.Context, id string, status string) error {
	return s.repo.UpdateMessageStatus(ctx, id, status)
}
