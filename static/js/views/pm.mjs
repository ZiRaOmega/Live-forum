export default {
    render: () => {
        return `<div class="convHolder">
        <div class="conv">
        </div>
        <input id="sender" type="text" placeholder="Send">
        <button id="sender_button">Envoyer</button>
        </div>
        <div class="recentconv">
        <div id="recent">Recent</div>
        <div class="convs">
        <div>
        </div>`;
    },
    postRender: () => {
        document.querySelector('#sender_button').addEventListener('click', () => {
            const field = document.querySelector('#sender');
            const text = field.value;
            field.value = "";
            
            const recipient = currentDiscussion;
            sendPrivateMessage(text, recipient);
        });
    },
}
