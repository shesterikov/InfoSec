// Package store implements PostgreSQL persistence for the teaching chat.
package store

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

type User struct {
	ID       int
	Username string
}

type Message struct {
	ID     int
	Author string
	Text   string
	Time   string
}

func Open(dsn string) (*Store, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}
