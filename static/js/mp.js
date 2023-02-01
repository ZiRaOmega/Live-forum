function StartMp(){


    console.log("here")
    let form = document.getElementById("sender_button")
    form.addEventListener("click", function() {
        let messageContainer = document.getElementById("conv")
        let username = document.getElementById("user2")
        let message = document.getElementById("sender")
        let receiverhtml = document.getElementById("receiver")
        let receiver=""
        if (receiverhtml == undefined||(receiverhtml.value == "") ){
            receiver = "guest"
        }
        if (message.value != "") {
            console.log(username.innerText,':',message.value)
            let messageText = message.value;
            message.value = "";

            let currentDate = new Date();
            let day = currentDate.getDate().toString().padStart(2, '0');
            let month = (currentDate.getMonth() + 1).toString().padStart(2, '0');
            let year = currentDate.getFullYear();
            let hour = currentDate.getHours().toString().padStart(2, '0');
            let minute = currentDate.getMinutes().toString().padStart(2, '0');
            let time = `${hour}:${minute}`; 
            let timeformat = `${day}/${month}/${year} ${hour}:${minute}`;
            let msg = new PrivateMessage(username.innerText,receiver ,messageText, timeformat)
            let wsmsg = new Message(receiver, msg.Stringify(), "private")
            let m = document.createElement("div")
            m.innerHTML = `${time} | <b>${msg.message.from}</b>: ${msg.message.content}`
            messageContainer.appendChild(m) 
            websocket.send(wsmsg.Stringify())
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