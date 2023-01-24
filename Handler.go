package main

import (
	"encoding/json"
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
		// Check if the UUID is in the database
		/* if UUIDandUsernameMatch(db, u.UUID) {
			// If it is, send 200
			w.WriteHeader(http.StatusOK)
		} else {
			// If it isn't, send 404
			w.WriteHeader(http.StatusNotFound)
		} */
		if uuidUser[u.UUID] != "" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
