class Message {
    message = {
        username: "",
        message: "",
        type: "",
    }
    Stringify() {
        return JSON.stringify(this.message);
    }
    constructor(username, message, type) {
        this.message.username = username;
        this.message.message = message;
        this.message.type = type;
    }
}
class LoginMessage {
    message = {
        username: "",
        password: "",
    }
    Stringify() {
        return JSON.stringify(this.message);
    }
    constructor(username, password) {
        this.message.username = username;
        this.message.password = password;
    }
}

class RegisterMessage {
    message = {
        username: "",
        email: "",
        age: "",
        gender: "",
        firstname: "",
        lastname: "",
        password: "",
    }
    Stringify() {
        return JSON.stringify(this.message);
    }
    constructor(username, email, age, gender, firstname, lastname, password) {
        this.message.username = username;
        this.message.email = email;
        this.message.age = age;
        this.message.gender = gender
        this.message.firstname = firstname;
        this.message.lastname = lastname;
        this.message.password = password;
    }
}

class PostMessage {
    message = {
        creator: "",
        title: "",
        content: "",
        categories: [],
        comments: [],
    }
    Stringify() {
        return JSON.stringify(this.message);
    }
    constructor(creator, title, content, categories) {
        this.message.creator = creator;
        this.message.title = title;
        this.message.content = content;
        this.message.categories = categories;
    }
}
class PrivateMessage {
    message = {
        from: "",
        to: "",
        content: "",
        date: "",
    }
    Stringify() {
        return JSON.stringify(this.message);
    }
    constructor(from, to, content, date) {
        this.message.from = from;
        this.message.to = to;
        this.message.content = content;
        this.message.date = date;
    }
}

class UuidMessage {
    message = {
        uuid: "",
        username: "",
        authenticated: "",
        expires: "",
    }
    Stringify() {
        return JSON.stringify(this.message);
    }
    constructor(uuid, username, authenticated, expires) {
        this.message.uuid = uuid;
        this.message.username = username;
        this.message.authenticated = authenticated;
        this.message.expires = expires;
    }
}

var Userconversations = []

/**
 * @type {Websocket}
 */
let websocket = null;

const synchronizeMessages = () => {
    var message = new Message("server", "request", "sync:messages");
    websocket.send(message.Stringify());
};

const synchronizeUsers = () => {
    var message = new Message("server", "request", "sync:users");
    websocket.send(message.Stringify());
};

const initWebsocket = () => {
    websocket = new WebSocket("ws://localhost:8080/ws");
    websocket.onopen = function () {
        console.log("Connected to server");
        synchronizeMessages();
        synchronizeUsers();
    }

    websocket.onmessage = function (event) {
        console.log(event.data);
        var message = JSON.parse(event.data);
        switch (message.type) {
            case "synchronize":
                console.log(message.Messages)
                Userconversations = message.Messages
                break;
        }
    }
};

initWebsocket();

/*
const login = () => {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    CreateLoginWS(username,password)
}

const register = () => {
    var username = document.getElementById("username").value;
    var email = document.getElementById("email").value;
    var age = document.getElementById("age").value;
    var gender = document.getElementById("gender").value
    var firstname = document.getElementById("firstname").value;
    var lastname = document.getElementById("lastname").value;
    var password = document.getElementById("password").value;
    CreateRegisterWS(username,email,age,gender,firstname,lastname,password)
}
*/

const post = () => {
    var creator = document.getElementById("creator").value;
    var title = document.getElementById("title").value;
    var content = document.getElementById("content").value;
    var categories = document.getElementById("categories").value;
    CreatePostWS(creator, title, content, categories)
}

const HelloWorld = () => {
    var message = new Message("Hello", "World", "hello");
    console.log(message.Stringify())
    websocket.send(message.Stringify());
}

const TestLogin = () => {
    var username = "test"
    var password = "test"
    var logmessage = new LoginMessage(username, password)
    var message = new Message(username, logmessage.Stringify(), "login")
    websocket.send(message.Stringify())
}

const TestRegister = () => {
    var username = "test"
    var email = "test@test"
    var age = "20"
    var gender = "M"
    var firstname = "Max"
    var lastname = "DIET"
    var password = "test"
    var registerMessage = new RegisterMessage(username, email, age, gender, firstname, lastname, password)
    var message = new Message(username, registerMessage.Stringify(), "register")
    websocket.send(message.Stringify())
}

function LoginClick() {
    var username = document.getElementById('username').value
    var password = document.getElementById('password').value
    document.cookie = "username=" + username;
    Username = username
    CreateLoginWS(username, password)
    console.log(username)
    document.getElementById('user').innerHTML = username
    router.navigate(event, '/forum')
    return false;
}

const CreateLoginWS = (username, password) => {
    var logmessage = new LoginMessage(username, password)
    var message = new Message(username, logmessage.Stringify(), "login")
    websocket.send(message.Stringify())
}

function RegisterClick() {
    var username = document.getElementById('username').value
    var email = document.getElementById('email').value
    var age = document.getElementById('age').value
    var gender = document.getElementById('gender').value
    var firstname = document.getElementById('firstname').value
    var lastname = document.getElementById('lastname').value
    var password = document.getElementById('password').value
    CreateRegisterWS(username, email, age, gender, firstname, lastname, password)
    Username = username
    router.navigate(event, '/forum')
    return false;
}

const CreateRegisterWS = (username, email, age, gender, firstname, lastname, password) => {
    var registerMessage = new RegisterMessage(username, email, age, gender, firstname, lastname, password)
    var message = new Message(username, registerMessage.Stringify(), "register")
    websocket.send(message.Stringify())
}

const CreatePostWS = (creator, title, content, categories) => {
    var postMessage = new PostMessage(creator, title, content, categories)
    var message = new Message(creator, postMessage.Stringify(), "post")
    websocket.send(message.Stringify())
}

const CreatePrivateMessageWS = (from, to, content, date) => {
    var privateMessage = new PrivateMessage(from, to, content, date)
    var message = new Message(from, privateMessage.Stringify(), "private")
    websocket.send(message.Stringify())
}

