import React, { createContext, useState, useEffect } from "react";
import axios, { HttpStatusCode } from "axios";
import { useNavigate } from "react-router-dom";
import api from "./api";

// Create a context for authentication
export const AuthContext = createContext();

export const AuthContextProvider = ({ children }) => {
  const [isAuthentzicated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const navigate = useNavigate();

  const checkAuthStatus = async () => {
    try {
      const response = await axios.get("/healthcheck", {
        withCredentials: true,
      });
      if (response.status === HttpStatusCode.Accepted) {
        setUser(response.data);
        setIsAuthenticated(true);
        navigate("/");
      } else if (response.status === HttpStatusCode.Forbidden) {
        setIsAuthenticated(false);
        setUser(null);
        navigate("/login");
      } else {
        navigate("/error/" + response.status);
      }
    } catch (error) {
      setIsAuthenticated(false);
      setUser(null);
      navigate("/login");
    }
  };

  const signup = async (formData) => {
    try {
      const response = await api.post("/signup", formData);
      if (response.status === HttpStatusCode.SeeOther) {
        navigate("/login");
      } else {
        navigate("/error/" + response.status);
      }
    } catch (error) {
      console.error("Error submitting data:", error); // Handle error
    }
  };
  const login = async (formData) => {
    try {
      const response = await api.post("/login", formData, {
        withCredentials: false,
      });
      if (response.status === 200) {
        setUser(response.data);
        setIsAuthenticated(true);
        navigate("/");
      } else {
        navigate("/error/" + response.status);
      }
    } catch (error) {
      console.error("Error occurred:", error);
    }
  };

  // Function to logout user
  const logout = async () => {
    try {
      // Make a request to invalidate the user session or token on the backend
      // Example:
      // await axios.post("/logout");
      setUser(null);
      setIsAuthenticated(false);
      navigate("/login"); // Redirect to login page after logout
    } catch (error) {
      console.error("Error occurred:", error);
    }
  };

  useEffect(() => {
    checkAuthStatus();
  }, []);

  return (
    <AuthContext.Provider
      value={{
        isAuthenticated,
        user,
        signup,
        login,
        logout,
        isLoading,
      }}>
      {children}
    </AuthContext.Provider>
  );
};
