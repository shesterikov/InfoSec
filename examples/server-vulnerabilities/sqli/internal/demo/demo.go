// Package demo contrasts SQL text construction with value binding.
package demo

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func New() (http.Handler, func(), error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, nil, err
	}
	mux := http.NewServeMux()
	for _, mode := range []string{"vulnerable", "safe"} {
		mux.HandleFunc("GET /"+mode+"/users", usersHandler(db, mode == "safe"))
	}
	return mux, func() { db.Close() }, nil
}
