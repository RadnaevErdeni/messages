package repository

import (
	"context"
	ms "messageService"

	"github.com/jmoiron/sqlx"
)

type Message interface {
	CreateMessage(mes ms.NewMessage) (int, error)
	StatusMessage() ([]ms.MessageDB, error)
}
type Kafka interface {
	UpdateMessageStatus(ctx context.Context, id string, status string) error
}

type Repository struct {
	Message
	Kafka
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Message: NewMessage(db),
		Kafka:   NewKafka(db),
	}
}
