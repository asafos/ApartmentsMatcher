import { AMSHttpClient } from '../../clients/AMS'
import { Location } from '../types'

export type Apartment = {
  id: number
  user_id: number
  numberOfRooms: number
  price: number
  location: Location
  availableDates: [Date, Date]
  balcony: boolean
  roof: boolean
  parking: boolean
  elevator: boolean
  petsAllowed: boolean
  renovated: boolean
}
export type ApartmentToAdd = Omit<Apartment, 'id'>

const fetchUserApartments = (userID: number) => {
  return AMSHttpClient.get<Apartment[]>(`/api/v1/apartments/user/${userID}`)
}

const addUserApartment = (apartment: ApartmentToAdd) => {
  return AMSHttpClient.post<Apartment>('/api/v1/apartments', apartment)
}

export { fetchUserApartments, addUserApartment }
