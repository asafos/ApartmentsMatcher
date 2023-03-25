import { useAuthStore } from "../../DAL/stores/auth"
import { Login } from "./Login"
import { GoogleOAuthProvider } from '@react-oauth/google';
import { useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { DataObjectState } from "../../DAL/stores/types";

const LoginContainer = () => {
    const facebookLogin = useAuthStore(({ facebookLogin }) => facebookLogin)
    const googleLogin = useAuthStore(({ googleLogin }) => googleLogin)
    const fetchUser = useAuthStore((state) => state.fetchUser)
    const user = useAuthStore(({ user }) => user)
    const navigate = useNavigate();
    const location = useLocation();
  
    useEffect(() => {
      if (user.state === DataObjectState.NotStarted) {
        fetchUser()
      }
    }, [])

    useEffect(() => {
        if (user.data) {
            navigate("/", { state: { from: location?.state?.from || '/' } });
        }
    }, [user])

    if (user.state === DataObjectState.NotStarted) {
        return null
    }

    return (
        <GoogleOAuthProvider clientId="79402643972-cdochoam7jcng81nvmumircp12jb5skp.apps.googleusercontent.com" >
            <Login onFacebookClick={facebookLogin} onGoogleLogin={googleLogin} />
        </GoogleOAuthProvider>
    )
}

export { LoginContainer }