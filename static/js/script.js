const SwitchPage=(page)=>{
    //if page=="" then page="/"
    FetchPage(page);
    if(page==""){
        page="/";
    }

    history.pushState({}, page, page);
    
}
const FetchPage=(page)=>{
    fetch(`/${page}`)
    .then(response => response.text())
    .then(data => {
        document.body.innerHTML = data;
    });
}