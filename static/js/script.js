const SwitchPage=(page)=>{
    history.pushState({}, page, page);
    FetchPage(page);
}
const FetchPage=(page)=>{
    fetch(`/${page}`)
    .then(response => response.text())
    .then(data => {
        document.body.innerHTML = data;
    });
}