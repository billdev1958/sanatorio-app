import { createSignal, createEffect, Show } from "solid-js";
import { useNavigate } from "@solidjs/router";
import logoCMS from '../assets/logo_cms.png'; 
import FormInput from '../components/FormInput'; 
import { useLoginService } from '../services/LoginService'; 
import { LoginUser } from '../models/Login&Registers'; 
import NavBar from '../components/NavBar';
import AlertMessage from '../core/components/AlertMessage';
import { useMessage } from '../core/domain/messageProvider';
import ModalEmailForward from "../components/modalEmailForward";

const Login = () => {
  const { login, loginError, isLoggingIn, token } = useLoginService(); 
  const [email, setEmail] = createSignal<string>(""); 
  const [password, setPassword] = createSignal<string>(""); 
  const [showModal, setShowModal] = createSignal(false);
  const [modalProps, setModalProps] = createSignal({
    type: "error" as "success" | "error",
    message: "",
    email: "",
    onClose: () => setShowModal(false),
  });
  const navigate = useNavigate();
  const { successMessage, setSuccessMessage } = useMessage();

  const handleSubmit = async (e: Event) => {
    e.preventDefault(); 

    const user: LoginUser = {
      email: email(),
      password: password(),
    };

    await login(user);

    if (loginError()) {
      setModalProps({
        type: "error",
        message: loginError() ?? "Ocurrió un error inesperado.",
        email: email(),
        onClose: () => setShowModal(false),
      });
      setShowModal(true);
    }
  };

  createEffect(() => {
    if (token()) {
      console.log("Login successful, navigating to home page...");
      navigate("/", { replace: true });
    }
  });

  return (
    <>
      <Show when={successMessage()}>
        <AlertMessage 
          type="success" 
          message={successMessage()!} 
          onClose={() => setSuccessMessage(null)} 
        />
      </Show>

      <NavBar toggleMenu={() => {}} />

      <div class="form-container">
        <div class="form-card">
          <div class="form-logo">
            <img src={logoCMS} alt="Logo CMS" />
          </div>

          <h2>Iniciar Sesión</h2>
          <form class="form" onSubmit={handleSubmit}>
            <FormInput
              type="email"
              name="email"
              placeholder="Correo electrónico"
              required={true}
              value={email()}
              onInput={(e: InputEvent) => setEmail((e.target as HTMLInputElement).value)} 
            />
            <FormInput
              type="password"
              name="password"
              placeholder="Contraseña"
              required={true}
              value={password()} 
              onInput={(e: InputEvent) => setPassword((e.target as HTMLInputElement).value)} 
            />
            <button type="submit" class="form-button" disabled={isLoggingIn()}>
              {isLoggingIn() ? "Entrando..." : "Entrar"}
            </button>
          </form>

          <div class="form-links">
            <p>Conoce nuestro <a href="#privacy">Aviso de Privacidad</a></p>
            <p>¿Eres nuevo? <a href="register">Regístrate</a></p>
            <p>¿Olvidaste tu contraseña? <a href="#recover">Recupérala</a></p>
          </div>
        </div>
      </div>

      <Show when={showModal()}>
        <ModalEmailForward {...modalProps()} />
      </Show>
    </>
  );
};

export default Login;
