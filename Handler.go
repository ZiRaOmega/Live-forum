package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Page struct {
	Username string
}

type PageUsers struct {
	Username string
	Users    []User
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t := template.New("home")
	t, _ = t.ParseFiles("templates/home.html")
	// Parse all files in the templates/static folder
	t, _ = t.ParseGlob("./templates/static/*.html")

	// If the user is logged in, send the username to the template
	// so that the navbar can display the username
	// if uuid cookie exists
	username := ""
	if _, err := r.Cookie("uuid"); err == nil {
		UUID, err := r.Cookie("uuid")
		if err != nil {
			fmt.Println(err)
		}
		username = uuidUser[UUID.Value]
	}
	if username == "" {
		username = "Guest"
	}
	p := Page{Username: username}
	t.ExecuteTemplate(w, "home", p)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("login")
	t, _ = t.ParseFiles("templates/login.html")
	t, _ = t.ParseGlob("./templates/static/*.html")
	// if uuid cookie exists
	username := ""
	if _, err := r.Cookie("uuid"); err == nil {
		UUID, err := r.Cookie("uuid")
		if err != nil {
			fmt.Println(err)
		}
		username = uuidUser[UUID.Value]
	}
	if username == "" {
		username = "Guest"
	}
	p := Page{Username: username}
	fmt.Println(p)
	t.ExecuteTemplate(w, "login", p)
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("register")
	t, _ = t.ParseFiles("templates/register.html")
	t, _ = t.ParseGlob("./templates/static/*.html")
	// if uuid cookie exists
	username := ""
	if _, err := r.Cookie("uuid"); err == nil {
		UUID, err := r.Cookie("uuid")
		if err != nil {
			fmt.Println(err)
		}
		username = uuidUser[UUID.Value]
	}
	if username == "" {
		username = "Guest"
	}
	p := Page{Username: username}
	t.ExecuteTemplate(w, "register", p)
}

func mpPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("mp")
	t, _ = t.ParseFiles("templates/mp.html")
	t, _ = t.ParseGlob("./templates/static/*.html")
	// if uuid cookie exists
	username := ""
	if _, err := r.Cookie("uuid"); err == nil {
		UUID, err := r.Cookie("uuid")
		if err != nil {
			fmt.Println(err)
		}
		username = uuidUser[UUID.Value]
	}
	if username == "" {
		username = "Guest"
	}
	db := GetDB()
	users := GetAllUsers(db)
	p := PageUsers{Username: username, Users: users}
	t.ExecuteTemplate(w, "mp", p)
}

func forumPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("forum")
	t, _ = t.ParseFiles("templates/forum.html")
	t, _ = t.ParseGlob("./templates/static/*.html")
	// if uuid cookie exists
	username := ""
	if _, err := r.Cookie("uuid"); err == nil {
		UUID, err := r.Cookie("uuid")
		if err != nil {
			fmt.Println(err)
		}
		username = uuidUser[UUID.Value]
	}
	if username == "" {
		username = "Guest"
	}
	p := Page{Username: username}
	t.ExecuteTemplate(w, "forum", p)
}

func accountPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("account")
	t, _ = t.ParseFiles("templates/account.html")
	t, _ = t.ParseGlob("./templates/static/*.html")
	// if uuid cookie exists
	username := ""
	if _, err := r.Cookie("uuid"); err == nil {
		UUID, err := r.Cookie("uuid")
		if err != nil {
			fmt.Println(err)
		}
		username = uuidUser[UUID.Value]
	}
	if username == "" {
		username = "Guest"
	}
	p := Page{Username: username}
	t.ExecuteTemplate(w, "account", p)
}

// LE PROBLEME VIENT DE LA FONCTION CI-DESSOUS
func uuidCheck(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	case "POST":
		type uuid struct {
			UUID     string `json:"uuid"`
			Username string `json:"username"`
		}
		var u uuid
		db := GetDB()
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&u)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		defer r.Body.Close()

		fmt.Println("UUID: "+u.UUID, "Username: "+u.Username)
		if uuidUser[u.UUID] == u.Username && UserCookie[u.Username].Expires.After(time.Now()) {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
