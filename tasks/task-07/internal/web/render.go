package web

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed views/*.html public/style.css
var assets embed.FS
var templates = template.Must(template.ParseFS(assets, "views/*.html"))

func render(w http.ResponseWriter, name string, status int, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if err := templates.ExecuteTemplate(w, name, data); err != nil {
		log.Printf("Template: %v", err)
	}
}

func loginPage(w http.ResponseWriter, status int, message, username string) {
	render(w, "login.html", status, struct {
		Error    string
		Username string
	}{
		Error:    message,
		Username: username,
	})
}
