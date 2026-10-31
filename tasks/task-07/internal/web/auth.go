package web

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"unicode/utf8"
)

func (a *application) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		loginPage(w, 200, "", "")
		return
	}

	if !method(w, r, http.MethodPost) {
		return
	}

	if err := r.ParseForm(); err != nil {
		loginPage(w, 400, "Некорректная форма.", "")
		return
	}

	username := strings.TrimSpace(r.PostForm.Get("username"))
	password := r.PostForm.Get("password")
	if username == "" || utf8.RuneCountInString(username) > 100 ||
		password == "" || utf8.RuneCountInString(password) > 200 {
		loginPage(w, 400, "Заполните имя и пароль допустимой длины.", username)
		return
	}
	u, err := a.store.Authenticate(r.Context(), username, password)
	if errors.Is(err, sql.ErrNoRows) {
		loginPage(w, 401, "Неверное имя пользователя или пароль.", username)
		return
	}

	if err != nil {
		loginPage(w, 400, "Ошибка SQL-запроса.", username)
		return
	}
	raw, err := a.store.RotateSession(r.Context(), sessionToken(r), u.ID)
	if err != nil {
		http.Error(w, "Session creation failed", 500)
		return
	}

	setCookie(w, raw, 3600)
	http.Redirect(w, r, "/", 303)
}

func (a *application) logout(w http.ResponseWriter, r *http.Request) {
	if !method(w, r, http.MethodPost) {
		return
	}

	if err := a.store.DeleteSession(r.Context(), sessionToken(r)); err != nil {
		http.Error(w, "Logout failed", 500)
		return
	}

	setCookie(w, "", -1)
	http.Redirect(w, r, "/login", 303)
}
