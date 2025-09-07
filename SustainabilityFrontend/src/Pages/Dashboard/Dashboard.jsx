import styles from './Dashboard.module.css'
import { useNavigate } from 'react-router-dom'
import { useEffect } from 'react'

function Dashboard() {
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
        <div>
            <button onClick={test}>Press me</button>
            <p>Test</p>
        </div>
    )
}

export default Dashboard