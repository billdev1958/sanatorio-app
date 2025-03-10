import { createSignal, createEffect, onCleanup } from "solid-js";
import { sendVerificationEmail, verifyCode } from "../services/RegisterService";

interface ModalProps {
  type: "success" | "error";
  message: string;
  email: string;
  onClose: () => void;
}

function ModalEmailForward(props: ModalProps) {
  const [counter, setCounter] = createSignal(60);
  const [isCounting, setIsCounting] = createSignal(false);
  const [showInput, setShowInput] = createSignal(false);
  const [verificationCode, setVerificationCode] = createSignal("");

  createEffect(() => {
    let timer: number;
    if (isCounting()) {
      timer = setInterval(() => {
        setCounter((prev) => {
          if (prev <= 1) {
            setIsCounting(false);
            clearInterval(timer);
            return 60;
          }
          return prev - 1;
        });
      }, 1000);
    }
    onCleanup(() => timer && clearInterval(timer));
  });

  const handleResend = async () => {
    if (!isCounting()) {
      try {
        console.log("Reenviando código...");
        const result = await sendVerificationEmail(props.email);
        console.log("Código de verificación enviado:", result);
        setIsCounting(true);
        setShowInput(true);
      } catch (error: any) {
        console.error("Error al reenviar código:", error);
      }
    }
  };

  const handleConfirm = async () => {
    try {
      const confirmationData = { email: props.email, code: verificationCode() };
      const result = await verifyCode(confirmationData);
      console.log("Código verificado:", result);
      props.onClose();
    } catch (error: any) {
      console.error("Error al verificar el código:", error);
    }
  };

  return (
    <div class="modal-container" onClick={props.onClose}>
      <div
        class={`modal-content ${props.type === "success" ? "modal-success" : "modal-error"}`}
        onClick={(e) => e.stopPropagation()}
      >
        <div class="modal-header">
          <h3 class="modal-title">
            {props.type === "success" ? "Éxito" : "Error"}
          </h3>
          <button class="modal-close-btn" onClick={props.onClose}>
            &times;
          </button>
        </div>
        <div class="modal-body">
          <p>{props.message}</p>
          {showInput() && (
            <input
              type="text"
              placeholder="Ingresa el código de verificación"
              value={verificationCode()}
              onInput={(e) => setVerificationCode(e.currentTarget.value)}
            />
          )}
        </div>
        <div class="modal-footer">
          <button class="btn-action btn-cancel" onClick={props.onClose}>
            Cancelar
          </button>
          <button class="btn-action btn-confirm" onClick={handleConfirm}>
            Confirmar
          </button>
          <button
            class="btn-action btn-resend"
            onClick={handleResend}
            disabled={isCounting()}
          >
            {isCounting() ? `Reenviar código (${counter()}s)` : "Reenviar código"}
          </button>
        </div>
      </div>
    </div>
  );
}

export default ModalEmailForward;
