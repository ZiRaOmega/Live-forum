package main

//Import sql package
import (
	"database/sql"
	"log"

	//mattn/go-sqlite3
	_ "github.com/mattn/go-sqlite3"
	//bcrypt
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB = nil

func IsGoodCredentials(username string, password string) bool {
	//Get Password from Database
	var passwordFromDB string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&passwordFromDB)
	if err != nil {
		log.Fatal(err)
	}

	//Compare Passwords
	err = bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(password))
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func RegisterUser(username string, email string, age string, gender string, firstname string, lastname string, password string) {
	//Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	//Insert User into Database
	stmt, err := db.Prepare("INSERT INTO users(username, email, age, gender, firstname, lastname, password) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, email, age, gender, firstname, lastname, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func DidUserExist(username string) bool {
	//Get Password from Database
	//Check if email or username already exists
	var usernameFromDB string
	err := db.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&usernameFromDB)
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = db.QueryRow("SELECT email FROM users WHERE email = ?", username).Scan(&usernameFromDB)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true

}
func CreatePost(Creator string, Title string, Content string, Categories []string, Comments []string) {
	//Insert Post into Database
	stmt, err := db.Prepare("INSERT INTO posts(creator, title, content, categories, comments) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(Creator, Title, Content, Categories, Comments)
	if err != nil {
		log.Fatal(err)
	}
}

func CreatePrivateMessage(From string, To string, Content string, Date string) {

	//Insert Private Message into Database
	stmt, err := db.Prepare("INSERT INTO private_messages(from, to, content, date) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(From, To, Content, Date)
	if err != nil {
		log.Fatal(err)
	}
}
