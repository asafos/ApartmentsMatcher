import { create } from 'zustand'
import {
  Apartment,
  fetchUserApartments,
  addUserApartment,
  ApartmentToAdd,
} from '../services/apartments/apartments'
import { DataObject, DataObjectState } from './types'

interface State {
  apartment: DataObject<null | Apartment>
  getUserApartment: (userID: number) => void
  addUserApartment: (apartment: ApartmentToAdd) => void
}

const useApartmentsStore = create<State>((set) => ({
  apartment: { data: null, state: DataObjectState.NotStarted },
  getUserApartment: async (userID) => {
    set({ apartment: { data: null, state: DataObjectState.Fetching } })
    try {
      const { data } = await fetchUserApartments(userID)
      set({ apartment: { data: data?.[0] || null, state: DataObjectState.Succeeded } })
    } catch (error) {
      set({ apartment: { data: null, state: DataObjectState.Failed, error } })
    }
  },
  addUserApartment: async (apartment) => {
    set({ apartment: { data: null, state: DataObjectState.Adding } })
    try {
      const { data } = await addUserApartment(apartment)
      set({ apartment: { data, state: DataObjectState.Succeeded } })
    } catch (error) {
      set({ apartment: { data: null, state: DataObjectState.Failed, error } })
    }
  },
}))

export { useApartmentsStore }
