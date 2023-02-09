
const login = async(ev) => {
    ev.preventDefault();

    const form = ev.target;
    const formData = new FormData(form);

    await fetch("/api/login", {
        method: "post",
        body: formData,
    }).then(r => {
        if (r.status != 200) throw new Error("Wrong username or password.");
        router.navigate(null, "/forum");
    }).catch(r => {
        alert(r);
        return r;
    });
};

const register = async(ev) => {
    ev.preventDefault();

    const form = ev.target;
    const formData = new FormData(form);

    await fetch("/api/register", {
        method: "post",
        body: formData,
    }).then(r => {
        if (r.status != 200) throw new Error("S'thing got wrong, man....");
        router.navigate(null, "/login");
    }).catch(r => {
        alert(r);
        return r;
    });
};
