package main

import (
	"html/template"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t := template.New("home")
	t, _ = t.ParseFiles("templates/home.html")
	//Parse all files in the templates/static folder
	t, _ = t.ParseGlob("./templates/static/*.html")
	t.ExecuteTemplate(w, "home", nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("login")
	t, _ = t.ParseFiles("templates/login.html")
	t, _ = t.ParseGlob("./templates/static/*.html")
	t.ExecuteTemplate(w, "login", nil)
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("register")
	t, _ = t.ParseFiles("templates/register.html")
	t, _ = t.ParseGlob("./templates/static/*.html")
	t.ExecuteTemplate(w, "register", nil)
}

func mpPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("mp")
	t, _ = t.ParseFiles("templates/mp.html")
	t, _ = t.ParseGlob("./templates/static/*.html")
	t.ExecuteTemplate(w, "mp", nil)
}

func forumPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("forum")
	t, _ = t.ParseFiles("templates/forum.html")
	t, _ = t.ParseGlob("./templates/static/*.html")
	t.ExecuteTemplate(w, "forum", nil)
}

func accountPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("account")
	t, _ = t.ParseFiles("templates/account.html")
	t, _ = t.ParseGlob("./templates/static/*.html")
	t.ExecuteTemplate(w, "account", nil)
}
