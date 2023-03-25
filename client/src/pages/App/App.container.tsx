import { useAuthStore } from "../../DAL/stores/auth";
import { App } from "./App";

function AppContainer() {
    const user = useAuthStore((state) => state.user)

    return (
        <App user={user} />
    );
}

export { AppContainer };