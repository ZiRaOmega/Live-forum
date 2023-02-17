
const login = async(ev) => {
    ev.preventDefault();

    const form = ev.target;
    const formData = new FormData(form);

    await fetch("/api/login", {
        method: "post",
        body: formData,
    }).then(r => {
        if (r.status != 200) throw new Error("Wrong username or password.");
        initWebsocket();

        router.navigate(null, "/forum");
    }).catch(r => {
        alert(r);
        return r;
    });
};

const register = async(ev) => {
    ev.preventDefault();

    const form = ev.target;
    const formData = new FormData(form);

    await fetch("/api/register", {
        method: "post",
        body: formData,
    }).then(r => {
        if (r.status != 200) throw new Error("Something got wrong, man....");
        router.navigate(null, "/login");
    }).catch(r => {
        alert(r);
        return r;
    });
};

const logout = async (ev) => {
    websocket.close();
    user.username = "";
    UserConversations = [];
    UsersOnline = [];
    UserList = [];
    Posts = [];
    userss = [];
    currentDiscussion = "";
    await fetch("/api/logout");
};

