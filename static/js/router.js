
// A simple router in vanilla javascript

const router = {
    root: document.body,
    routes: {
        "/": { from: "/" },
        "/register": { from: "/register" },
        "/login": { from: "/login" },
        "/mp": { from: "/mp" },
        "/account": { from: "/account" },
        "/forum": { from: "/forum" },
    },
    request: async(pathname) => {
        const route = router.routes[pathname];
        return fetch(route.from, { cache: "default" },)
            .then(r => r.text());
    },
    navigate: async(pathname) => {
        window.history.pushState(
            {},
            pathname,
            window.location.origin + pathname
        );
        router.root.innerHTML = await router.request(pathname);
    },
};

window.onpopstate = async () => {
    router.root.innerHTML = await router.request(window.location.pathname);
};
