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

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func IsGoodCredentials(db *sql.DB, username string, password string) bool {
	//Get Password from Database
	var passwordFromDB string
	err := db.QueryRow("SELECT password FROM user WHERE username = ?", username).Scan(&passwordFromDB)
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

func RegisterUser(db *sql.DB, username string, email string, age string, gender string, firstname string, lastname string, password string) {
	//Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	//Insert User into Database
	stmt, err := db.Prepare("INSERT INTO user(name, mail, age, gender, firstname, lastname, password) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, email, age, gender, firstname, lastname, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func DidUserExist(db *sql.DB, username string) bool {
	//Get Password from Database
	//Check if email or username already exists
	SqlQuery := "SELECT name FROM user WHERE name = ? or mail = ?"
	//prepare
	stmt, err := db.Prepare(SqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	//execute
	rows, err := stmt.Query(username, username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	//check if user exists
	return rows.Next()
}
func CreatePost(db *sql.DB, Creator string, Title string, Content string, Categories []string, Comments []string) {
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

func CreatePrivateMessage(db *sql.DB, From string, To string, Content string, Date string) {

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
