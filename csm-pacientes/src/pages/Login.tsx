import logoCMS from '../assets/logo_cms.png'; // Asegúrate de importar correctamente el logo
import FormInput from '../components/FormInput'; // Importa el componente reutilizable

const Login = () => {
  return (
    <>
      <div class="form-container">
        <div class="form-card">
          <div class="form-logo">
            <img src={logoCMS} alt="Logo CMS" />
          </div>

          <h2>Iniciar Sesión</h2>
          <form class="form">
            <FormInput
              type="email"
              name="email"
              placeholder="Correo electrónico"
              required={true} 
            />
            <FormInput
              type="password"
              name="password"
              placeholder="Contraseña"
              required={true}
            />
            <button type="submit" class="form-button">Entrar</button>
          </form>

          <div class="form-links">
            <p>Conoce nuestro <a href="#privacy">Aviso de Privacidad</a></p>
            <p>¿Eres nuevo? <a href="#register">Regístrate</a></p>
            <p>¿Olvidaste tu contraseña? <a href="#recover">Recupérala</a></p>
          </div>
        </div>
      </div>
    </>
  );
};

export default Login;
