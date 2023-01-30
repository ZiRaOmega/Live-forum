function StartMp(){


    console.log("here")
    let form = document.getElementById("sender_button")
    form.addEventListener("click", function() {
        let messageContainer = document.getElementById("conv")
        let username = document.getElementById("user2")
        let message = document.getElementById("sender")
        if (message.value != "") {
            console.log(username.innerText,':',message.value)
            let msg = new PrivateMessage(username.innerText,"guest" ,message.value, "30/12/2020")
            let wsmsg = new Message("guest", msg.Stringify(), "private")
            websocket.send(wsmsg.Stringify())
            
            message.value = ""
            let m = document.createElement("div")
            m.innerHTML = `<b>${msg.message.from}</b>: ${msg.message.content}`
            messageContainer.appendChild(m) 
        }
    })

}