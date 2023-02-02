/* var registersubmit = document.getElementById('registersubmit');
registersubmit.addEventListener('click', function() {
    return RegisterClick();
}); */

(async function() {
    if (await checkuuid()) {
        SwitchPage("forum")
    }    
})();