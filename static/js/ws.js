class Message {
  stringify() {
    return JSON.stringify(this);
  }

  constructor(username, message, type) {
    this.username = username;
    this.message = message;
    this.type = type;
  }
}

let user = {
  username: null,
};

let UserConversations = [];
let UsersOnline = [];
let UserList = [];

/**
 * @type {Websocket}
 */
let websocket = null;

const ping = () => {
  var message = new Message("server", "request", "ping");
  websocket.send(message.stringify());
};

const synchronizeProfile = () => {
  var message = new Message("server", "request", "sync:profile");
  websocket.send(message.stringify());
};

const synchronizeMessages = () => {
  var message = new Message("server", "request", "sync:messages");
  websocket.send(message.stringify());
};

const synchronizeUsers = () => {
  var message = new Message("server", "request", "sync:users");
  websocket.send(message.stringify());
};

const synchronizeUserList = () => {
  var message = new Message("server", "request", "sync:userList");
  websocket.send(message.stringify());
};

const initWebsocket = () => {
  if (websocket && websocket.readyState == WebSocket.OPEN) {
    console.error("already connected");
    return;
  }

  websocket = new WebSocket("ws://localhost:8080/ws");
  websocket.onopen = function () {
    console.log("Connected to server");
    synchronizeUserList();
    synchronizeProfile();
    synchronizeMessages();
    synchronizeUsers();
  };

  websocket.onmessage = function (event) {
    console.log(event.data);
    var message = JSON.parse(event.data);
    switch (message.type) {
      case "sync:profile":
        user = message.profile;
        break;
      case "sync:messages":
        console.log(message.Messages);
        UserConversations = message.Messages;
        break;
      case "sync:users":
        console.log(message.Users);
        UsersOnline = message.online;
        break;
      case "sync:userList":
        console.log(message.userList);
        UserList = message.userList;
        createList(UserList)
        break;
    }
  };
};

initWebsocket();

var userss = []

function createList(users) {
  users.forEach(item => {
      userss.push(item.username)
  });
  document.querySelector(".recentconv").appendChild(list);
}
var refresh = false
setInterval(() => {
  if (refresh) {
    refresh = false
  } else {
    refresh = true
  }}, 1000);
document.addEventListener('mousemove', ()=>{
  if (!refresh) {
    return
  }
  var crs = document.getElementsByClassName('cr');
  if (crs.length > 0) {
    return
  } else {
    const list = document.createElement("ul");
    userss.forEach(item => {
      if (item != user.username) {
        const listItem = document.createElement("li");
        const span = document.createElement("span");
        span.classList.add("dot");
        listItem.innerHTML = item;
        list.classList.add("cr");
        for (let i = 0; i < UsersOnline.length; i++) {
          if (UsersOnline[i].username == item) {
            span.classList.add("online")
          }
        }
        list.appendChild(span);
        list.appendChild(listItem);
      }
    });
    refresh=false
    document.querySelector(".recentconv").appendChild(list);
  }
});
