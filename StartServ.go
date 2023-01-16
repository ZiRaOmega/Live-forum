package main

import "net/http"

func StartServerHandler() {
	//Handlers
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsHandler)

	//Static Files
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//Start Server
	http.ListenAndServe(":8080", nil)
}
