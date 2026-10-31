// Package web implements HTTP routes and renders the chat interface.
package web

import (
	"fmt"
	"net/http"

	"lab_7/internal/store"
)

type application struct {
	store *store.Store
}

func New(db *store.Store) http.Handler {
	return (&application{store: db}).routes()
}

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := a.store.Ping(r.Context()); err != nil {
			http.Error(w, "Database unavailable", 503)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"ok"}`)
	})
	mux.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		css, _ := assets.ReadFile("public/style.css")
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		w.Write(css)
	})
	mux.HandleFunc("/login", a.login)
	mux.HandleFunc("/logout", a.logout)
	mux.HandleFunc("/messages", a.messages)
	mux.HandleFunc("/", a.chat)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		r.Body = http.MaxBytesReader(w, r.Body, 16*1024)
		mux.ServeHTTP(w, r)
	})
}

func method(w http.ResponseWriter, r *http.Request, allowed string) bool {
	if r.Method == allowed {
		return true
	}
	w.Header().Set("Allow", allowed)
	http.Error(w, "Method not allowed", 405)
	return false
}
