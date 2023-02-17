export default {
    render: () => {
        return `<div id="postcreator">
        <input id="post_title" type="text" placeholder="Post Title">
        <input id="post_categories" type="text" placeholder="Post Categories">
        <input id="post_content" type="text" placeholder="Post Content">
        <button id="post_button">Créer</button>
        </div>
        <div id="postList">
        </div>`;
    },
    postRender: () => { 
        loadPosts(Posts);
        document.querySelector('#post_button').addEventListener('click', function() {
            let title = document.querySelector('#post_title').value;
            let content = document.querySelector('#post_content').value;
            let categories = document.querySelector('#post_categories').value;
            if (title != "" && content != "" && categories != "") {
                CreatePost(title, content, categories);
            }
            document.querySelector('#post_title').value = "";
            document.querySelector('#post_content').value = "";
            document.querySelector('#post_categories').value = "";
            if (document.querySelectorAll('.success').length > 0) {
                let success = document.createElement('p');
                success.innerHTML = "Post created";
                success.classList.add('successMessage');
                document.querySelector('#postcreator').appendChild(success);
            }
        });
        const respButtons = document.querySelectorAll('.resp_button');

        respButtons.forEach(function(respButton) {
          respButton.addEventListener('click', function() {
            const postDiv = respButton.closest('.post_container');
            const postCommentResponse = postDiv.querySelector('.response');
            const postId = postDiv.querySelector('.post_id');
    
            const commentResponseValue = postCommentResponse.value;
            const postIdValue = postId.value;
    
            AddComment(commentResponseValue, postIdValue);
          });
        });
    },
}