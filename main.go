package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

const COOKIE_SESSION_NAME = "SESSION_ID"
const COOKIE_SESSION_DURATION = time.Hour * 3

func SendIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Client with IP: %s\n", r.RemoteAddr)

	tmpl, err := template.ParseFiles("templates/header.html", "templates/footer.html", "templates/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		SendIndex(w, r)
	})
	http.HandleFunc("/ws", wsHandler)

	db := GetDB()
	sqlMaker(db)
	defer CloseDB()
	for _, pattern := range []string{"/login", "/register", "/pm", "/forum", "/account"} {
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			SendIndex(w, r)
		})
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		email := r.FormValue("email")
		age := r.FormValue("age")
		gender := r.FormValue("gender")
		password := r.FormValue("password")

		allFilled := true
		for _, s := range []string{username, firstname, lastname, email, age, gender, password} {
			if s == "" {
				allFilled = false
			}
		}

		if allFilled {
			hashedPassword, _ := HashPassword(password)
			GetDB().Exec(
				"INSERT INTO user (name, mail, age, gender, firstname, lastname, password) VALUES (?,?,?,?,?,?,?)",
				username, email, age, gender, firstname, lastname, hashedPassword)

			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
	})

	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username != "" && password != "" {
			row := GetDB().QueryRow("SELECT idUser, password FROM user WHERE name = ?", username)
			if row != nil {
				var sqUserId int
				var sqHashedPassword string
				if row.Scan(&sqUserId, &sqHashedPassword) == nil {
					if CheckPasswordHash(password, sqHashedPassword) {
						sessionId := uuid.NewV4().String()
						GetDB().Exec("INSERT INTO session (session_id, user_id) VALUES (?, ?)", sessionId, sqUserId)

						http.SetCookie(w, &http.Cookie{
							Name:     COOKIE_SESSION_NAME,
							Value:    sessionId,
							Expires:  time.Now().Add(COOKIE_SESSION_DURATION),
							HttpOnly: true,
							Path:     "/",
						})

						w.WriteHeader(http.StatusOK)
						return
					}
				}
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
	})

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
