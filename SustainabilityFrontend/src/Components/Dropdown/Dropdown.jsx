import styles from './Dropdown.module.css'
import { useState } from 'react'

function Dropdown({dropdownItems, setFunc, selectedItem}) {
    const [showDropdown, setShowDropdown] = useState(false)

    return (
        <div className={styles.DropdownContainer}>
            <div className={styles.Dropdown} onClick={() => {
                setShowDropdown((prev) => {
                return !prev;
                })
            }}>
                <p>{selectedItem}</p>
            </div>

            {showDropdown ? 
            <div className={styles.DropdownBox}>
                { dropdownItems.map((Item, index) => {
                    return <p key={index} className={styles.DropdownItem} onClick={(e) => { 
                        setFunc(e.target.innerText) 
                        setShowDropdown(false)
                    }}>{Item}</p>
                }) }
            </div> 
            : 
            <></>}
        </div>
    )
}

export default Dropdown;