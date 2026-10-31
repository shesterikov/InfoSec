package web

import (
	"database/sql"
	"errors"
	"net/http"

	"lab_7/internal/store"
)

const cookieName = "sqli_session"

func sessionToken(r *http.Request) string {
	c, err := r.Cookie(cookieName)
	if err != nil {
		return ""
	}

	return c.Value
}

func setCookie(w http.ResponseWriter, token string, age int) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   age,
	})
}

func (a *application) currentUser(r *http.Request) (store.User, error) {
	token := sessionToken(r)
	if len(token) != 64 {
		return store.User{}, sql.ErrNoRows
	}

	return a.store.SessionUser(r.Context(), token)
}

func (a *application) requireUser(w http.ResponseWriter, r *http.Request) (store.User, bool) {
	u, err := a.currentUser(r)
	if errors.Is(err, sql.ErrNoRows) {
		http.Redirect(w, r, "/login", 303)
		return u, false
	}

	if err != nil {
		http.Error(w, "Session lookup failed", 500)
		return u, false
	}

	return u, true
}
