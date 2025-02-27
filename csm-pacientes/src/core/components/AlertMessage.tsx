import { Component } from "solid-js";

interface AlertMessageProps {
  type?: "error" | "success";
  message: string;
  onClose?: () => void;
}

const AlertMessage: Component<AlertMessageProps> = (props) => {
  const handleAnimationEnd = () => {
    if (props.onClose) {
      props.onClose();
    }
  };

  return (
    <div class="alert-wrapper">
      <div
        class={`alert-message ${props.type === "error" ? "error" : "success"}`}
        onAnimationEnd={handleAnimationEnd}
      >
        <span>{props.message}</span>
        {props.onClose && (
          <button class="close-btn" onClick={props.onClose}>
            ×
          </button>
        )}
      </div>
    </div>
  );
};

export default AlertMessage;
