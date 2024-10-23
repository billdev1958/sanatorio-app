import axios from "axios";
import { createSignal } from "solid-js";
import { LoginUser } from "../models/Login&Registers";
import { useAuth } from './AuthContext';

export function useLoginService() {
  const [loginError, setLoginError] = createSignal<string | null>(null);
  const [isLoggingIn, setIsLoggingIn] = createSignal(false);
  const auth = useAuth();

  async function login(user: LoginUser) {
    setIsLoggingIn(true);
    setLoginError(null);

    try {
      const response = await axios.post("http://localhost:8080/v1/login", user);
      console.log("Login successful:", response.data);

      const receivedToken = response.data?.data?.token;
      if (receivedToken) {
        auth.login(receivedToken); // Usa el método de login del contexto de autenticación
      } else {
        throw new Error("Token not found in response");
      }

      return response.data;
    } catch (error: any) {
      if (error.response) {
        setLoginError(error.response.data.message || "Login failed");
      } else {
        setLoginError("Network error, please try again later.");
      }
      console.error("Error during login:", error);
    } finally {
      setIsLoggingIn(false);
    }
  }

  function logout() {
    auth.logout(); // Utiliza el método de logout del contexto de autenticación
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
