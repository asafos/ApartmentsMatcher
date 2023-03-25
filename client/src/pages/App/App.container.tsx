import { useEffect } from "react";
import { useAuthStore } from "../../DAL/stores/auth";
import { DataObjectState } from "../../DAL/stores/types";
import { App } from "./App";

function AppContainer() {
    const fetchUser = useAuthStore((state) => state.fetchUser)
    const user = useAuthStore((state) => state.user)


    useEffect(() => {
        if (user.state === DataObjectState.NotStarted) {
            fetchUser()
        }
    }, [])

    return (
        <App user={user} />
    );
}

export { AppContainer };