const SwitchPage=(page)=>{
    history.pushState({}, page, page);
    FetchPage(page);
}
const FetchPage=(page)=>{
    fetch(`pages/${page}`)
    .then(response => response.text())
    .then(data => {
        document.innerHTML = data;
    });
}