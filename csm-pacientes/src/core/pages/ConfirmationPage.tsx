import { Component, createSignal, onMount } from "solid-js";
import { useParams, useNavigate } from "@solidjs/router";
import Modal from "../../components/Modal";

const ConfirmationPage: Component = () => {
	const params = useParams();
	const navigate = useNavigate();
	const [token, setToken] = createSignal<string | null>(null);
	const [modalProps, setModalProps] = createSignal<{ type: "success" | "error"; message: string } | null>(null);

	onMount(() => {
		if (params.token) {
			setToken(params.token);
			localStorage.setItem("confirmationToken", params.token);
		}
	});

	const handleConfirm = async () => {
		const currentToken = token();
		if (!currentToken) {
			setModalProps({ type: "error", message: "No se encontró un token válido en la URL." });
			return;
		}

		try {
			const response = await fetch("https://api.ax01.dev/v1/confirmation", {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify({ token: currentToken })
			});

			const result = await response.json();

			if (!response.ok) {
				throw new Error(result.message || "Error al confirmar la cuenta");
			}

			setModalProps({ type: "success", message: "Cuenta confirmada exitosamente. Redirigiendo al login..." });

			setTimeout(() => {
				navigate("/login");
			}, 3000);

		} catch (error: any) {
			console.error("Error en la confirmación:", error);
			setModalProps({ type: "error", message: error.message || "Ocurrió un error al confirmar la cuenta." });
		}
	};

	return (
		<div class="confirmation-wrapper">
			<div class="confirmation-page">
				<h1>Confirmación de Cuenta</h1>
				<p>Por favor, confirma tu cuenta haciendo clic en el botón.</p>

				{modalProps() && (
					<Modal
						type={modalProps()!.type}
						message={modalProps()!.message}
						onClose={() => setModalProps(null)}
					/>
				)}

				<button class="confirm-button" onClick={handleConfirm}>Confirmar Cuenta</button>
			</div>
		</div>
	);
};

export default ConfirmationPage;
