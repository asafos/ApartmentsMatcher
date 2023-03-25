import { AMSHttpClient } from "../../clients/AMS";

type User = {
    id: number
}

const facebookLogin = (authCode: string) => {
    const bodyFormData = new FormData();
    bodyFormData.append('code', authCode);

    return AMSHttpClient.post<User>('/auth/facebook', bodyFormData)
}

const googleLogin = (authCode: string) => {
    const bodyFormData = new FormData();
    bodyFormData.append('code', authCode);

    return AMSHttpClient.post<User>('/auth/google', bodyFormData)
}

const fetchUser = () => {
    return AMSHttpClient.get<User>('/auth/user')
}

const loginWithFacebook = () => {
    AMSHttpClient.post('/auth/facebook')
}

export { loginWithFacebook, fetchUser, facebookLogin, googleLogin }