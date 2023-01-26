package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func SendIndex(w http.ResponseWriter, r *http.Request) {
	templates, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	templates.Execute(w, nil)
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

	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
