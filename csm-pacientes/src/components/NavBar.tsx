import logo1 from '../assets/logo.png';
import menuIcon from '../assets/Menu_Hamburger.svg'; // Añade el icono de menú hamburguesa
import '../styles/index.scss';

const NavBar = (props: { toggleMenu: () => void }) => {
  return (
    <>
      <nav class="navbar">
        {/* Botón de menú hamburguesa en el NavBar */}
        <div class="menu-hamburguesa" onClick={props.toggleMenu}>
          <img src={menuIcon} alt="Menú" />
        </div>

        <div class="navbar-logos">
          <img src={logo1} alt="Logo 1" class="navbar-logo-img" />
        </div>

        <div class="navbar-menu">
          <a href="#home">Home</a>
          <a href="#about">About</a>
          <a href="#contact">Contact</a>
        </div>
        <button class="navbar-button">Registrate</button>
      </nav>
      <div class="marco-dorado"></div>
    </>
  );
};

export default NavBar;
