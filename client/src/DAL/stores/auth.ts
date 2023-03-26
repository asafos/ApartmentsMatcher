import { create } from "zustand";
import { facebookLogin, fetchUser, googleLogin, User } from "../services/auth/auth";
import { DataObject, DataObjectState } from "./types";

interface State {
    user: DataObject<null | User>
    fetchUser: () => void
    facebookLogin: (authCode: string) => void
    googleLogin: (authCode: string) => void
}

const useAuthStore = create<State>((set) => ({
    user: { data: null, state: DataObjectState.NotStarted },
    facebookLogin: async (authCode: string) => {
        const user = await facebookLogin(authCode)
        set({ user: { data: user.data, state: DataObjectState.Succeeded } })
    },
    googleLogin: async (authCode) => {
        const user = await googleLogin(authCode)
        set({ user: { data: user.data, state: DataObjectState.Succeeded } })
    },
    fetchUser: async () => {
        set({ user: { data: null, state: DataObjectState.Fetching } })
        try {
            const { data } = await fetchUser()
            set({ user: { data, state: DataObjectState.Succeeded } })
        } catch (error) {
            set({ user: { data: null, state: DataObjectState.Failed, error } })
        }
    }
}))

export { useAuthStore }