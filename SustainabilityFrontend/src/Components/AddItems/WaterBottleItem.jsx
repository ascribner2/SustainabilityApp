import styles from './WaterBottleItem.module.css'

function WaterBottleItem({ updateFunc }) {
    
    function addItem() {
        updateFunc((prev)=>{
            return !prev;
        })
    }
    
    return (
        <form className={styles.Form} action={addItem}>
            <div className={styles.Body}>
                <input type="text"  className={styles.ounces} name="ounces" placeholder="Ounces(Oz)"/>

                <input type="submit" className={styles.addButton} name="addButton" value="Add"/>
            </div>
        </form>
    )
}

export default WaterBottleItem;