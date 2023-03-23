import { useEffect } from "react";
import { redirect, useLocation, useNavigate } from "react-router-dom";
import { useAuthStore } from "../../DAL/stores/auth";
import { DataObjectState } from "../../DAL/stores/types";
import { App } from "./App";

function AppContainer() {
    const fetchUser = useAuthStore((state) => state.fetchUser)
    const user = useAuthStore((state) => state.user)
    const navigate = useNavigate();
    const location = useLocation();

    useEffect(() => {
        if (user.state === DataObjectState.NotStarted) {
            fetchUser()
        }
    }, [user])

    if (user.state === DataObjectState.InProgress || user.state === DataObjectState.NotStarted) {
        return <div>Loading...</div>
    }

    if (user.state === DataObjectState.Failed) {
        const from = location.state?.from?.pathname || "/";
        navigate("/login", { state: { from } });
        return null
    }

    return (
        <App />
    );
}

export { AppContainer };