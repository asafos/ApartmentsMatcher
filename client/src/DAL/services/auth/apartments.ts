import { AMSHttpClient } from "../../clients/AMS";

export type Apartment = {
    id: number
}

const fetchUserApartments = (userID: number) => {
    return AMSHttpClient.get<Apartment[]>(`/api/v1/apartments/user/${userID}`)
}

export { fetchUserApartments }