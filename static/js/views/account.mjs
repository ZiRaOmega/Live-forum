export default {
    render: () => {
        return `<div class="profile">
        <img class="pp">
        <div><p id="profile-username"></p></div>
        </div>`;
    },
    postRender: () => {
        setTimeout(() => {
            document.querySelector('#profile-username').innerText = user.username;
        }, 500);
    },
}