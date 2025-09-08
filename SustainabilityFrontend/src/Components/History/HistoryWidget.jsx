import { useState, useEffect } from 'react'
import styles from './History.module.css'
import HistoryEntry from '../HistoryEntry/HistoryEntry.jsx'

function HistoryWidget() {
    const [itemType, setItemType] = useState(<></>)

    useEffect(() => {

    }, []);
    
    return (
        <>
            <p className={styles.Header}>History</p>
            <div className={styles.widgetBox}>
                <HistoryEntry />
                <HistoryEntry />
                <HistoryEntry />
                <HistoryEntry />
                <HistoryEntry />
                <HistoryEntry />
                <HistoryEntry />
                <HistoryEntry />
                <HistoryEntry />
            </div>
        </>
    )
}

export default HistoryWidget;