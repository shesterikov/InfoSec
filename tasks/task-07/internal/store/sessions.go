package store

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
)

func tokenHash(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func (s *Store) SessionUser(ctx context.Context, token string) (User, error) {
	var u User
	if len(token) != 64 {
		return u, sql.ErrNoRows
	}
	const query = `
        SELECT u.id, u.username
        FROM sessions s
        JOIN users u ON u.id = s.user_id
        WHERE s.token_hash = $1 AND s.expires_at > NOW()
    `

	err := s.db.QueryRowContext(ctx, query, tokenHash(token)).Scan(&u.ID, &u.Username)
	return u, err
}

func (s *Store) RotateSession(ctx context.Context, oldToken string, userID int) (string, error) {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	raw := hex.EncodeToString(token)
	// Rotate the session atomically; never store the bearer token in the database.
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	defer tx.Rollback()
	const deleteQuery = `
        DELETE FROM sessions
        WHERE token_hash = $1 OR expires_at <= NOW()
    `
	const insertQuery = `
        INSERT INTO sessions (token_hash, user_id, expires_at)
        VALUES ($1, $2, NOW() + INTERVAL '1 hour')
    `

	if _, err = tx.ExecContext(ctx, deleteQuery, tokenHash(oldToken)); err == nil {
		_, err = tx.ExecContext(ctx, insertQuery, tokenHash(raw), userID)
	}

	if err != nil {
		return "", err
	}

	if err = tx.Commit(); err != nil {
		return "", err
	}

	return raw, nil
}

func (s *Store) DeleteSession(ctx context.Context, token string) error {
	const query = `DELETE FROM sessions WHERE token_hash = $1`

	_, err := s.db.ExecContext(ctx, query, tokenHash(token))
	return err
}
