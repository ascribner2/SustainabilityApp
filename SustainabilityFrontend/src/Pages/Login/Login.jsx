import styles from "./Login.module.css"
import {useNavigate} from 'react-router-dom'
import {useState, useEffect} from 'react'
import axios from 'axios'

function Login() {
    const [invalid, setInvalid] = useState(false)
    const [loginData, setLoginData] = useState(null)
    let navigate = useNavigate();
    

    async function login(LoginData) {
        let email = LoginData.get("email");
        let password = LoginData.get("password");

        try {
            let data = await axios.post('http://127.0.0.1:8080/login', {
                "email": email,
                "password": password
            }, {
                withCredentials: true
            })
            setLoginData(data)
        } catch (error) {
            error.code !== "ERR_NETWORK" ? console.log(error.response.data) : console.log("Error Connecting To Server")
            setLoginData({"status": 400})
        }
    }

    useEffect(() => {
        if (loginData !== null) {
            if (loginData.status === 200) {
                navigate('/')
            } else {
                setInvalid(true)
            }
        }

    }, [loginData])

    return (
    <div className={styles.loginBackground}>
        <form action={login}>
            <div className={styles.loginBox}>
                <p className={styles.loginBoxHeader}>Login</p>

                {/* Crendentials */}
                <input className={styles.credentialBox} name="email" placeholder="Email"/>

                <input type="password" className={styles.credentialBox} name="password" placeholder="Password"/>

                {invalid ? <p className={styles.invalid}>Incorrect Email or Password</p> : <></>}

                <input type="submit" className={styles.loginButton} name="loginButton" value="Login"/>

                {/* Register link */}
                <p className={[styles.registerLinkPretext, styles.loginBoxText].join(" ")}>Don't have an account? <span onClick={() => navigate('/register')} className={styles.registerLink}>Register</span></p>
            </div>
        </form>
    </div>
    )
}

export default Login