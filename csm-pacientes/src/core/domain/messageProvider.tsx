import { createContext, useContext, ParentComponent, createSignal } from "solid-js";

interface MessageContextProps {
  successMessage: () => string | null;
  setSuccessMessage: (msg: string | null) => void;
  errorMessage: () => string | null;
  setErrorMessage: (msg: string | null) => void;
}

const MessageContext = createContext<MessageContextProps>();

export const MessageProvider: ParentComponent = (props) => {
  const [successMessage, setSuccessMessage] = createSignal<string | null>(null);
  const [errorMessage, setErrorMessage] = createSignal<string | null>(null);

  return (
    <MessageContext.Provider
      value={{
        successMessage,
        setSuccessMessage,
        errorMessage,
        setErrorMessage,
      }}
    >
      {props.children}
    </MessageContext.Provider>
  );
};

export const useMessage = () => {
  const context = useContext(MessageContext);
  if (!context) {
    throw new Error("useMessage must be used within a MessageProvider");
  }
  return context;
};
