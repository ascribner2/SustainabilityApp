import { useState, useEffect } from 'react'
import styles from './NewItem.module.css'
import Dropdown from '../Dropdown/Dropdown.jsx'
import WaterBottleItem from '../AddItems/WaterBottleItem.jsx'

function NewItem({ updateFunc }) {
    const [itemType, setItemType] = useState("Water Bottle")
    const [item, setItem] = useState(<></>)

    useEffect(() => {
        let itemMap = {
            "Water Bottle": <WaterBottleItem updateFunc={updateFunc}/>
        }

        setItem(itemMap[itemType])

    }, [itemType]);
    
    return (
        <>
            <p className={styles.Header}>Add New Item</p>
            <div className={styles.widgetBox}>
                <Dropdown dropdownItems={["Water Bottle"]} setFunc={setItemType} selectedItem={itemType} />
                { item }
            </div>
        </>
    )
}

export default NewItem;