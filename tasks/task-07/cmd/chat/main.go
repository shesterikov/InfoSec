package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"lab_8/internal/store"
	"lab_8/internal/web"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://student:HjjnGytfnhg@db/db?sslmode=disable"
	}
	db, err := store.Open(dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	server := &http.Server{
		Addr:              ":8080",
		Handler:           web.New(db),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Println("Учебный чат: :8080")
	return server.ListenAndServe()
}
