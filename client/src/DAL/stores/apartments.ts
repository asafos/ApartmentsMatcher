import { create } from "zustand";
import { Apartment, fetchUserApartments } from "../services/auth/apartments";
import { DataObject, DataObjectState } from "./types";

interface State {
    apartment: DataObject<null | Apartment>
    getUserApartment: (userID: number) => void
}

const useApartmentsStore = create<State>((set) => ({
    apartment: { data: null, state: DataObjectState.NotStarted },
    getUserApartment: async (userID) => {
        set({ apartment: { data: null, state: DataObjectState.InProgress } })
        try {
            const { data } = await fetchUserApartments(userID)
            set({ apartment: { data: data?.[0] || null, state: DataObjectState.Succeeded } })
        } catch (error) {
            set({ apartment: { data: null, state: DataObjectState.Failed, error } })
        }
    }
}))

export { useApartmentsStore }