import styles from './HistoryEntry.module.css'
import dayjs from 'dayjs'
import { useRef } from 'react'

function HistoryEntry({ item }) {
    const time = useRef(dayjs(item.date).format('M/D/YY'))

    return (
    <div className={styles.container}>
        <div className={styles.topRow}>
            <p className={styles.topText}>{item.title}</p>
            <p className={styles.topText}>{item.offset} kgs</p>
        </div>
        <p className={styles.bottomText}>{time.current}</p>
    </div>
    )
}

export default HistoryEntry;