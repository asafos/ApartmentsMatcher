import { useAuthStore } from "../../DAL/stores/auth"
import { Login } from "./Login"
import { GoogleOAuthProvider } from '@react-oauth/google';

const LoginContainer = () => {
    const facebookLogin = useAuthStore(({ facebookLogin }) => facebookLogin)
    const googleLogin = useAuthStore(({ googleLogin }) => googleLogin)
    

    const handleFacebookLogin = () => {
        // facebookLogin()
    }
    const handleGoogleLogin = (authCode: string) => {
        googleLogin(authCode)
    }
    
    return (
        <GoogleOAuthProvider clientId="79402643972-cdochoam7jcng81nvmumircp12jb5skp.apps.googleusercontent.com" >
            <Login onFacebookClick={handleFacebookLogin} onGoogleLogin={handleGoogleLogin} />
        </GoogleOAuthProvider>
    )
}

export { LoginContainer }