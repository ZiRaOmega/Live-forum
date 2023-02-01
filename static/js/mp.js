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
            let currentDate = new Date();
            let hour = currentDate.getHours().toString().padStart(2, '0');
            let minute = currentDate.getMinutes().toString().padStart(2, '0');
            let time = `${hour}:${minute}`; 
            let m = document.createElement("div")
            m.innerHTML = `${time} | <b>${msg.message.from}</b>: ${msg.message.content}`
            messageContainer.appendChild(m) 
        }
    })

}

function PrintMP(msg){
    console.log(msg)
    let messageContainer = document.getElementById("conv")
    let m = document.createElement("div")
    if (msg.from == Username){
        return}
    let currentDate = new Date();
    let hour = currentDate.getHours().toString().padStart(2, '0');
    let minute = currentDate.getMinutes().toString().padStart(2, '0');
    let time = `${hour}:${minute}`;   
    m.innerHTML = `${time} | <b>${msg.from}</b>: ${msg.content}`
    messageContainer.appendChild(m) 
}