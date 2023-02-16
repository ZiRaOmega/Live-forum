export default {
    render: () => {
        return `<div id="postcreator">
        <input id="post_title" type="text" placeholder="Post Title">
        <input id="post_categories" type="text" placeholder="Post Categories">
        <input id="post_content" type="text" placeholder="Post Content">
        <button id="post_button">Créer</button>
        </div>`;
    },
    postRender: () => { },
}