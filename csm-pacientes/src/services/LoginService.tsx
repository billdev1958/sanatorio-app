import axios from "axios";
import { createSignal } from "solid-js";
import { LoginUser } from "../models/Login&Registers"; // Asegúrate de que esta ruta sea correcta

export function useLoginService() {
  const [loginError, setLoginError] = createSignal<string | null>(null);
  const [isLoggingIn, setIsLoggingIn] = createSignal(false);
  const [token, setToken] = createSignal<string | null>(null);

  async function login(user: LoginUser) {
    setIsLoggingIn(true);
    setLoginError(null);

    try {
      const response = await axios.post("https://api.example.com/login", user);
      setToken(response.data.token); // Asume que el token viene en response.data.token
      console.log("Login successful:", response.data);
      return response.data;
    } catch (error: any) {
      // Si el servidor devuelve un error, capturamos el mensaje de error
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

  return {
    login,
    loginError,
    isLoggingIn,
    token,
  };
}
