import { Component, createSignal, onMount } from "solid-js";
import { useParams } from "@solidjs/router";

const ConfirmationPage: Component = () => {
	const params = useParams();
	const [token, setToken] = createSignal<string | null>(null);

	onMount(() => {
		if (params.token) {
			setToken(params.token);
			localStorage.setItem("token", params.token);
		}
	});

	const handleConfirm = async () => {
		const currentToken = token();
		if (!currentToken) {
			alert("No se encontró token en la URL.");
			return;
		}

		try {
			const response = await fetch("https://api.ax01.dev/v1/confirm", {
				method: "POST",
				headers: {
					"Content-Type": "application/json"
				},
				body: JSON.stringify({ token: currentToken })
			});

			if (!response.ok) {
				throw new Error("Error al confirmar la cuenta");
			}

			alert("Cuenta confirmada exitosamente.");
		} catch (error) {
			console.error("Error en la confirmación:", error);
			alert("Ocurrió un error al confirmar la cuenta.");
		}
	};

	return (
            <div class="confirmation-wrapper">
                <div class="confirmation-page">
                    <h1>Confirmación de Cuenta</h1>
                    <p>Por favor, confirma tu cuenta haciendo clic en el botón.</p>
                    <button class="confirm-button" onClick={handleConfirm}>
                        Confirmar Cuenta
                    </button>
                </div>
            </div>
    
	);
};

export default ConfirmationPage;
