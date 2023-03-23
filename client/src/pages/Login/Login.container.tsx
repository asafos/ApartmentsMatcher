import { Login } from "./Login"


const LoginContainer = () => {
    const handleFacebookLogin = () => {
        
    }
    const handleGoogleLogin = () => {

    }
    return <Login onFacebookClick={handleFacebookLogin} onGoogleClick={handleGoogleLogin}  />
}

export { LoginContainer }