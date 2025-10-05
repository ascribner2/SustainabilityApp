import styles from './Navbar.module.css'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'

function Navbar() {
    let navigate = useNavigate()

    async function logout() {
        try {
            await axios.get(`http://127.0.0.1:8080/logout`, { withCredentials: true })
            navigate('/login')
        } catch (error) {
            console.log(error)
        }
    }

    return (
        <div className={styles.navBar}>
            <div className={styles.LogoutButton} onClick={() => logout()}>
                <p>Logout</p>
            </div>
        </div>
    )
}

export default Navbar