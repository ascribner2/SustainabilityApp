import { useState, useEffect } from 'react'
import styles from './History.module.css'
import HistoryEntry from '../HistoryEntry/HistoryEntry.jsx'

function HistoryWidget({ data }) {
    return (
        <>
            <p className={styles.Header}>History</p>
            <div className={styles.widgetBox}>
                { data["Items"].map(item => <HistoryEntry item={item} />) }
            </div>
        </>
    )
}

export default HistoryWidget;