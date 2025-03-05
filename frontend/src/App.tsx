import './App.css'
import Signup from './components/Signup'
import Login from './components/Login'
import Home from './components/Home'
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { useState, useEffect } from 'react'

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      setIsLoggedIn(true);
    }
  }, []);

  const ProtectedRoute = ({ children }) => { // protect the / route
    if (!isLoggedIn) {
      return <Navigate to="/login" />;
    }
    return children;
  }

  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<ProtectedRoute>
            <Home />
          </ProtectedRoute>} /> {/*only after login */}
          <Route path="/signup" element={<Signup />} />
          <Route path="/login" element={<Login />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
