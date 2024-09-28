import { createSignal } from 'solid-js';
import NavBar from './components/NavBar';
import LateralMenu from './components/LateralMenu';
import { JSX } from 'solid-js';

const App = (props: { children?: JSX.Element }) => { 
  const [menuOpen, setMenuOpen] = createSignal(false);

  const toggleMenu = () => {
    setMenuOpen(!menuOpen());
  };

  return (
    <div class="app-container">
      <NavBar toggleMenu={toggleMenu} />
      <div class="main-layout">
        <LateralMenu open={menuOpen()} toggleMenu={toggleMenu} />
        <div class={`content-area ${menuOpen() ? 'menu-open' : 'menu-closed'}`}>
          {props.children}
        </div>
      </div>
    </div>
  );
};

export default App;
