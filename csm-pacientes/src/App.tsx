import { createSignal, Show } from 'solid-js';
import NavBar from './components/NavBar';
import LateralMenu from './components/LateralMenu';
import { JSX } from 'solid-js';
import { useAuth } from './services/AuthContext'; // Importamos el hook de autenticación
import { useLocation } from '@solidjs/router'; // Usamos para obtener la ruta actual

const App = (props: { children?: JSX.Element }) => {
  const [menuOpen, setMenuOpen] = createSignal(false);
  const auth = useAuth(); // Obtenemos el contexto de autenticación
  const location = useLocation(); // Obtenemos la ubicación actual

  const toggleMenu = () => {
    setMenuOpen(!menuOpen());
  };

  // Determinamos si estamos en una ruta pública
  const isPublicRoute = location.pathname === "/login" || location.pathname === "/register";

  return (
    <div class="app-container">
      {/* Siempre mostramos el NavBar */}
      <NavBar toggleMenu={toggleMenu} />

      <div class="main-layout">
        {/* Mostramos el LateralMenu solo si el usuario está autenticado y no está en una ruta pública */}
        <Show when={auth.isAuthenticated() && !isPublicRoute}>
          <LateralMenu open={menuOpen()} toggleMenu={toggleMenu} />
        </Show>

        {/* Ajustamos la clase content-area dependiendo del estado del menú y la autenticación */}
        <div class={`content-area ${auth.isAuthenticated() && menuOpen() && !isPublicRoute ? 'menu-open' : 'menu-closed'}`}>
          {props.children}
        </div>
      </div>
    </div>
  );
};

export default App;
