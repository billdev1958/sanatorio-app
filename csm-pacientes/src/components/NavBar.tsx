import { useAuth } from '../services/AuthContext'; // Importamos el hook de autenticación
import logo1 from '../assets/logo.png';
import menuIcon from '../assets/Menu_Hamburger.svg'; // Añade el icono de menú hamburguesa
import '../styles/index.scss';

const NavBar = (props: { toggleMenu: () => void }) => {
  const auth = useAuth(); // Obtenemos el contexto de autenticación

  return (
    <>
      <nav class="navbar">
        {/* Botón de menú hamburguesa solo si el usuario está autenticado */}
        {auth.isAuthenticated() && (
          <div class="menu-hamburguesa" onClick={props.toggleMenu}>
            <img src={menuIcon} alt="Menú" />
          </div>
        )}

        <div class="navbar-logos">
          <img src={logo1} alt="Logo 1" class="navbar-logo-img" />
          <div>
            <p>UNIVERSIDAD AUTONOMA DEL ESTADO DE MEXICO</p>
            <hr></hr>
            <p>UAEM</p>
            <p>CLINICA MULTIDISCIPLINARIA DE SALUD</p>
          </div>
        </div>

        {/*<button class="navbar-button">Registrate</button>*/}
      </nav>
      <div class="marco-dorado"></div>
    </>
  );
};

export default NavBar;
