import './App.css'
import Login from './Pages/Login/Login.jsx'
import Layout from './Pages/Layout/Layout.jsx'
import Dashboard from './Pages/Dashboard/Dashboard.jsx'
import { BrowserRouter, Routes, Route } from 'react-router-dom'

function App() {
  

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />}/>
        <Route element={<Layout/>}>
          <Route path="/" element={<Dashboard />}/>
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
