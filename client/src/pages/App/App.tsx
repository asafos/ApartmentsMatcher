import { useRoutes } from "react-router-dom";
import { appRoutes } from "../../routes";

function App() {
    const routeResult = useRoutes(appRoutes);
    return (
        <>
            {routeResult}
        </>
    );
}

export { App };