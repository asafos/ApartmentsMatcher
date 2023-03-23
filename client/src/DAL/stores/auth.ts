import { create } from "zustand";
import { fetchUser } from "../services/auth/auth";
import { DataObject, DataObjectState } from "./types";

type User = {
    id: number
}

interface State {
    user: DataObject<null | User>
    fetchUser: () => void
}

const useAuthStore = create<State>((set) => ({
    user: { data: null, state: DataObjectState.NotStarted },
    fetchUser: async () => {
        set({ user: { data: null, state: DataObjectState.InProgress } })
        try {
            const {data} = await fetchUser()
            set({ user: { data, state: DataObjectState.Succeeded } })
        } catch (error) {
            set({ user: { data: null, state: DataObjectState.Failed, error } })
        }
    }
}))

export { useAuthStore }