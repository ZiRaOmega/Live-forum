const SwitchPage=async (page)=>{
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
    await FetchPage(page);
    document.getElementById('user').innerText=Username
    if(page==""){
        page="/";
    }
    history.pushState({}, page, page);
    if (page=="mp"){
        //timeout to wait for the page to load
        setTimeout(StartMp, 1000)
        
    }
    
}
const FetchPage=async (page)=>{
    fetch(`/${page}`)
    .then(response => response.text())
    .then(data => {
        document.body.innerHTML = data;
    });
}
var UUID = ""
var Username = "Guest"
//LE PROBLEME VIENT DE LA FONCTION CHECKUUID
const checkuuid=async ()=>{
    if (UUID==""){
        SwitchPage("login")
    }else{
        //Post request fetch to /uuidcheck if 200, then switch to main page else switch to login page
        return fetch("/uuidcheck", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                uuid: UUID,
                username: Username
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