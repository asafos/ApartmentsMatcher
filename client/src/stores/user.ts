import { create } from "zustand";

const useUserStore = create(() => ({
    user: null,
    snout: true,
    fur: true
}))

export { useUserStore }