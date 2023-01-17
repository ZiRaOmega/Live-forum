package main

import (
	"net/http"
	"fmt"
)

func StartServerHandler() {
	//Handlers
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsHandler)

	//Static Files
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//Start Server
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
