import styles from "./Login.module.css"
import {useNavigate} from 'react-router-dom'
import {useState, useEffect} from 'react'
import axios from 'axios'

function Login() {
    const [invalid, setInvalid] = useState(false)
    const [loginData, setLoginData] = useState(null)
    let navigate = useNavigate();
    

    async function login(LoginData) {
        let username = LoginData.get("username");
        let password = LoginData.get("password");

        try {
            let data = await axios.post('http://127.0.0.1:8080/login', {
                "email": username,
                "password": password
            })
            setLoginData(data)
        } catch (error) {
            setLoginData(error.response.data)
        }
    }

    useEffect(() => {
        if (loginData !== null) {
            if (loginData.status === 200) {
                navigate('/')
            } else {
                console.log(loginData)
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
                <input className={styles.credentialBox} name="username" placeholder="Username"/>

                <input type="password" className={styles.credentialBox} name="password" placeholder="Password"/>

                {invalid ? <p className={styles.invalid}>Incorrect Username or Password</p> : <></>}

                <input type="submit" className={styles.loginButton} name="loginButton" value="Login"/>

                {/* Register link */}
                <p className={[styles.registerLinkPretext, styles.loginBoxText].join(" ")}>Don't have an account? <span className={styles.registerLink}>Register</span></p>
            </div>
        </form>
    </div>
    )
}

export default Login