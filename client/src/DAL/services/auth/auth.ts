import { AMSHttpClient } from "../../clients/AMS";

type User = {
    id: number
}

const facebookLogin = () => {
    return AMSHttpClient.get<User>('/auth/facebook')
}

const googleLogin = (code: string) => {
    const bodyFormData = new FormData();
    bodyFormData.append('code', code);

    return AMSHttpClient.post<User>('/auth/google', bodyFormData)
}

const fetchUser = () => {
    return AMSHttpClient.get<User>('/auth/user')
}

const loginWithFacebook = () => {
    AMSHttpClient.post('/auth/facebook')
}

export { loginWithFacebook, fetchUser, facebookLogin, googleLogin }