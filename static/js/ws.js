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
  const list = document.createElement("ul");
  userss = []
  users.forEach(item => {
    if (item.username != user) {
      userss.push(item.username)
      const listItem = document.createElement("li");
      listItem.innerHTML = item.username;
      list.appendChild(listItem);
    }
  });
  document.querySelector(".recentconv").appendChild(list);
}

document.addEventListener('mousemove', ()=>{
  var crs = document.getElementsByClassName('cr');
  if (crs.length > 0) {
    return
  } else {
    const list = document.createElement("ul");
    userss.forEach(item => {
        const listItem = document.createElement("li");
        listItem.innerHTML = item;
        list.classList.add("cr")
        list.appendChild(listItem);
    });
    document.querySelector(".recentconv").appendChild(list);
  }
});

function LiveUpdateUserList() {
  var crs = document.getElementsByClassName('LiveUserList');
  if (crs.length > 0) {
    return
  } else {
    const list = document.createElement("ul");
    userss.forEach(item => {
        const listItem = document.createElement("li");
        listItem.innerHTML = item;
        list.classList.add("LiveUserList")
        list.appendChild(listItem);
    });
    document.querySelector(".recentconv").appendChild(list);
  }
}