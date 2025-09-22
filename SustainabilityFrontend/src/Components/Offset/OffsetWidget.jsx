import styles from './Offset.module.css'
import Dropdown from '../Dropdown/Dropdown.jsx'

function OffsetWidget({ data, dateFunc, date }) {
    return (
        <>
            <p className={styles.Header}>CO<span className={styles.smallNumber}>2</span> Offset</p>
            <div className={styles.widgetBox}>
                <Dropdown dropdownItems={["Annual", "Month", "Week", "Day"]} setFunc={dateFunc} selectedItem={date}/>
                <p className={styles.offsetAmount}>{data["TotalOffset"].toFixed(3)}</p>
                <p className={styles.offsetAmount}>kgs of CO₂e</p>
            </div>
        </>
    )
}

export default OffsetWidget;