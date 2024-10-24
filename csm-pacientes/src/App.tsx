import { createSignal, Show } from 'solid-js';
import NavBar from './components/NavBar';
import LateralMenu from './components/LateralMenu';
import { JSX } from 'solid-js';
import { useAuth } from './services/AuthContext'; 
import { useLocation } from '@solidjs/router'; 

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
      <NavBar toggleMenu={toggleMenu} />

      <div class="main-layout">
        <Show when={auth.isAuthenticated() && !isPublicRoute}>
          <LateralMenu open={menuOpen()} toggleMenu={toggleMenu} />
        </Show>

        <div class={`content-area ${auth.isAuthenticated() && menuOpen() && !isPublicRoute ? 'menu-open' : 'menu-closed'}`}>
          {props.children}
        </div>
      </div>
    </div>
  );
};

export default App;
