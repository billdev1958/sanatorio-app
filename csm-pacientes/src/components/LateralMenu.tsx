import historiaClinicaIcon from '../assets/historiaclinica.png';
import cita from '../assets/citaicon.png';
import notaEvolucionIcon from '../assets/notaevolucion.png';
import recetaIcon from '../assets/receta.png';
import incapacidadIcon from '../assets/incapacidad.png';
import fileIcon from '../assets/file.png';
import laboratorioIcon from '../assets/laboratorio.png';
import logoutIcon from '../assets/logout.png';
import { A } from "@solidjs/router";

const LateralMenu = (props: { open: boolean; toggleMenu: () => void }) => {
  return (
    <>
      {/* Menú lateral */}
      <div class={`lateral-menu ${props.open ? 'open' : ''}`}>
        <ul>
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
            <A href="/logout" activeClass="active" class="menu-link">
              <img src={logoutIcon} alt="Cerrar sesión" />
              <span>Cerrar sesión</span>
            </A>
          </li>
        </ul>
      </div>
    </>
  );
};

export default LateralMenu;
