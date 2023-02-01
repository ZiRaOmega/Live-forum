package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"     // Import go-sqlite3 library
	bcrypt "golang.org/x/crypto/bcrypt" // Import bcrypt library
)

type Profile struct {
	Username        string
	Email           string
	Password        string
}

func sqlMaker() {
	if !fileExists("sqlite-database.db") {
		os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records.
		// SQLite is a file based database.

		go Log("Creating sqlite-database.db...")
		file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		go Log("sqlite-database.db created")

		sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
		defer sqliteDatabase.Close()                                     // Defer Closing the database
		createUserTable(sqliteDatabase)
		createMpTable(sqliteDatabase) // Create Database Tables
		CreateUUIDTable(sqliteDatabase)
		createConversationsTable(sqliteDatabase)

		// INSERT RECORDS
		// passtest, _ := HashPassword("test")

		// fmt.Println(CheckPasswordHash("d", HashD), CheckPasswordHash("", HashD))

	}
}

func ConnectToDB() *sql.DB {
	sqlMaker()
	// Open the created SQLite File
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	return sqliteDatabase
}

func DisconnectFromDB(db *sql.DB) {
	db.Close()
}

func Log(texttolog ...interface{}) {
	log.Println(texttolog...)
}

func createUserTable(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE user (
		"idUser" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"mail" TEXT,
		"age" TEXT,
		"gender" TEXT,
		"firstname" TEXT,
		"lastname" TEXT,
		"password" TEXT
	  );` // SQL Statement for Create Table

	go Log("Create user table...")
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	go Log("user table created")
}

func createNotifsTable(db *sql.DB) {
	createNotifsTableSQL := `CREATE TABLE notifs (
		"commentID" integer,
		"username" TEXT,
		"liked" BIT,
		"disliked" BIT,
		"commented" BIT,
		"postID" integer,
		FOREIGN KEY(postID) REFERENCES post(idPost)

	  );` // SQL Statement for Create Table

	go Log("Create notifs table...")
	statement, err := db.Prepare(createNotifsTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	go Log("notifs table created")
}

func createLikeTable(db *sql.DB) {
	createLikeTableSQL := `CREATE TABLE likes (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT,
		"postID" integer,
		FOREIGN KEY(postID) REFERENCES post(idPost) 
	);`
	go Log("Create likes table...")
	statement, err := db.Prepare(createLikeTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	go Log("likes table created")
}

func createDisLikeTable(db *sql.DB) {
	createLikeTableSQL := `CREATE TABLE dislikes (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT,
		"postID" integer,
		FOREIGN KEY(postID) REFERENCES post(idPost) 
	);`
	go Log("Create Dislikes table...")
	statement, err := db.Prepare(createLikeTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	go Log("Dislikes table created")
}

func insertLike(db *sql.DB, username string, postID int) {
	go Log("[\033[33m>\033[0m] Inserting like")
	insertLikeSQL := `INSERT INTO likes (username, postID) VALUES (?, ?)`
	statement, err := db.Prepare(insertLikeSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(username, postID)
	if err != nil {
		log.Fatalln(err.Error())
	}
	insertLikeNotifSQL := `INSERT INTO notifs (commentID, username, liked, disliked, commented, postID) VALUES (0, ?, 1, 0, 0, ?)`
	statement2, err := db.Prepare(insertLikeNotifSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement2.Exec(username, postID)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func insertDisLike(db *sql.DB, username string, postID int) {
	go Log("[\033[33m>\033[0m] Inserting dislike")
	insertLikeSQL := `INSERT INTO dislikes (username, postID) VALUES (?, ?)`
	statement, err := db.Prepare(insertLikeSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(username, postID)
	if err != nil {
		log.Fatalln(err.Error())
	}
	insertDislikeNotifSQL := `INSERT INTO notifs (commentID, username, liked, disliked, commented, postID) VALUES (0, ?, 0, 1, 0, ?)`
	statement2, err := db.Prepare(insertDislikeNotifSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement2.Exec(username, postID)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func DeleteLike(db *sql.DB, username string, postID int) {
	go Log("[\033[33m>\033[0m] Removing like")
	removeLikeSQL := `DELETE FROM likes WHERE username = ? AND postID = ?`
	statement, err := db.Prepare(removeLikeSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(username, postID)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func DeleteDisLike(db *sql.DB, username string, postID int) {
	go Log("[\033[33m>\033[0m] Removing dislike")
	removeLikeSQL := `DELETE FROM dislikes WHERE username = ? AND postID = ?`
	statement, err := db.Prepare(removeLikeSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(username, postID)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func RemoveOneLike(db *sql.DB, idPost int) {
	go Log("[\033[31m-\033[0m] Remove like on post with ID :", idPost)
	RemoveOneLikeSQL := `UPDATE post SET nbr_likes = nbr_likes - 1 WHERE idPost = ?`
	statement, err := db.Prepare(RemoveOneLikeSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(idPost)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func RemoveOneDisLike(db *sql.DB, idPost int) {
	go Log("[\033[31m-\033[0m] Remove dislike on post with ID :", idPost)
	RemoveOneLikeSQL := `UPDATE post SET nbr_dislikes = nbr_dislikes - 1 WHERE idPost = ?`
	statement, err := db.Prepare(RemoveOneLikeSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(idPost)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func AddOneComment(db *sql.DB, idPost int) {
	go Log("[\033[32m+\033[0m] Adding comment on post with ID :", idPost)
	AddOneCommentSQL := `UPDATE post SET nbr_comments = nbr_comments + 1 WHERE idPost = ?`
	statement, err := db.Prepare(AddOneCommentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(idPost)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func createCommentsTable(db *sql.DB) {
	createCommentsTableSQL := `CREATE TABLE comments (
		"commentID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"comment" TEXT,
		"username" TEXT,
		"date" TEXT,
		"postID" integer,
		FOREIGN KEY(postID) REFERENCES post(idPost) 
	);`
	// Explain FOREIGN KEY(postID) REFERENCES post(idPost) :
	// The FOREIGN KEY constraint is used to prevent actions that would destroy links between tables.
	// The FOREIGN KEY constraint also prevents invalid data from being inserted into the foreign key column,
	// because it has to be one of the values contained in the table it points to.
	// The FOREIGN KEY constraint requires an INDEX on the foreign key columns if the table is to be referenced by other tables.
	// The FOREIGN KEY constraint requires that the referenced columns are indexed.
	// The FOREIGN KEY constraint requires that the referenced columns are NOT NULL.

	go Log("Create commentaries table...")
	statement, err := db.Prepare(createCommentsTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	go Log("commentaries table created")
}

func createPostTable(db *sql.DB) {
	createPostTableSQL := `CREATE TABLE post (
		"idPost" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,
		"username" TEXT,
		"profile_picture" TEXT,	
		"date" TEXT,
		"content" TEXT,
		"image_in_post" TEXT,
		"nbr_likes" integer,
		"nbr_dislikes" integer,
		"nbr_comments" integer,
		"categories" string

	  );` // SQL Statement for Create Table

	go Log("Create post table...")
	statement, err := db.Prepare(createPostTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	go Log("post table created")
}

func createMpTable(db *sql.DB) {
	createMpTableSQL := `CREATE TABLE mp (
        "sender" TEXT,        
        "receiver" TEXT,
        "content" TEXT,
        "date" TEXT
      );` // SQL Statement for Create Table

	go Log("Create mp table...")
	statement, err := db.Prepare(createMpTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	go Log("mp table created")
}

func createConversationsTable(db *sql.DB) {
	createConversationsTableSQL := `CREATE TABLE conversations (
        "sender" TEXT,        
        "receiver" TEXT,
        "lastMessage" TEXT
      );` // SQL Statement for Create Table

	go Log("Create conversations table...")
	statement, err := db.Prepare(createConversationsTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	go Log("conversations table created")
}

func insertComment(db *sql.DB, comment string, usernames string, postID int) {
	go Log("[\033[33m>\033[0m] Inserting comment")
	insertCommentarySQL := `INSERT INTO comments(comment, username, date, postID) VALUES (?, ?, ?, ?)`

	statement, err := db.Prepare(insertCommentarySQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	date := time.Now().Format(time.Stamp)
	_, err = statement.Exec(comment, usernames, date, postID)
	AddOneComment(db, postID)
	if err != nil {
		log.Fatal(err.Error())
	}
	commentID := getLatestCommentID(db)
	insertCommentNotifSQL := `INSERT INTO notifs (commentID, username, liked, disliked, commented, postID) VALUES (?, ?, 0, 0, 1, ?)`
	statement2, err := db.Prepare(insertCommentNotifSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement2.Exec(commentID, usernames, postID)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// We are passing db reference connection from main to our method with other parameters
func insertUser(db *sql.DB, name string, mail string, age string, gender string, firstname string, lastname string, password string, profile_picture string, rank string) {
	EncodedPassword, _ := HashPassword(password)
	go Log("[\033[33m>\033[0m] Inserting user record")
	insertUserSQL := `INSERT INTO user (name, mail, age, gender, firstname, lastname, password, profile_picture, rank) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertUserSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(name, mail, age, gender, firstname, lastname, EncodedPassword, profile_picture, rank)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func insertPost(db *sql.DB, title string, username string, profile_picture string, date string, content string, image_in_post string, nbr_likes int, nbr_dislikes int, nbr_comments int, categories []string) {
	go Log("[\033[33m>\033[0m] Inserting post record")
	insertPostSQL := `INSERT INTO post(title, username, profile_picture, date, content, image_in_post, nbr_likes, nbr_dislikes,nbr_comments, categories) VALUES (?, ?, ?, ?, ?, ?, ?, ?,?,?)`
	statement, err := db.Prepare(insertPostSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	categories = NoEmptyCategory(categories)
	_, err = statement.Exec(title, username, profile_picture, date, content, image_in_post, nbr_likes, nbr_dislikes, nbr_comments, strings.Join(categories, ";"))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

/*
func removePost(db *sql.DB, postID string, Username string) {
	if !IsUserAdmin(Username) {
		return
	}
	go Log("[\033[31m-\033[0m] Removing post")
	removeLikeSQL := `DELETE FROM post WHERE idPost = ?`
	statement, err := db.Prepare(removeLikeSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	idPost, err := strconv.Atoi(postID)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(idPost)
	if err != nil {
		log.Fatalln(err.Error())
	}
	RemoveAllLike(db, postID)
}

func removeComment(db *sql.DB, postID string, Username string) {
	if !IsUserAdmin(Username) {
		return
	}
	go Log("[\033[31m-\033[0m] Removing post")
	removeLikeSQL := `DELETE FROM comments WHERE commentID = ?`
	statement, err1 := db.Prepare(removeLikeSQL) // Prepare statement.
	if err1 != nil {
		log.Fatalln(err1.Error())
	}
	idPost, err1 := strconv.Atoi(postID)
	if err1 != nil {
		log.Fatalln(err1.Error())
	}
	_, err1 = statement.Exec(idPost)
	if err1 != nil {
		log.Fatalln(err1.Error())
	}
	removeFromNotifSQL := `DELETE FROM notifs WHERE commentID = ?`
	statement2, err2 := db.Prepare(removeFromNotifSQL) // Prepare statement.
	if err2 != nil {
		log.Fatalln(err2.Error())
	}
	if err2 != nil {
		log.Fatalln(err2.Error())
	}
	_, err2 = statement2.Exec(idPost)
	if err2 != nil {
		log.Fatalln(err2.Error())
	}
	RemoveAllLike(db, postID)
}

func editPost(db *sql.DB, postID string, Content string, Username string) {
	if !IsUserAdmin(Username) {
		return
	}
	go Log("[\033[31m-\033[0m] Editing post")
	editPostSQL := `UPDATE post SET content = ? WHERE idPost = ?`
	statement, err := db.Prepare(editPostSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	postId, err := strconv.Atoi(postID)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(Content, postId)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func editComment(db *sql.DB, commentID string, Content string, Username string) {
	if !IsUserAdmin(Username) {
		return
	}
	go Log("[\033[31m-\033[0m] Editing post")
	editCommentSQL := `UPDATE comments SET comment = ? WHERE commentID = ?`
	statement, err := db.Prepare(editCommentSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	commentId, err := strconv.Atoi(commentID)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(Content, commentId)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
*/

func RemoveAllLike(db *sql.DB, postID string) {
	go Log("[\033[31m-\033[0m] Removing like")
	removeLikeSQL := `DELETE FROM likes WHERE postID = ?`
	statement, err := db.Prepare(removeLikeSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	idPost, err := strconv.Atoi(postID)
	_, err = statement.Exec(idPost)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func NoEmptyCategory(categorie []string) (NoEmpty []string) {
	for i := 0; i < len(categorie); i++ {
		if categorie[i] != "" && categorie[i] != " " {
			NoEmpty = append(NoEmpty, categorie[i])
		}
	}
	return NoEmpty
}

func AddOneLike(db *sql.DB, idPost int) {
	go Log("[\033[32m+\033[0m] Adding like on post with ID :", idPost)
	addOneLikeSQL := `UPDATE post SET nbr_likes = nbr_likes + 1 WHERE idPost = ?`
	statement, err := db.Prepare(addOneLikeSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(idPost)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func AddOneDisLike(db *sql.DB, idPost int) {
	go Log("[\033[32m+\033[0m] Adding dislike on post with ID :", idPost)
	addOneLikeSQL := `UPDATE post SET nbr_dislikes = nbr_dislikes + 1 WHERE idPost = ?`
	statement, err := db.Prepare(addOneLikeSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(idPost)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func EditUserEmail(db *sql.DB, name string, newEmail string) {
	go Log("[\033[33m>\033[0m] Editing user email")
	editUserEmailSQL := `UPDATE user SET mail = ? WHERE name = ?`
	statement, err := db.Prepare(editUserEmailSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(newEmail, name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func EditUserPicture(db *sql.DB, name string, newPicture string) {
	go Log("[\033[33m>\033[0m] Editing user profile picture")
	editUserPictureSQL := `UPDATE user SET profile_picture = ? WHERE name = ?`
	statement, err := db.Prepare(editUserPictureSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(newPicture, name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func EditUserPassword(db *sql.DB, name string, newPassword string) {
	go Log("[\033[33m>\033[0m] Editing user password")
	editUserPasswordSQL := `UPDATE user SET password = ? WHERE name = ?`
	statement, err := db.Prepare(editUserPasswordSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(newPassword, name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayUsers(db *sql.DB) {
	row, err := db.Query("SELECT * FROM user ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var mail string
		var password string
		var profile_picture string
		var rank string
		row.Scan(&id, &name, &mail, &password, &profile_picture, &rank)
	}
}

func displayPosts(db *sql.DB) {
	row, err := db.Query("SELECT * FROM post ORDER BY title")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var title string
		var username string
		var profile_picture string
		var date string
		var content string
		var image_in_post string
		var nbr_likes int
		var nbr_dislikes int
		var nbr_comments int
		var categories string
		row.Scan(&id, &title, &username, &profile_picture, &date, &content, &image_in_post, &nbr_likes, &nbr_dislikes, &nbr_comments, &categories)
	}
}

func displayComments(db *sql.DB) {
	row, err := db.Query("SELECT * FROM comments ORDER BY postID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var comment string
		var username string
		var date string
		var postID int
		row.Scan(&id, &comment, &username, &date, &postID)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getLatestCommentID(db *sql.DB) int {
	row, err := db.Query("SELECT commentID FROM comments ORDER BY commentID DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var commentID int
		row.Scan(&commentID)
		return commentID
	}
	return 0
}

func IsOnline(Username string) bool {
	for k, v := range UserCookie {
		if k == Username && v.Expires.After(time.Now()) {
			return true
		}
	}
	return false
}

func GetAllUsers(db *sql.DB) (AllUSers []User) {
	row, err := db.Query("SELECT * FROM user ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var mail string
		var password string
		row.Scan(&id, &name, &mail, &password)
		AllUSers = append(AllUSers, User{Username: name})
	}
	return AllUSers
}

func GetProfileInfo(db *sql.DB, username string) Profile {
	var profile Profile
	row, err := db.Query("SELECT * FROM user ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var mail string
		var password string
		row.Scan(&id, &name, &mail, &password)
		if username == name {
			profile = Profile{Username: name, Email: mail, Password: password}
		}
	}
	return profile
}

func GetOnlineUsers(db *sql.DB) []User {
	var result []User
	row, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var name string
		var mail string
		var password string
		row.Scan(&id, &name, &mail, &password)
		result = append(result, User{Username: name, IsOnline: IsOnline(name)})
	}
	return result
}