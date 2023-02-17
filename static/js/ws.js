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
  websocket.send(
    JSON.stringify({
      type: "private",
      message: {
        from: user.username,
        to: recipient,
        content: message,
        date: Date.now().toString(),
      },
    }),
  );
};
const CreatePost = (title, content, categories) => {
  websocket.send(
    JSON.stringify({
      type: "post",
      message: {
        title: title,
        username: user.username,
        date: Date.now().toString(),
        content: content,
        categories: categories,
      },
    }),
  );
};
const AddComment = (content, postID) => {
  websocket.send(
    JSON.stringify({
      type: "comment",
      message: {
        username: user.username,
        content: content,
        postID: postID,
      },
      username: user.username,
    }),
  );
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
        createList(UserList, message.Messages);
        break;
      case "sync:posts":
        console.log(message.posts);
        Posts = message.posts;
        loadPosts(Posts);
        break;
    }
    loadConversation(currentDiscussion);
  };
};

initWebsocket();

//PM page

var userss = [];

const GetConversation = (user) => {
  return UserConversations.filter((conversation) => {
    return conversation.From === user || conversation.To === user;
  });
};

const GetMessagesSorted = (user) => {
  return GetConversation(user).sort((a, b) => {
    if (parseInt(a.Date) < parseInt(b.Date)) {
      return -1;
    }
    if (parseInt(a.Date) > parseInt(b.Date)) {
      return 1;
    }
    return 0;
  });
};

const GetLastMessage = (user) => {
  const messages = GetMessagesSorted(user);
  return messages[messages.length - 1] || { Date: "0" };
};

function createList(users) {
  userss = [];
  users.forEach((item) => {
    userss.push(item.username);
    });
    //document.querySelector(".convs").appendChild(list);
  }

  var refresh = false;
  setInterval(() => {
    if (refresh) {
      refresh = false;
    } else {
      refresh = true;
    }
  }, 1000);
  document.addEventListener("mousemove", () => {
  if (!refresh) {
    return;
  }
  var crs = document.getElementsByClassName("cr");
  const list = document.createElement("ul");

  userss.sort((a, b) => {
    const aMsgDate = GetLastMessage(a).Date || "0";
    const bMsgDate = GetLastMessage(b).Date || "0";
    if (parseInt(aMsgDate) < parseInt(bMsgDate)) return 1;
    else if (parseInt(aMsgDate) > parseInt(bMsgDate)) return -1;
    else return 0;
  });

  userss.forEach((item) => {
    if (item != user.username) {
      const span = document.createElement("span");
      const user = document.createElement("p");
      user.addEventListener("click", function () {
        loadConversation(item);
      });
      user.textContent = item;
      span.classList.add("dot");
      list.classList.add("cr");
      for (let i = 0; i < UsersOnline.length; i++) {
        if (UsersOnline[i].username == item) {
          span.classList.add("online");
        }
      }
      list.appendChild(span);
      list.appendChild(user);
    }
  });
  refresh = false;
  if (document.querySelector(".convs") != null) {
    document.querySelector(".convs").innerHTML = "";
    document.querySelector(".convs").appendChild(list);
  }
});

setTimeout(() => {
  document.querySelector("#user").innerText = user.username;
}, 500);

function loadConversation(user) {
  let userMessages = [];
  currentDiscussion = user;
  if (document.querySelector("#currentDiscussion") != null) {
    document.querySelector("#currentDiscussion").innerText = currentDiscussion;
  }
  if (UserConversations != undefined) {
    for (let i = 0; i < UserConversations.length; i++) {
      if (
        UserConversations[i].To == user ||
        UserConversations[i].From == user
      ) {
        userMessages.push(UserConversations[i]);
      }
    }
  }
  if (document.querySelector(".conv") != null) {
    document.querySelector(".conv").innerHTML = "";
  }
  for (let j = 0; j < userMessages.length; j++) {
    let p = document.createElement("p");
    if (userMessages[j].To == user) {
      p.classList.add("sent");
    } else {
      p.classList.add("received");
    }
    p.innerText = userMessages[j].Content;
    if (document.querySelector(".conv") != null) {
      document.querySelector(".conv").appendChild(p);
    }
  }
}

//Forum page

function loadPosts(posts) {
  if (document.querySelector("#postList") != null) {
    document.querySelector("#postList").innerHTML = "";
  }
  for (let i = 0; i < posts.length; i++) {
    let container = document.createElement("div");
    let title = document.createElement("h2");
    let username = document.createElement("p");
    let date = document.createElement("p");
    let content = document.createElement("p");
    let categories = document.createElement("p");
    let comment = document.createElement("p");
    title.innerHTML = posts[i].title;
    username.innerHTML = posts[i].username;
    postDate = new Date(Number(posts[i].date)).toUTCString();
    date.innerHTML = postDate;
    content.innerHTML = posts[i].content;
    categories.innerHTML = posts[i].categories;
    for (let j = 0; j < posts[i].comments.length; j++) {
      comment.innerHTML +=
        posts[i].comments[j].username +
        ": " +
        posts[i].comments[j].comment +
        "<br>";
    }
    container.classList.add("post_container");
    title.classList.add("post_title");
    username.classList.add("post_username");
    date.classList.add("post_date");
    content.classList.add("post_content");
    categories.classList.add("post_categories");
    comment.classList.add("post_comment");
    comment.style.display = "none";
    container.appendChild(title);
    container.appendChild(username);
    container.appendChild(date);
    container.appendChild(content);
    container.appendChild(categories);
    container.appendChild(comment);
    document.querySelector("#postList").appendChild(container);

    container.addEventListener("click", (ev) => {
      const comments = ev.currentTarget.querySelector(".post_comment");
      if (comments.style.display === "none") comments.style.display = "block";
      else comments.style.display = "none";
    });
  }
}
