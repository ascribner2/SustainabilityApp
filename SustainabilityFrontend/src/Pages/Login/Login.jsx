import styles from "./Login.module.css"

function Login() {
    function login(LoginData) {
        let username = LoginData.get("username");
        let password = LoginData.get("password");

        console.log(username, password);
    }

    return (
    <div className={styles.loginBackground}>
        <form action={login}>
            <div className={styles.loginBox}>
                <p className={styles.loginBoxHeader}>Login</p>

                {/* Crendentials */}
                <input className={styles.credentialBox} name="username" placeholder="Username"/>

                <input type="password" className={styles.credentialBox} name="password" placeholder="Password"/>

                <input type="submit" className={styles.loginButton} name="loginButton" value="Login"/>

                {/* Register link */}
                <p className={[styles.registerLinkPretext, styles.loginBoxText].join(" ")}>Don't have an account? <span className={styles.registerLink}>Register</span></p>
            </div>
        </form>
    </div>
    )
}

export default Login