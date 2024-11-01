import { createSignal, createEffect, Show } from "solid-js";
import { useNavigate } from "@solidjs/router";
import logoCMS from '../assets/logo_cms.png'; 
import FormInput from '../components/FormInput'; 
import { useLoginService } from '../services/LoginService'; 
import { LoginUser } from '../models/Login&Registers'; 
import NavBar from '../components/NavBar'; // Importamos el NavBar para mostrarlo
import Modal from "../components/Modal"; // Importamos el modal

const Login = () => {
  // Utiliza el servicio de login
  const { login, loginError, isLoggingIn, token } = useLoginService(); 

  // Señales para el email y la contraseña
  const [email, setEmail] = createSignal<string>(""); 
  const [password, setPassword] = createSignal<string>(""); 

  // Señales para controlar el modal
  const [showModal, setShowModal] = createSignal(false);
  const [modalProps, setModalProps] = createSignal({
    type: "error" as "success" | "error",
    message: "",
    onClose: () => setShowModal(false),
  });

  // Hook para la navegación
  const navigate = useNavigate(); 

  // Manejador del evento de submit del formulario
  const handleSubmit = async (e: Event) => {
    e.preventDefault(); 

    const user: LoginUser = {
      email: email(),
      password: password(),
    };

    // Intenta loguear al usuario
    await login(user);

    // Si hay un error de login, mostramos el modal
    if (loginError()) {
      setModalProps({
        type: "error",
        message: loginError() ?? "Ocurrió un error inesperado.", // Aseguramos que message siempre sea un string
        onClose: () => setShowModal(false),
      });
      setShowModal(true);
    }
  };

  // Efecto para redirigir al usuario si hay un token presente
  createEffect(() => {
    if (token()) {
      console.log("Login successful, navigating to home page...");
      navigate("/", { replace: true });
    }
  });

  return (
    <>
      {/* Incluimos el NavBar siempre */}
      <NavBar toggleMenu={() => {}} /> {/* No necesitamos el menú hamburguesa en el login, por eso la función está vacía */}

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

      {/* Modal para mostrar errores o mensajes */}
      <Show when={showModal()}>
        <Modal {...modalProps()} />
      </Show>
    </>
  );
};

export default Login;
