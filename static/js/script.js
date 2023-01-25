const SwitchPage=(page)=>{
    if (page=="forum"){
        if (checkuuid()){
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
const FetchPage=(page)=>{
    fetch(`/${page}`)
    .then(response => response.text())
    .then(data => {
        document.body.innerHTML = data;
    });
}
//LE PROBLEME VIENT DE LA FONCTION CHECKUUID
const checkuuid=()=>{
    console.log(document.cookie)
    if (UUID==""){
        console.log("here")
        SwitchPage("login")
    }else{
        //Post request fetch to /uuidcheck if 200, then switch to main page else switch to login page
        return fetch("/uuidcheck", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                uuid: UUID
            })
        }).then(res => {
            console.log(res.status)
            if (res.status == 200){
                return true
            }else{
                return false
            }
        })
    }
}