import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom"; // Import BrowserRouter
import Login from "./Pages/Login";
import Signup from "./Pages/SignUp";
import Home from "./Pages/Home";
import { AuthContextProvider } from "./Context/AuthContext";
import Navbar from "./Components/Nav/Navbar";
import ErrorPage from "./Pages/Error";

export default function App() {
  return (
    <Router>
      <AuthContextProvider>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/error/:errorcode" element={<ErrorPage />} />
        </Routes>
      </AuthContextProvider>
    </Router>
  );
}
