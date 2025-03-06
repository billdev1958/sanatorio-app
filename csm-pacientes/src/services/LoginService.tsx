import { createSignal } from "solid-js";
import { LoginUser } from "../models/Login&Registers";
import { useAuth } from "./AuthContext";
import api from "./Api";

export function useLoginService() {
	const [loginError, setLoginError] = createSignal<string | null>(null);
	const [isLoggingIn, setIsLoggingIn] = createSignal(false);
	const auth = useAuth();

	async function login(user: LoginUser) {
		setIsLoggingIn(true);
		setLoginError(null);

		try {
			console.log("🟡 Iniciando solicitud de login con usuario:", user);

			const response = await api.post("/login", user);

			console.log("🟢 Respuesta del login:", response.data);

			if (response.data.status === "error") {
				const combinedError = response.data.errors
					? `${response.data.message} - ${response.data.errors}`
					: response.data.message || "Login fallido";

				setLoginError(combinedError);
				return;
			}

			const receivedToken = response.data?.data?.token;
			if (receivedToken) {
				auth.login(receivedToken);
			} else {
				throw new Error("⚠ Token no encontrado en la respuesta.");
			}

			return response.data;
		} catch (error: any) {
			console.error("❌ Error durante el login:", error);

			if (error.response) {
				setLoginError(error.response.data?.message || "Error HTTP desconocido");
			} else {
				setLoginError("⚠ Error de red, intenta nuevamente.");
			}
		} finally {
			setIsLoggingIn(false);
		}
	}

	function logout() {
		auth.logout();
		console.log("🟠 Usuario cerrado sesión");
	}

	return {
		login,
		logout,
		loginError,
		isLoggingIn,
		token: auth.token,
	};
}
