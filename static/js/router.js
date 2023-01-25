
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
        const cachedRouteData = sessionStorage.getItem(pathname);
        if (cachedRouteData !== null) return cachedRouteData;
        const data = await fetch(route.from, { cache: "no-cache" })
            .then(r => r.text());
        sessionStorage.setItem(pathname, data);
        return data;
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

window.addEventListener('DOMContentLoaded', async () => {
    const data = await router.request(window.location.pathname);
    console.log(data);
    router.root.innerHTML = data;
});

window.addEventListener('popstate', async () => {
    router.root.innerHTML = await router.request(window.location.pathname);
});
