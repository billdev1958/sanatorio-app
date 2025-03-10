import { createSignal, createEffect, Show, Switch, Match } from "solid-js";
import { useNavigate } from "@solidjs/router";
import logoCMS from '../assets/logo_cms.png'; 
import FormInput from '../components/FormInput'; 
import { useLoginService } from '../services/LoginService'; 
import { LoginUser } from '../models/Login&Registers'; 
import NavBar from '../components/NavBar';
import AlertMessage from '../core/components/AlertMessage';
import { useMessage } from '../core/domain/messageProvider';
import ModalEmailForward from "../components/modalEmailForward";
import Modal from "../components/Modal";

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
  const [isActive, setIsActive] = createSignal<boolean | null>(null);
  const navigate = useNavigate();
  const { successMessage, setSuccessMessage } = useMessage();

  const handleSubmit = async (e: Event) => {
    e.preventDefault(); 

    const user: LoginUser = {
      email: email(),
      password: password(),
    };

    const response = await login(user);

    if (loginError()) {
      setModalProps({
        type: "error",
        message: loginError() ?? "Ocurrió un error inesperado.",
        email: "",
        onClose: () => setShowModal(false),
      });
      setShowModal(true);
      return;
    }

    if (response && response.status === "success") {
      setIsActive(response.data.isActive);
      if (response.data.isActive === false) {
        setModalProps({
          type: "error",
          message: "Tu cuenta no está verificada. Por favor, reenvía el código de verificación.",
          email: email(),
          onClose: () => setShowModal(false),
        });
        setShowModal(true);
      }
    }
  };

  createEffect(() => {
    if (token() && isActive() === true) {
      console.log("Login exitoso, navegando a la home...");
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
        <Switch>
          <Match when={modalProps().email !== ""}>
            <ModalEmailForward {...modalProps()} />
          </Match>
          <Match when={true}>
            <Modal 
              type={modalProps().type} 
              message={modalProps().message} 
              onClose={modalProps().onClose} 
            />
          </Match>
        </Switch>
      </Show>
    </>
  );
};

export default Login;
