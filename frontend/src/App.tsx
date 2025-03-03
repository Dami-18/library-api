import './App.css'
import Signup from './components/Signup'
import Login from './components/Login'
import Home from './components/Home'
import {BrowserRouter, Routes, Route} from "react-router-dom";

function App() {

  return (
    <>
      <BrowserRouter>
      <Routes>
      <Route path="/" element={<Home />} /> {/*only after login */}
      <Route path="/signup" element={<Signup />} />
      <Route path="/login" element={<Login />} />
      </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
