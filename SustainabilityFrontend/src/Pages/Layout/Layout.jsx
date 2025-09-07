import styles from "./Layout.module.css"
import { Outlet } from 'react-router-dom'

function Layout() {
    return (
        <div className={styles.LayoutBody}>
            <Outlet />
        </div>
    )
}

export default Layout