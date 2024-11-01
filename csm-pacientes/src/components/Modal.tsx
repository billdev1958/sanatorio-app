
interface ModalProps {
  type: "success" | "error";
  message: string;
  onClose: () => void;
}

function Modal(props: ModalProps) {
  return (
    <div class="modal-container" onClick={props.onClose}>
      <div
        class={`modal-content ${props.type === "success" ? "modal-success" : "modal-error"}`}
        onClick={(e) => e.stopPropagation()}
      >
        <div class="modal-header">
          <h3 class="modal-title">{props.type === "success" ? "Éxito" : "Error"}</h3>
          <button class="modal-close-btn" onClick={props.onClose}>
            &times;
          </button>
        </div>
        <div class="modal-body">
          <p>{props.message}</p>
        </div>
        <div class="modal-footer">
          <button class="btn-action btn-close" onClick={props.onClose}>
            Cerrar
          </button>
        </div>
      </div>
    </div>
  );
}



export default Modal;