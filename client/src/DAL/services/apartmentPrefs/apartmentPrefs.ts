import { AMSHttpClient } from '../../clients/AMS'
import { Location } from '../types'

export type ApartmentPref = {
  id: number
  user_id: number
  numberOfRooms: [number, number]
  price: [number, number]
  location: Location[]
  availableDates: [Date, Date]
  balcony: boolean
  roof: boolean
  parking: boolean
  elevator: boolean
  petsAllowed: boolean
  renovated: boolean
}
export type ApartmentPrefToAdd = Omit<ApartmentPref, 'id'>

const fetchUserApartmentPrefs = (userID: number) => {
  return AMSHttpClient.get<ApartmentPref[]>(`/api/v1/apartmentPrefs/user/${userID}`)
}

const addUserApartmentPref = (apartmentPref: ApartmentPrefToAdd) => {
  return AMSHttpClient.post<ApartmentPref>('/api/v1/apartmentPrefs', apartmentPref)
}

export { fetchUserApartmentPrefs, addUserApartmentPref }
