
var UUID = "";
var Username = "Guest";

const SwitchPage = async(page) => {
    if (page === "forum" || page === "mp" || page === "account") {
        let uuidIsValid = await checkuuid();
        if (uuidIsValid) {
            console.log("uuid is valid");
        } else {
            console.log("uuid is invalid");
            page="login";
        }
    }

    await FetchPage(page);
    document.getElementById('user').innerText = Username;
    
    if (page === "") page = "/";

    history.pushState({}, page, page);
    if (page === "mp") {
        //timeout to wait for the page to load
        setTimeout(StartMp, 1000);
    }
}

const FetchPage = async (page) => {
    fetch(`/${page}`)
        .then(response => response.text())
        .then(data => {
            document.body.innerHTML = data;
        });
}

const checkuuid = async () => {
    return fetch("/uuidcheck", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            uuid: UUID,
            username: Username
        })
    }).then(response => response.status === 200);
}

