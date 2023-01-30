function StartMp(){


    console.log("here")
    let form = document.getElementById("sender_button")
    form.addEventListener("click", function() {
        let messageContainer = document.getElementById("conv")
        let username = document.getElementById("user2")
        let message = document.getElementById("sender")
        if (message.value != "") {
            console.log(username.innerText,':',message.value)
            let msg = {Username: username.innerText,Message: message.value}
            websocket.send(JSON.stringify(msg))
            message.value = ""
            let m = document.createElement("div")
            m.innerHTML = `<b>${msg.Username}</b>: ${msg.Message}`
            messageContainer.appendChild(m) 
        }
    })

}