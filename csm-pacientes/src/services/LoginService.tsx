import axios from "axios";
import { createSignal } from "solid-js";
import { LoginUser } from "../models/Login&Registers";
import { useAuth } from './AuthContext';

const API_BASE_URL = import.meta.env.VITE_BACKEND_HOST;

export function useLoginService() {
  const [loginError, setLoginError] = createSignal<string | null>(null);
  const [isLoggingIn, setIsLoggingIn] = createSignal(false);
  const auth = useAuth();

  async function login(user: LoginUser) {
    setIsLoggingIn(true);
    setLoginError(null);

    try {
      const response = await axios.post(`${API_BASE_URL}/v1/login`, user);
    
      console.log("Login response:", response.data);
    
      if (response.data.status === "error") {
        const combinedError = response.data.errors
          ? `${response.data.message} - ${response.data.errors}`
          : response.data.message || "Login failed";
    
        setLoginError(combinedError);
        return; 
      }
    
      const receivedToken = response.data?.data?.token;
      if (receivedToken) {
        auth.login(receivedToken);
      } else {
        throw new Error("Token not found in response");
      }
    
      return response.data;
    } catch (error: any) {
      if (error.response) {
        setLoginError(error.response.data?.message || "Error HTTP desconocido");
      } else {
        setLoginError("Network error, please try again later.");
      }
      console.error("Error during login:", error);
    } finally {
      setIsLoggingIn(false);
    }    
  }

  function logout() {
    auth.logout();
    console.log("User logged out");
  }

  return {
    login,
    logout,
    loginError,
    isLoggingIn,
    token: auth.token,
  };
}