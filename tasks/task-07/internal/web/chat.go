package web

import (
	"net/http"
	"strings"
	"unicode/utf8"

	"lab_7/internal/store"
)

func (a *application) chat(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if !method(w, r, http.MethodGet) {
		return
	}
	u, ok := a.requireUser(w, r)
	if !ok {
		return
	}

	messages, err := a.store.ListMessages(r.Context())
	if err != nil {
		http.Error(w, "Messages unavailable", 500)
		return
	}

	render(w, "chat.html", 200, struct {
		User     store.User
		Messages []store.Message
	}{
		User:     u,
		Messages: messages,
	})
}

func (a *application) messages(w http.ResponseWriter, r *http.Request) {
	if !method(w, r, http.MethodPost) {
		return
	}
	u, ok := a.requireUser(w, r)
	if !ok {
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Некорректная форма", 400)
		return
	}

	text := strings.TrimSpace(r.PostForm.Get("text"))
	if text == "" || utf8.RuneCountInString(text) > 1000 {
		http.Error(w, "Сообщение должно содержать 1–1000 символов.", 400)
		return
	}

	if err := a.store.AddMessage(r.Context(), u.ID, text); err != nil {
		http.Error(w, "Message creation failed", 500)
		return
	}
	http.Redirect(w, r, "/", 303)
}
