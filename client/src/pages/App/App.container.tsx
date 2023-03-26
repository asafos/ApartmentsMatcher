import { useEffect } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import { useApartmentsStore } from "../../DAL/stores/apartments";
import { useAuthStore } from "../../DAL/stores/auth";
import { DataObjectState } from "../../DAL/stores/types";
import { App } from "./App";

function AppContainer() {
    const user = useAuthStore(({ user }) => user)
    const getUserApartment = useApartmentsStore(({ getUserApartment }) => getUserApartment)
    const apartment = useApartmentsStore(({ apartment }) => apartment)
    const navigate = useNavigate();
    const location = useLocation();

    useEffect(() => {
        if (user.data) {
            getUserApartment(user.data.id)
        }
        if (user.state === DataObjectState.Failed) {
            const from = location.state?.from?.pathname || "/";
            navigate("/login", { state: { from } });
        }
    }, [user])

    useEffect(() => {
        if (apartment.state === DataObjectState.Succeeded) {
            if (!apartment.data) {
                navigate("/apartment/add");
            } else {
                navigate("/");
            }
        }
    }, [apartment])

    const isLoading = user.state === DataObjectState.Fetching
        || user.state === DataObjectState.NotStarted
        || apartment.state === DataObjectState.NotStarted
        || apartment.state === DataObjectState.Fetching

    return (
        <App user={user.data} isLoading={isLoading} />
    );
}

export { AppContainer };