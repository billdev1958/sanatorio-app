import { useState } from 'react';
import IonIcon from "@reacticons/ionicons";
import { Link, useNavigate } from 'react-router-dom';

/*const scrollDown = () => {
  const scrollHeight = document.body.scrollHeight;

  window.scrollTo({
      top: scrollHeight,
      behavior: "smooth"
  });
}*/

const NavBar = () => {
  const [isMobileActive, setIsMobileActive] = useState(false);
  // const [isSubMenuActive, setIsSubMenuActive] = useState(false);
  const navigate = useNavigate(); // Hook para navegar a diferentes rutas

  const closeMobileMenu = () => {
    setIsMobileActive(false);
  };

  const handleLogout = () => {
    // Eliminar el token JWT del localStorage
    localStorage.removeItem('jwtToken');
    
    // Redirigir al usuario a la página de login
    navigate('/login');

    // Cerrar el menú móvil si está abierto
    closeMobileMenu();
  };

return (
  <nav className="navbar">
    <div className="navbar-logo">AX01</div>
    <div className="navbar-menu" onClick={() => setIsMobileActive(!isMobileActive)}>
      <IonIcon className="icon-edit" name="menu-outline" />
    </div>
    <ul className={`navbar-links ${isMobileActive ? 'mobile-active' : ''}`}>
      <li className="navbar-item" onClick={closeMobileMenu}><Link to="/">Inicio</Link></li>

      {/* Menú de Registros comentado para deshabilitarlo */}
      {/*
      <li className="navbar-item">
        <a href="#" onClick={(e) => { e.preventDefault(); setIsSubMenuActive(!isSubMenuActive); }}>
          Registros
        </a>
        {isSubMenuActive && (
          <ul className="submenu">
            <li onClick={closeMobileMenu}>
              <Link to="/register/user">Super usuarios y pacientes</Link>
            </li>
            <li onClick={closeMobileMenu}>
              <Link to="/register/doctor">Doctores</Link>
            </li>
            <li onClick={closeMobileMenu}>
              <Link to="/servicios/investigacion">Pacientes</Link>
            </li>
          </ul>
        )}
      </li>
      */}
            <li className="navbar-item">
              <Link to="/control">Control Hospital</Link>
            </li>
      <li className="navbar-item" onClick={handleLogout}><a href="#">Cerrar sesión</a></li>
    </ul>
  </nav>
);

}

export default NavBar;
