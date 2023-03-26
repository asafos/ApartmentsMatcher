import { useEffect } from "react";
import { useApartmentsStore } from "../../DAL/stores/apartments";
import { useAuthStore } from "../../DAL/stores/auth";
import { App } from "./App";

function AppContainer() {
    const user = useAuthStore(({ user }) => user)
    const getUserApartment = useApartmentsStore(({ getUserApartment }) => getUserApartment)
    const apartment = useApartmentsStore(({ apartment }) => apartment)

    useEffect(() => {
        if (user.data) {
            getUserApartment(user.data.id)
        }
    }, [user])

    return (
        <App user={user} apartment={apartment} />
    );
}

export { AppContainer };