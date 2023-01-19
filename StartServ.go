package main

import (
	"fmt"
	"net/http"
)

func StartServerHandler() {
	//Handlers
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/register", registerPage)

	//Static Files
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//Start Server
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
