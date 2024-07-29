package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type KafkaDB struct {
	db *sqlx.DB
}

func NewKafka(db *sqlx.DB) *KafkaDB {
	return &KafkaDB{db: db}
}

/*
	func (s *KafkaService) ReadMessageFromKafka(ctx context.Context) (string, error) {
		msg, err := s.consumer.ReadMessage(ctx)
		if err != nil {
			return "", err
		}
		return string(msg.Value), nil
	}
*/
func (r *KafkaDB) UpdateMessageStatus(ctx context.Context, id string, status string) error {
	query := fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2", messageTable)
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}
