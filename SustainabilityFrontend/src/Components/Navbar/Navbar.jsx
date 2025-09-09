import styles from './Navbar.module.css'
import { useNavigate } from 'react-router-dom'

function Navbar() {
    let navigate = useNavigate()

    return (
        <div className={styles.navBar}>
            <div className={styles.LogoutButton} onClick={() => navigate('/login')}>
                <p>Logout</p>
            </div>
        </div>
    )
}

export default Navbar