import { useLocation, useNavigate, useRoutes } from "react-router-dom";
import { User } from "../../DAL/stores/auth";
import { DataObject, DataObjectState } from "../../DAL/stores/types";
import { appRoutes } from "../../routes";

type Props = {
    user: DataObject<User | null>
}

function App(props: Props) {
    const { user } = props
    const routeResult = useRoutes(appRoutes);
    const navigate = useNavigate();
    const location = useLocation();

    if (user.state === DataObjectState.InProgress || user.state === DataObjectState.NotStarted) {
        return <div>Loading...</div>
    }

    if (user.state === DataObjectState.Failed) {
        const from = location.state?.from?.pathname || "/";
        navigate("/login", { state: { from } });
        return null
    }

    return (
        <>
            {routeResult}
        </>
    );
}

export { App };