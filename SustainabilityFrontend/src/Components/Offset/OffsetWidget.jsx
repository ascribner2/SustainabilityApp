import { useState, useEffect } from 'react'
import styles from './Offset.module.css'
import Dropdown from '../Dropdown/Dropdown.jsx'

function OffsetWidget({ data, dateFunc, date }) {
    const nums = {
        "Annual": 15.0,
        "Month": 5.0, 
        "Week": 2.0, 
        "Day": 1.0
    }

    return (
        <>
            <p className={styles.Header}>CO<span className={styles.smallNumber}>2</span> Offset</p>
            <div className={styles.widgetBox}>
                <Dropdown dropdownItems={["Annual", "Month", "Week", "Day"]} setFunc={dateFunc} selectedItem={date}/>
                <p className={styles.offsetAmount}>{data.toFixed(3)}</p>
                <p className={styles.offsetAmount}>kgs of CO₂e</p>
            </div>
        </>
    )
}

export default OffsetWidget;