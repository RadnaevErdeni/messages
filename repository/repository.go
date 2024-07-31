package repository

import (
	"context"
	ms "messageService"

	"github.com/jmoiron/sqlx"
)

type Message interface {
	CreateMessage(mes ms.NewMessage) (int, error)
	StatusMessage() ([]ms.MessageDB, error)
	UpdateStatus(ctx context.Context, id int, status string) error
	UpdateStatusErr(ctx context.Context, id int, status string) error
}

type Repository struct {
	Message
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Message: NewMessage(db),
	}
}
