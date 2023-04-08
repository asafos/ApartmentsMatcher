import { create } from 'zustand'
import {
  addUserApartmentPref,
  ApartmentPref,
  ApartmentPrefToAdd,
  fetchUserApartmentPrefs,
} from '../services/apartmentPrefs/apartmentPrefs'

import { DataObject, DataObjectState } from './types'

interface State {
  apartmentPrefs: DataObject<null | ApartmentPref[]>
  getUserApartmentPrefs: (userID: number) => void
  addUserApartmentPref: (apartment: ApartmentPrefToAdd) => void
}

const useApartmentPrefsStore = create<State>((set, get) => ({
  apartmentPrefs: { data: null, state: DataObjectState.NotStarted },
  getUserApartmentPrefs: async (userID) => {
    set({ apartmentPrefs: { data: null, state: DataObjectState.Fetching } })
    try {
      const { data } = await fetchUserApartmentPrefs(userID)
      set({ apartmentPrefs: { data, state: DataObjectState.Succeeded } })
    } catch (error) {
      set({ apartmentPrefs: { data: null, state: DataObjectState.Failed, error } })
    }
  },
  addUserApartmentPref: async (apartmentPref) => {
    set({ apartmentPrefs: { data: null, state: DataObjectState.Adding } })
    try {
      const { data } = await addUserApartmentPref(apartmentPref)
      set(({ apartmentPrefs }) => ({
        apartmentPrefs: {
          data: [...(apartmentPrefs.data || []), data],
          state: DataObjectState.Succeeded,
        },
      }))
    } catch (error) {
      set({ apartmentPrefs: { data: null, state: DataObjectState.Failed, error } })
    }
  },
}))

export { useApartmentPrefsStore }
