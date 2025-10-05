import styles from './Dashboard.module.css'
import { useNavigate } from 'react-router-dom'
import { useEffect, useState } from 'react'
import Navbar from '../../Components/Navbar/Navbar.jsx'
import OffsetWidget from '../../Components/Offset/OffsetWidget.jsx'
import NewItem from '../../Components/NewItem/NewItem.jsx'
import HistoryWidget from '../../Components/History/HistoryWidget.jsx'
import axios from 'axios'

function Dashboard() {
    const [updateData, setUpdateData] = useState(false)
    const [offsetData, setOffsetData] = useState({"Items": [], "TotalOffset": 0.0})
    const [date, setDate] = useState("Annual")
    let navigate = useNavigate()

    // Authenticate
    useEffect(()=>{
        async function verify() {
            try {
                await axios.get(`http://127.0.0.1:8080/verify`, { withCredentials: true })
            } catch (error) {
                console.log(error)
                navigate("/login")
            }
        }

        verify()
    }, [])

    // Whenever the date filter is changed or a new item is added
    useEffect(()=>{
        async function update() {
            try {
                let data = await axios.get(`http://127.0.0.1:8080/getitems?timespan=${date}`, { withCredentials: true })

                // console.log(data.data)

                setOffsetData(data.data)
            } catch (error) {
                console.log(error)
            }
        }
        update()
    }, [updateData, date])

    return (
        <>
        <Navbar />
        <div className={styles.Body}>
            <div className={styles.LeftCol}>
                <OffsetWidget data={offsetData} dateFunc={setDate} date={date} />
                <NewItem updateFunc={setUpdateData} />
            </div>

            <div className={styles.RightCol}>
                <HistoryWidget data={offsetData} />
            </div>
        </div>
        </>
    )
}

export default Dashboard