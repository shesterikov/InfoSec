package demo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func usersHandler(db *sql.DB, safe bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		query := "SELECT id, name, role FROM users WHERE name = $1 ORDER BY id"
		args := []any{name}
		if !safe {
			// INTENTIONALLY VULNERABLE: user input becomes SQL syntax.
			query = fmt.Sprintf("SELECT id, name, role FROM users WHERE name = '%s' ORDER BY id", name)
			args = nil
		}
		rows, err := db.QueryContext(r.Context(), query, args...)
		if err != nil {
			http.Error(w, "query failed", 400)
			return
		}
		defer rows.Close()
		users := []user{}
		for rows.Next() {
			var u user
			if err := rows.Scan(&u.ID, &u.Name, &u.Role); err != nil {
				http.Error(w, "scan failed", 500)
				return
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, "read failed", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
