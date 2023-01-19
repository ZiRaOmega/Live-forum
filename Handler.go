package main

import (
	"html/template"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t := template.New("home")
	t, _ = t.ParseFiles("templates/home.html", "./templates/static/header.html", "./templates/static/footer.html", "./templates/static/ws.html", "./templates/static/js.html")
	t.ExecuteTemplate(w, "home", nil)
}
func loginPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("login")
	t, _ = t.ParseFiles("templates/login.html", "./templates/static/header.html", "./templates/static/footer.html", "./templates/static/ws.html", "./templates/static/js.html")
	t.ExecuteTemplate(w, "login", nil)
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("register")
	t, _ = t.ParseFiles("templates/register.html", "./templates/static/header.html", "./templates/static/footer.html", "./templates/static/ws.html", "./templates/static/js.html")
	t.ExecuteTemplate(w, "register", nil)
}
