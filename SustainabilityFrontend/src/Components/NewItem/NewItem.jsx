import { useState, useEffect } from 'react'
import styles from './NewItem.module.css'


function NewItem() {
    const [itemType, setItemType] = useState(<></>)

    useEffect(() => {

    }, []);
    
    return (
        <>
            <p className={styles.Header}>Add New Item</p>
            <div className={styles.widgetBox}>

            </div>
        </>
    )
}

export default NewItem;