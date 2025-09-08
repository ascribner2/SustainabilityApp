import { useState, useEffect } from 'react'
import styles from './Offset.module.css'


function OffsetWidget() {
    const [timePeriod, setTimePeriod] = useState(<></>)

    useEffect(() => {

    }, []);
    
    return (
        <>
            <p className={styles.Header}>CO<span className={styles.smallNumber}>2</span> Offset</p>
            <div className={styles.widgetBox}>

                <p className={styles.offsetAmount}>15.0</p>
                <p className={styles.offsetAmount}>Tons of CO₂</p>
            </div>
        </>
    )
}

export default OffsetWidget;