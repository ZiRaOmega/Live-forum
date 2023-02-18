export default {
  render: () => {
    return `<div class="convHolder">
        <div class="conv">
        </div>
        <div id="currentDiscussion"></div>
        <form id="form-send-message">
        <input id="sender" type="text" placeholder="Send">
        <input type="submit" id="sender_button" value="Envoyer"></input>
        </form>
        </div>
        <div class="recentconv">
        <div id="recent">Recent</div>
        <div class="convs">
        <div>
        </div>`;
  },
  postRender: () => {
    document
      .querySelector("#form-send-message")
      .addEventListener("submit", (ev) => {
        ev.preventDefault();

        const field = document.querySelector("#sender");
        let text = field.value;
        if (text != "") {
          field.value = "";

          text = text.replaceAll("$", "🐧");

          const recipient = currentDiscussion;
          sendPrivateMessage(text, recipient);
        }
      });
  },
};
