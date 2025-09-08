import styles from './Dashboard.module.css'
import { useNavigate } from 'react-router-dom'
import { useEffect, useState } from 'react'
import OffsetWidget from '../../Components/Offset/OffsetWidget.jsx'
import NewItem from '../../Components/NewItem/NewItem.jsx'
import HistoryWidget from '../../Components/History/HistoryWidget.jsx'

function Dashboard() {
    const [updateData, setUpdateData] = useState(false)
    let navigate = useNavigate()
    
    function test() {
        navigate("/login")
    }

    // Authenticate
    useEffect(()=>{
        if (true) {
            // navigate("/login")
        }
    }, [])

    return (
        <div className={styles.Body}>
            <div className={styles.LeftCol}>
                <OffsetWidget />
                <NewItem />
            </div>

            <div className={styles.RightCol}>
                <HistoryWidget />
            </div>
        </div>
    )
}

export default Dashboard