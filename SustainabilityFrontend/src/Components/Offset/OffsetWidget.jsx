import { useState, useEffect } from 'react'
import styles from './Offset.module.css'
import Dropdown from '../Dropdown/Dropdown.jsx'

function OffsetWidget() {
    const [timePeriod, setTimePeriod] = useState("Annual")

    useEffect(() => {
        console.log(timePeriod);
    }, [timePeriod]);
    
    return (
        <>
            <p className={styles.Header}>CO<span className={styles.smallNumber}>2</span> Offset</p>
            <div className={styles.widgetBox}>
                <Dropdown dropdownItems={["Annual", "Month", "Week", "Day"]} setFunc={setTimePeriod} selectedItem={timePeriod}/>
                <p className={styles.offsetAmount}>15.0</p>
                <p className={styles.offsetAmount}>Tons of CO₂</p>
            </div>
        </>
    )
}

export default OffsetWidget;