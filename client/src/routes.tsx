import { Home } from "./pages/Home/Home";
import { LoginContainer } from "./pages/Login/Login.container";

const mainRoutes = [
    {
        path: "/",
        element: <Home />,
    },
    {
        path: "/login",
        element: <LoginContainer />,
    },
];

export { mainRoutes }; 