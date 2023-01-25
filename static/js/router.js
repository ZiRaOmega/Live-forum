
// A simple router in vanilla javascript
// Should be global and stay alive at all time.
// Register load.

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
        return await fetch(route.from, { cache: "no-cache" })
            .then(r => r.text());
    },
    navigate: async(event, pathname) => {
        if (event) event.preventDefault();
        window.history.pushState(
            {},
            pathname,
            window.location.origin + pathname
        );
        router.root.innerHTML = await router.request(pathname);
        
        const navigateEvent = new Event("navigate");
        dispatchEvent(navigateEvent);
    },
};

window.addEventListener('DOMContentLoaded', async () => {
    router.root.innerHTML = await router.request(window.location.pathname);

    const navigateEvent = new Event("navigate");
    dispatchEvent(navigateEvent);
});

window.addEventListener('popstate', async () => {
    router.root.innerHTML = await router.request(window.location.pathname);
});
