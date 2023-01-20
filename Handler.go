package main

import (
	"html/template"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	makeSql(w, r)
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

func makeSql(w http.ResponseWriter, r *http.Request) {
	sqlMaker()
}
