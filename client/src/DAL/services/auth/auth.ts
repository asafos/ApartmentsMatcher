import { AMSHttpClient } from "../../clients/AMS";

type User = {
    id: number
}

const fetchUser = () => {
    return AMSHttpClient.get<User>('/auth/user')
}

const loginWithFacebook = () => {
    AMSHttpClient.post('/auth/facebook')
}

export { loginWithFacebook, fetchUser }