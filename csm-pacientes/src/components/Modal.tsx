import { createSignal, createEffect, onCleanup } from "solid-js";

interface ModalProps {
  type: "success" | "error";
  message: string;
  onClose: () => void;
}

function Modal(props: ModalProps) {
  // Contador (60 segundos) para deshabilitar el botón tras el reenvío
  const [counter, setCounter] = createSignal(60);
  // Bandera para saber si el contador está activo (cuando se presiona el botón de reenvío)
  const [isCounting, setIsCounting] = createSignal(false);
  // Bandera para mostrar el input solo cuando se haya presionado "Reenviar código"
  const [showInput, setShowInput] = createSignal(false);
  // Señal para almacenar el código de verificación ingresado por el usuario
  const [verificationCode, setVerificationCode] = createSignal("");

  // Efecto para decrementar el contador cada segundo cuando esté activo
  createEffect(() => {
    let timer: number;
    if (isCounting()) {
      timer = setInterval(() => {
        setCounter((prev) => {
          if (prev <= 1) {
            // Cuando el contador llega a 0, se reinicia y se habilita el botón nuevamente
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

  // Handler para reenviar el código; inicia el contador y muestra el input al ejecutarse
  const handleResend = () => {
    if (!isCounting()) {
      // Aquí puedes agregar la lógica para reenviar el código, por ejemplo llamar a una API.
      console.log("Reenviando código...");
      setIsCounting(true);
      setShowInput(true);
    }
  };

  // Handler para confirmar el código ingresado
  const handleConfirm = () => {
    console.log("Código ingresado:", verificationCode());
    // Agrega validaciones o envía el código al backend según sea necesario
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

export default Modal;
