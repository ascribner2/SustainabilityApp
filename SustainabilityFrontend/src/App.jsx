import './App.css'

function App() {

  return (
    <>
      <div id="loginBackground">
        <div id="loginBox">
          <p id="loginBoxHeader">Login</p>

          {/* Crendentials */}
          <div class="credentialBox">
            <p class="loginBoxText">Username</p>
          </div>
          <div class="credentialBox">
            <p class="loginBoxText">Password</p>
          </div>

          <div id="loginButton">
            <p>Login</p>
          </div>

          {/* Register link */}
          <p id="registerLinkPretext" class="loginBoxText">Don't have an account? <span id="registerLink">Register</span></p>
        </div>
      </div>
    </>
  )
}

export default App
