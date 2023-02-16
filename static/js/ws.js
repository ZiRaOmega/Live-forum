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
let Posts = [];
let currentDiscussion = "";

/**
 * @type {Websocket}
 */
let websocket = null;

const ping = () => {
  var message = new Message("server", "request", "ping");
  websocket.send(message.stringify());
};

const sendPrivateMessage = (message, recipient) => {
  websocket.send(JSON.stringify({
    "type": "private",
    "message": {
      "from": user.username,
      "to": recipient,
      "content": message,
      "date": Date.now().toString(),
    }
  }));
};

const synchronizeProfile = () => {
  var message = new Message("server", "request", "sync:profile");
  websocket.send(message.stringify());
};

const synchronizeMessages = () => {
  var message = new Message("server", "request", "sync:messages");
  websocket.send(message.stringify());
};
const synchronizePosts = () => {
  var message = new Message("server", "request", "sync:posts");
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
    synchronizePosts();
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
        UserConversations = message.Messages;
        createList(UserList, message.Messages)
        break;
      case "sync:posts":
        console.log(message.posts);
        Posts = message.posts;
    }
  };
};

initWebsocket();

var userss = []

function createList(users) {
  userss = []
  users.forEach(item => {
      userss.push(item.username)
  });
  //document.querySelector(".convs").appendChild(list);
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
    const list = document.createElement("ul");
    userss.forEach(item => {
      if (item != user.username) {
        const span = document.createElement("span");
        const user = document.createElement("p");
        user.addEventListener('click', function() {
          loadConversation(item);
        });
        user.textContent = item;
        span.classList.add("dot");
        list.classList.add("cr");
        for (let i = 0; i < UsersOnline.length; i++) {
          if (UsersOnline[i].username == item) {
            span.classList.add("online")
          }
        }
        list.appendChild(span);
        list.appendChild(user);
      }
    });
    refresh=false
    document.querySelector(".convs").innerHTML = "";
    document.querySelector(".convs").appendChild(list);
});

setTimeout(() => {
  document.querySelector('#user').innerText = user.username;
}, 500);

function loadConversation(user) {
  let userMessages = []
  for (let i = 0; i < UserConversations.length; i++) {
    if (UserConversations[i].To == user || UserConversations[i].From == user) {
      userMessages.push(UserConversations[i]);
      currentDiscussion = user;
    }
  }
  document.querySelector('.conv').innerHTML = ""
  for (let j = 0; j < userMessages.length; j++) {
    let p = document.createElement('p');
    if (userMessages[j].To == user) {
      p.classList.add('sent')
    } else {
      p.classList.add('received')
    }
    p.innerText = userMessages[j].Content;
    document.querySelector('.conv').appendChild(p);
  }
}