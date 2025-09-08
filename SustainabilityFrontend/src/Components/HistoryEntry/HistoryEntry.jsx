import styles from './HistoryEntry.module.css'

function HistoryEntry() {
    return (
    <div className={styles.container}>
        <div className={styles.topRow}>
            <p className={styles.topText}>Water Bottle</p>
            <p className={styles.topText}>.006 tons</p>
        </div>
        <p className={styles.bottomText}>9/7/25</p>
    </div>
    )
}

export default HistoryEntry;