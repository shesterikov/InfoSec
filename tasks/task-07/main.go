package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"html/template"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {
	// Initialize PostgreSQL connection
	connStr := "user=student dbname=db password=HjjnGytfnhg host=db sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Set up HTTP handlers
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("login.html"))
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	query := fmt.Sprintf("SELECT id, username, password FROM users WHERE username='%s' AND password='%s'", username, password)
	
    fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Query failed", http.StatusInternalServerError)
		log.Printf("Query failed: %v", err)
		return
	}
	defer rows.Close()

	var id int
	var dbUser, dbPass string

	if rows.Next() {
		err := rows.Scan(&id, &dbUser, &dbPass)
		if err != nil {
			http.Error(w, "Scan failed", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Welcome, %s!", dbUser)
	} else {
		fmt.Fprintf(w, "Login failed!")
	}
}