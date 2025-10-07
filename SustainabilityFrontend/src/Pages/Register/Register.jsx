import styles from "./Register.module.css"
import {useNavigate} from 'react-router-dom'
import {useState, useEffect} from 'react'
import axios from 'axios'

function Register() {
    const [invalid, setInvalid] = useState(false)
    const [registerData, setRegisterData] = useState(null)
    let navigate = useNavigate();
    

    async function Register(RegisterData) {
        let email = RegisterData.get("email");
        let password = RegisterData.get("password");

        try {
            let data = await axios.post('http://127.0.0.1:8080/register', {
                "email": email,
                "password": password
            })
            setRegisterData(data)
        } catch (error) {
            error.code !== "ERR_NETWORK" ? console.log(error.response.data) : console.log("Error Connecting To Server")
            setRegisterData({"status": 400})
        }
    }

    useEffect(() => {
        if (registerData !== null) {
            if (registerData.status === 200) {
                navigate('/login')
            } else {
                setInvalid(true)
            }
        }

    }, [registerData])

    return (
    <div className={styles.registerBackground}>
        <form action={Register}>
            <div className={styles.registerBox}>
                <p className={styles.registerBoxHeader}>Register</p>

                {/* Crendentials */}
                <input className={styles.credentialBox} name="email" placeholder="Email"/>

                <input type="password" className={styles.credentialBox} name="password" placeholder="Password"/>

                {invalid ? <p className={styles.invalid}>Invalid Email or Password</p> : <></>}

                <input type="submit" className={styles.registerButton} name="registerButton" value="Register"/>

                {/* Register link */}
                <p className={[styles.registerLinkPretext, styles.registerBoxText].join(" ")}>Already have an account? <span onClick={() => navigate('/login')} className={styles.registerLink}>Login</span></p>
            </div>
        </form>
    </div>
    )
}

export default Register