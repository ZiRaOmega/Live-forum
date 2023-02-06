
export default class {
    
    constructor(root, routes) {
        this.root = root;
        this.routes = routes;
        
        window.addEventListener('popstate', async () => {
            this.root.innerHTML = this.routes[window.location.pathname].view;
        });
    }

    DOMContentLoadedHandler() {
        this.root.innerHTML = this.routes[window.location.pathname].view;
        this.dispatch();
    }

    dispatch() {
        const routeEvent = new Event("route");
        dispatchEvent(routeEvent);
    }

    async navigate(event, pathname) {
        if (event) event.preventDefault();

        window.history.pushState(
            {},
            pathname,
            window.location.origin + pathname
        );
        this.root.innerHTML = this.routes[pathname].view;
        this.dispatch();
    }
};
