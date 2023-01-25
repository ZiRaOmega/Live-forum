let SwitchPage=async (page)=>{
    if (page=="forum"){
        console.log(await checkuuid())
        if (await checkuuid()){
            console.log("uuid is valid")
            page="forum"
        }else{
            console.log("uuid is invalid")
            page="login"
        }
    }
    FetchPage(page);
    if(page==""){
        page="/";
    }
    history.pushState({}, page, page);
    
}
const FetchPage = async (page) => {
    await fetch(page)
        .then(response => response.text())
        .then(data => {
            document.body.innerHTML = data;
        });
}

///////////////////////////////////////////////////////////

const _IsAuthorized = async () => {
    return await fetch("/uuidcheck", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            uuid: UUID,
        })
    }).then(r => r.status === 200);
};

const _SwitchPage = async (pageId) => {
    pageId = "/" + pageId;
    if (pageId === "/forum") {
        const isAuthorized = await _IsAuthorized();
        if (!isAuthorized) pageId = "/login";
    }

    await FetchPage(pageId);
    history.pushState({
        urlPath: pageId,
    }, "", pageId);
};

SwitchPage = _SwitchPage;
