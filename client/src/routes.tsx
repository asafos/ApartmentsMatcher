import { Home } from "./pages/Home";
import { Login } from "./pages/Login";

const mainRoutes = [
    {
        path: "/",
        element: <Home />,
    },
    {
        path: "/login",
        element: <Login />,
    },
];

export { mainRoutes }; 