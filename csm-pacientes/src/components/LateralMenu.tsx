import historiaClinicaIcon from '../assets/historiaclinica.png';
import cita from '../assets/citaicon.png';
import notaEvolucionIcon from '../assets/notaevolucion.png';
import recetaIcon from '../assets/receta.png';
import incapacidadIcon from '../assets/incapacidad.png';
import fileIcon from '../assets/file.png';
import laboratorioIcon from '../assets/laboratorio.png';
import logoutIcon from '../assets/logout.png';

const LateralMenu = (props: { open: boolean; toggleMenu: () => void }) => {
  return (
    <>
      {/* Menú lateral */}
      <div class={`lateral-menu ${props.open ? 'open' : ''}`}>
        <ul>
          <li>
            <img src={historiaClinicaIcon} alt="Historia Clínica" />
            <span>Historia Clínica</span>
          </li>
          <li>
            <img src={cita} alt="Agendar Cita" />
            <span>Agendar Cita</span>
          </li>
          <li>
            <img src={notaEvolucionIcon} alt="Nota de Evolución" />
            <span>Nota de Evolución</span>
          </li>
          <li>
            <img src={recetaIcon} alt="Receta" />
            <span>Receta</span>
          </li>
          <li>
            <img src={incapacidadIcon} alt="Incapacidad" />
            <span>Incapacidad</span>
          </li>
          <li>
            <img src={fileIcon} alt="Archivo" />
            <span>Archivo</span>
          </li>
          <li>
            <img src={laboratorioIcon} alt="Laboratorio" />
            <span>Laboratorio</span>
          </li>
          <li>
            <img src={logoutIcon} alt="Cerrar sesión" />
            <span>Cerrar sesión</span>
          </li>
        </ul>
      </div>
    </>
  );
};

export default LateralMenu;
