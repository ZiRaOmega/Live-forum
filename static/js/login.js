/* var loginsubmit = document.getElementById('loginsubmit');
loginsubmit.addEventListener('click', function() {
    return LoginClick();
}); */

(async function() {
    if (await checkuuid()) {
        SwitchPage("forum")
    }    
})();