if (!checkuuid()){
    SwitchPage("login")
}

function converseWith(username) {
    if (username != document.getElementById('receiver').value) {
        document.getElementById('receiver').value = username
        console.log("Targeting  : ",username)
        const element = document.getElementById("conv");
        const divs = element.getElementsByTagName("div");
        while (divs.length > 0) {
            element.removeChild(divs[0]);
        }
        var messages = GetMessageFrom(document.getElementById('user2').innerHTML, username)
        for (var i = 0; i < messages.length; i++) {
            console.log(messages[i])
            PrintMP(messages[i])
        }
        document.getElementById('target').innerHTML = username
    }
    if (document.getElementById('user2').innerHTML == username) {
        document.getElementById('target').innerHTML = username + " (You)"
    }
}
function GetMessageFrom(From, To) {
    var messages = []
    for (var i = 0; i < userMessages.length; i++) {
        if (userMessages[i].From == From && userMessages[i].To == To) {
            messages.push(userMessages[i])
        }
    }
    return messages
}
function StartMp(){

    let form = document.getElementById("mp__send-form")
    form.addEventListener("submit", function(e) {
        e.preventDefault();

        let messageContainer = document.getElementById("conv")
        let username = document.getElementById("user2")
        let message = document.getElementById("sender")
        let receiverhtml = document.getElementById("receiver")
        console.log("Sending message to ",receiverhtml.value)
        let receiver=""
        if (receiverhtml.value == ""){
            receiver = "guest"
        } else {
            receiver = receiverhtml.value
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
    m.innerHTML = `${time} | <b>${msg.From}</b>: ${msg.Content}`
    messageContainer.appendChild(m) 
}
//hihi