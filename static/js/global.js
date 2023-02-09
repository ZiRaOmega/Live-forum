
const login = (ev) => {
    ev.preventDefault();

    const form = ev.target;
    const formData = new FormData(form);
    fetch("/api/login", {
        method: "post",
        body: formData,
    }).then(r => {
        if (r.status != 200) throw new Error("unauthorized");
        return r.status;
    }).catch(r => alert(r));

}
