package repository

import (
	ms "messageService"

	"github.com/jmoiron/sqlx"
)

type Message interface {
	CreateMessage(mes ms.NewMessage) (int, error)
	StatusMessage() ([]ms.MessageDB, error)
}

type Repository struct {
	Message
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Message: NewMessage(db),
	}
}
