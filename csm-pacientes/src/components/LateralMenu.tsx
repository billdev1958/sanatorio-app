import home from '../assets/home-outline.svg';

import historiaClinicaIcon from '../assets/historiaclinica.png';
import cita from '../assets/citaicon.png';
import notaEvolucionIcon from '../assets/notaevolucion.png';
import recetaIcon from '../assets/receta.png';
import incapacidadIcon from '../assets/incapacidad.png';
import fileIcon from '../assets/file.png';
import laboratorioIcon from '../assets/laboratorio.png';
import logoutIcon from '../assets/logout.png';
import { A, useNavigate } from "@solidjs/router";
import { useLoginService } from '../services/LoginService';

const LateralMenu = (props: { open: boolean; toggleMenu: () => void }) => {
  const { logout } = useLoginService(); // Traemos el método logout del servicio de login
  const navigate = useNavigate();

  // Función para manejar el logout
  const handleLogout = () => {
    logout(); // Llama a la función de logout del servicio
    navigate("/login", { replace: true }); // Redirige al usuario a la página de login
  };

  return (
    <>
      {/* Menú lateral */}
      <div class={`lateral-menu ${props.open ? 'open' : ''}`}>
        <ul>
        <li>
            <A href="/" activeClass="active" class="menu-link">
              <img src={home} alt="inicio" />
              <span>Inicio</span>
            </A>
          </li>
          <li>
            <A href="/medicalhistory" activeClass="active" class="menu-link">
              <img src={historiaClinicaIcon} alt="Historia Clínica" />
              <span>Historia Clínica</span>
            </A>
          </li>
          <li>
            <A href="/citas" activeClass="active" class="menu-link">
              <img src={cita} alt="Agendar Cita" />
              <span>Agendar Cita</span>
            </A>
          </li>
          <li>
            <A href="/nota-evolucion" activeClass="active" class="menu-link">
              <img src={notaEvolucionIcon} alt="Nota de Evolución" />
              <span>Nota de Evolución</span>
            </A>
          </li>
          <li>
            <A href="/receta" activeClass="active" class="menu-link">
              <img src={recetaIcon} alt="Receta" />
              <span>Receta</span>
            </A>
          </li>
          <li>
            <A href="/incapacidad" activeClass="active" class="menu-link">
              <img src={incapacidadIcon} alt="Incapacidad" />
              <span>Incapacidad</span>
            </A>
          </li>
          <li>
            <A href="/archivo" activeClass="active" class="menu-link">
              <img src={fileIcon} alt="Archivo" />
              <span>Archivo</span>
            </A>
          </li>
          <li>
            <A href="/laboratorio" activeClass="active" class="menu-link">
              <img src={laboratorioIcon} alt="Laboratorio" />
              <span>Laboratorio</span>
            </A>
          </li>
          <li>
            {/* Utilizamos un botón para manejar el logout */}
            <button onClick={handleLogout} class="menu-link logout-button">
              <img src={logoutIcon} alt="Cerrar sesión" />
              <span>Cerrar sesión</span>
            </button>
          </li>
        </ul>
      </div>
    </>
  );
};

export default LateralMenu;
