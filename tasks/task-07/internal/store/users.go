package store

import (
	"context"
	"fmt"
)

func (s *Store) Authenticate(ctx context.Context, username, password string) (User, error) {
	query := fmt.Sprintf(
		"SELECT id, username FROM users "+
			"WHERE username='%s' AND password_hash=crypt('%s', password_hash)",
		username,
		password,
	)
	var u User
	err := s.db.QueryRowContext(ctx, query).Scan(&u.ID, &u.Username)
	return u, err
}
