import { AddApartmentContainer } from "./pages/App/AddApartment/AddApartment.container";
import { AppContainer } from "./pages/App/App.container";
import { Home } from "./pages/App/Home/Home";
import { LoginContainer } from "./pages/Login/Login.container";

const mainRoutes = [
    {
        path: "/",
        element: <AppContainer />,
    },
    {
        path: "/login",
        element: <LoginContainer />,
    },
];

const appRoutes = [
    {
        path: "/",
        element: <Home />,
    },
    {
        path: "/apartment/add",
        element: <AddApartmentContainer />,
    },
];

export { mainRoutes, appRoutes }; 