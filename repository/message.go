package repository

import (
	"context"
	"fmt"
	ms "messageService"

	"github.com/jmoiron/sqlx"
)

type MessageDB struct {
	db *sqlx.DB
}

func NewMessage(db *sqlx.DB) *MessageDB {
	return &MessageDB{db: db}
}

func (r *MessageDB) CreateMessage(mes ms.NewMessage) (int, error) {
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (message, status,date_create) values ($1, 'processing',Now()) RETURNING id", messageTable)
	row := r.db.QueryRow(createQuery, mes.Payload)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (r *MessageDB) StatusMessage() ([]ms.MessageDB, error) {
	var messages []ms.MessageDB
	query := fmt.Sprintf("SELECT id,message,status,processed_time,date_create FROM %s", messageTable)
	if err := r.db.Select(&messages, query); err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *MessageDB) UpdateStatus(ctx context.Context, id int, status string) error {
	query := fmt.Sprintf("UPDATE %s SET status = $1,processed_time=Now() WHERE id = $2", messageTable)
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}

func (r *MessageDB) UpdateStatusErr(ctx context.Context, id int, status string) error {
	query := fmt.Sprintf("UPDATE %s SET status = $1,processed_time=Now() WHERE id = $2", messageTable)
	_, err := r.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return err
	}
	return nil
}
