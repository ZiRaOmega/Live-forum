package main

//Websocket

import (
	"fmt"
	"log"
	"net/http"

	//import json
	"encoding/json"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024}
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
)

type Message struct {
	Username     string      `json:"username"`
	Message      interface{} `json:"message"`
	Message_Type string      `json:"type"`
}
type LoginMessage struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterMessage struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}
type PostMessage struct {
	Creator    string   `json:"creator"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
	Comments   []string `json:"comments"`
}
type PrivateMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

func ListenforMessages(ws *websocket.Conn) {
	go MessageHandler()
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Client Connected")
	clients[ws] = true

	go ListenforMessages(ws)

}

func MessageHandler() {
	for {
		Message := <-broadcast
		fmt.Println(Message.Message_Type)

		switch Message.Message_Type {
		case "login":
			WsLogin(Message)
		case "register":
			WsRegister(Message)
		case "post":
			WsPost(Message)
		case "private":
			WsPrivate(Message)
		case "hello":
			fmt.Println(Message)
		}
	}
}
func WsLogin(Message Message) { //Working
	Content := LoginMessage{}
	//convert interface to LoginMessage
	json.Unmarshal(Message.ConvertInterface(), &Content)
	//login user
	Username := Content.Username
	Password := Content.Password
	//fmt.Println("login", Username)
	if IsGoodCredentials(Username, Password) {
		//login
	} else {
		//error
	}
}
func (m *Message) ConvertInterface() []byte {
	//convert Message to []byte
	Mes := m.Message.(string)
	return []byte(Mes)
}
func WsRegister(Message Message) {
	Content := RegisterMessage{}
	json.Unmarshal(Message.ConvertInterface(), &Content)
	//register user
	Username := Content.Username
	Email := Content.Email
	Age := Content.Age
	Gender := Content.Gender
	FirstName := Content.FirstName
	LastName := Content.LastName
	Password := Content.Password
	fmt.Println(Username, Email, Age, Gender, FirstName, LastName, Password)
	if DidUserExist(Username) || DidUserExist(Email) {
		//error
	} else {
		//register
		RegisterUser(Username, Email, Age, Gender, FirstName, LastName, Password)
	}

}

func WsPost(Message Message) {
	Content := PostMessage{}
	json.Unmarshal(Message.ConvertInterface(), &Content)
	//post
	Creator := Content.Creator
	Title := Content.Title
	Contentt := Content.Content
	Categories := Content.Categories
	Comments := Content.Comments
	CreatePost(Creator, Title, Contentt, Categories, Comments)
}

func WsPrivate(Message Message) {
	Content := PrivateMessage{}
	json.Unmarshal(Message.ConvertInterface(), &Content)
	//private message
	From := Content.From
	To := Content.To
	Contentt := Content.Content
	Date := Content.Date
	CreatePrivateMessage(From, To, Contentt, Date)
}
