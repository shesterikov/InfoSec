package store

import (
	"context"
	"time"
)

func (s *Store) ListMessages(ctx context.Context) ([]Message, error) {
	const query = `
        SELECT id, author, text, created_at
        FROM (
            SELECT m.id, u.username AS author, m.text, m.created_at
            FROM messages m
            JOIN users u ON u.id = m.user_id
            ORDER BY m.id DESC
            LIMIT 100
        ) recent
        ORDER BY id
    `

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []Message{}
	for rows.Next() {
		var m Message
		var created time.Time

		if err := rows.Scan(&m.ID, &m.Author, &m.Text, &created); err != nil {
			return nil, err
		}

		m.Time = created.UTC().Format(time.RFC3339)
		messages = append(messages, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *Store) AddMessage(ctx context.Context, userID int, text string) error {
	const query = `
        INSERT INTO messages (user_id, text)
        VALUES ($1, $2)
    `

	_, err := s.db.ExecContext(ctx, query, userID, text)
	return err
}
