package main

// Websocket

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	// import json
	"encoding/json"

	"github.com/gorilla/websocket"
	// import uuid
	UUID "github.com/satori/go.uuid"
)

type ClientWS struct {
	SessionId string
	UserId    int
	Username  string
}

// Used for sending messages Message = switch (login, register, post, private) in json need to be parsed
var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	clients = make(map[*websocket.Conn]*ClientWS)
)

var (
	uuidUser   = make(map[string]string)
	UserCookie = make(map[string]*http.Cookie)
)

// Used for sending messages Message = switch (login, register, post, private) in json need to be parsed
type Message struct {
	Username     string      `json:"username"`
	Message      interface{} `json:"message"`
	Message_Type string      `json:"type"`
}

// Used for sending login (for user)
type LoginMessage struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Used for sending register (for user)
type RegisterMessage struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

// Used for sending posts
type PostMessage struct {
	Creator    string   `json:"creator"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
	Comments   []string `json:"comments"`
}

// Used for sending private messages
type PrivateMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

// Used for sending answer to client
type ServerAnswer struct {
	Answer string `json:"answer"`
	UUID   string `json:"uuid"`
	Type   string `json:"type"`
}

// Used for registering uuid and username in database
type UuidMessage struct {
	Uuid          string `json:"uuid"`
	Username      string `json:"username"`
	Authenticated string `json:"authenticated"`
	Expires       string `json:"expires"`
}

type User struct {
	Username string
}

// handle websocket connection
func wsHandler(w http.ResponseWriter, r *http.Request) {
	var cookie, err = r.Cookie(COOKIE_SESSION_NAME)
	if err == nil && cookie != nil {
		var sessionId = cookie.Value
		row := GetDB().QueryRow("SELECT user_id FROM session WHERE session_id = ?", sessionId)

		var sqUsername string
		var sqUserId int
		if row.Scan(&sqUserId) != sql.ErrNoRows {
			row = GetDB().QueryRow("SELECT name FROM user WHERE idUser = ?", sqUserId)
			if row.Scan(&sqUsername) != sql.ErrNoRows {
				ws, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					log.Printf("Error upgrading client: %s\n", err)
					delete(clients, ws)
					return
				}

				log.Printf("[WebSocket] New client with session %s.\n", sessionId)
				clients[ws] = &ClientWS{SessionId: sessionId, UserId: sqUserId, Username: sqUsername}

				go MessageHandler(ws)
				return
			}
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
}

type userMessage struct {
	To      string
	From    string
	Content string
}

type wsSynchronize struct {
	Messages []userMessage
	Type     string `json:"type"`
}

// handle messages from websocket
func MessageHandler(ws *websocket.Conn) {
	db := GetDB()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("[WebSocket] Dropped client.\n")
			delete(clients, ws)
			break
		}

		// switch message type (login, register, post, private) and call function
		switch msg.Message_Type {
		case "post":
			WsPost(db, ws, msg)
		case "private":
			WsPrivate(db, ws, msg)
		case "sync:messages":
			WsSynchronizeMessages(db, ws, msg)
		case "sync:users":
			WsSynchronizeUsers(db, ws)
		case "ping":
			fmt.Printf("Client %d/%s (%s) has pinged.\n", clients[ws].UserId, clients[ws].Username, clients[ws].SessionId)
			ws.WriteJSON(map[string]string{
				"request": "ping",
			})
		}
	}
}
func WsSynchronizeUsers(db *sql.DB, ws *websocket.Conn) {
	type OnlineUser struct {
		Username string `json:"username"`
		UserID   int    `json:"idUser"`
	}
	type Online struct {
		OnlineUsers []OnlineUser `json:"online"`
		Type        string       `json:"type"`
	}
	OnlineUsers := Online{Type: "sync:users"}

	for _, clientInfo := range clients {
		OnlineUsers.OnlineUsers = append(OnlineUsers.OnlineUsers, OnlineUser{Username: clientInfo.Username, UserID: clientInfo.UserId})
	}

	for client := range clients {
		client.WriteJSON(OnlineUsers)
	}

}

func WsSynchronizeMessages(db *sql.DB, ws *websocket.Conn, Message Message) {
	username := clients[ws].Username

	// get all messages from user
	rows, err := db.Query("SELECT * FROM conversations WHERE receiver = ? OR sender = ?", username, username)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var messages []userMessage
	for rows.Next() {
		var to string
		var from string
		var content string
		err := rows.Scan(&to, &from, &content)
		if err != nil {
			fmt.Println(err)
		}
		messages = append(messages, userMessage{To: to, From: from, Content: content})
	}
	// send messages to client
	fmt.Println(messages)
	ws.WriteJSON(wsSynchronize{Messages: messages, Type: "sync:messages"})
}

func RemoveUserFromUuid(username string) {
	for key, value := range uuidUser {
		if value == username {
			delete(uuidUser, key)
		}
	}
}

func CreateUserUUIDandStoreit(Username string) string {
	uuid := UUID.NewV4()
	uuidUser[uuid.String()] = Username
	// expire in 5 hours
	cookie := http.Cookie{Name: "uuid", Value: uuid.String(), Expires: time.Now().Add(5 * time.Hour)}
	UserCookie[Username] = &cookie
	return uuid.String()
}

// convert interface to []byte for json
func (m *Message) ConvertInterface() []byte {
	// convert Message to []byte
	Mes := m.Message.(string)
	return []byte(Mes)
}

// post using websocket
func WsPost(db *sql.DB, ws *websocket.Conn, Message Message) {
	Content := PostMessage{}
	json.Unmarshal(Message.ConvertInterface(), &Content)
	// post
	Creator := Content.Creator
	Title := Content.Title
	Contentt := Content.Content
	Categories := Content.Categories
	Comments := Content.Comments
	CreatePost(db, Creator, Title, Contentt, Categories, Comments)
	Answer := ServerAnswer{Answer: "success", Type: "post"}
	ws.WriteJSON(Answer)
}

// private message using websocket
func WsPrivate(db *sql.DB, ws *websocket.Conn, message Message) {
	Content := PrivateMessage{}
	json.Unmarshal(message.ConvertInterface(), &Content)
	// private message
	fmt.Println(Content)
	From := Content.From
	To := Content.To
	Contentt := Content.Content
	Date := Content.Date
	CreatePrivateMessage(db, From, To, Contentt, Date)

	newMessage := Message{Message_Type: "private", Message: Content}
	broadcastMessage(newMessage)
}

func broadcastMessage(msg Message) {
	for client := range clients {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}
