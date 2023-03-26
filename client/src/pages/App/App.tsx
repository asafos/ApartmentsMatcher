import { useEffect } from "react";
import { useLocation, useNavigate, useRoutes } from "react-router-dom";
import { Apartment } from "../../DAL/services/apartments/apartments";
import { User } from "../../DAL/services/auth/auth";
import { DataObject, DataObjectState } from "../../DAL/stores/types";
import { appRoutes } from "../../routes";

type Props = {
    user: DataObject<User | null>
    apartment: DataObject<Apartment | null>
}

function App(props: Props) {
    const { user, apartment } = props
    const routeResult = useRoutes(appRoutes);
    const navigate = useNavigate();
    const location = useLocation();

    useEffect(() => {
        if (user.state === DataObjectState.Failed) {
            const from = location.state?.from?.pathname || "/";
            navigate("/login", { state: { from } });
        }
    }, [user])

    useEffect(() => {
        if (apartment.state === DataObjectState.Succeeded && !apartment.data) {
            navigate("/apartment/add");
        }
    }, [apartment])

    if (
        user.state === DataObjectState.Fetching
        || user.state === DataObjectState.NotStarted
        || apartment.state === DataObjectState.NotStarted
        || apartment.state === DataObjectState.Fetching
    ) {
        return <div>Loading...</div>
    }

    if (user.state === DataObjectState.Failed) {
        return null
    }

    return (
        <>
            {routeResult}
        </>
    );
}

export { App };