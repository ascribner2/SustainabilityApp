import styles from './WaterBottleItem.module.css'

function WaterBottleItem({ updateFunc }) {
    
    function addItem() {
        updateFunc((prev)=>{
            return !prev;
        })
    }
    
    return (
        <>
            <input className={styles.ounces} name="ounces" placeholder="Ounces(Oz)"/>
        </>
    )
}

export default WaterBottleItem;