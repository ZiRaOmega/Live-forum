//if cookies not include uuid, then switch to login page
if (document.cookie.indexOf("uuid") == -1){
    SwitchPage("login")
}