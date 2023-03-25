import { create } from "zustand";
import { facebookLogin, fetchUser, googleLogin } from "../services/auth/auth";
import { DataObject, DataObjectState } from "./types";

export type User = {
    id: number
}

interface State {
    user: DataObject<null | User>
    fetchUser: () => void
    facebookLogin: (authCode: string) => void
    googleLogin: (authCode: string) => void
}

const useAuthStore = create<State>((set) => ({
    user: { data: null, state: DataObjectState.NotStarted },
    facebookLogin: async (authCode: string) => {
        facebookLogin(authCode)
    },
    googleLogin: async (authCode) => {
        googleLogin(authCode)
    },
    fetchUser: async () => {
        set({ user: { data: null, state: DataObjectState.InProgress } })
        try {
            const { data } = await fetchUser()
            set({ user: { data, state: DataObjectState.Succeeded } })
        } catch (error) {
            set({ user: { data: null, state: DataObjectState.Failed, error } })
        }
    }
}))

export { useAuthStore }