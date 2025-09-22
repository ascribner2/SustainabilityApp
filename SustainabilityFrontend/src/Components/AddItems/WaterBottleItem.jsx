import styles from './WaterBottleItem.module.css'
import axios from 'axios'
import dayjs from 'dayjs'

function WaterBottleItem({ updateFunc }) {
    
    async function addItem(inputData) {
        try {
            let ounces = inputData.get("ounces")
            let timestamp = dayjs().format('YYYY-MM-DD HH:mm:ss')

            console.log(ounces)

            let offsetCalc = 0.00125 * ounces
            offsetCalc = (Math.round(offsetCalc * 1000) / 1000)
            offsetCalc = parseFloat(offsetCalc.toFixed(3))

            let status = await axios.post('http://localhost:8080/additem', {
                "title": "Water Bottle",
                "offset": offsetCalc,
                "date": timestamp
            })


            updateFunc((prev)=>{
                return !prev;
            })
        } catch (error) {
            console.log(error)
        }
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