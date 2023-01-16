package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t := template.New("home")
	t, _ = t.ParseFiles("templates/home.html", "./templates/static/header.html", "./templates/static/footer.html", "./templates/static/ws.html")
	t.ExecuteTemplate(w, "home", nil)
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Websocket Endpoint")
}
