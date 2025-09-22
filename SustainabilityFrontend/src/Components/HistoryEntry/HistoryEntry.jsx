import styles from './HistoryEntry.module.css'
import dayjs from 'dayjs'
import { useRef } from 'react'

function HistoryEntry({ item }) {
    let time = dayjs(item.date).format('M/D/YY')

    return (
    <div className={styles.container}>
        <div className={styles.topRow}>
            <p className={styles.topText}>{item.title}</p>
            <p className={styles.topText}>{item.offset} kgs</p>
        </div>
        <p className={styles.bottomText}>{time}</p>
    </div>
    )
}

export default HistoryEntry;