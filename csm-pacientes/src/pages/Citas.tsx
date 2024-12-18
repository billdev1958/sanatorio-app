import Calendario from '../components/Calendario'; // Importa el componente del calendario
import HorarioSelector from '../components/HorarioSelector'; // Importa el componente del selector de horario

const Citas = () => {
  return (
    <div class="citas-container">
      <div class="cita-card">
        <h1>Selecciona tu Cita</h1>
        
        <div class="cita-sections">
          <div class="form-section">
          <h2>Selecciona al beneficiario</h2>
            <select>
              <option value="">Beneficiarios --</option>
              <option value="1">Billy Rivera Salinas</option>
              <option value="2">Ana Patricia Sanchez Jacobo</option>
              <option value="3">Leonardo Jesus Rivera Sanchez</option>
            </select>
            <h2>Selecciona Servicio</h2>
            <select>
              <option value="">Servicios --</option>
              <option value="medicina-general">Medicina General</option>
              <option value="pediatria">Pediatría</option>
              <option value="dermatologia">Dermatología</option>
              <option value="ginecologia">Ginecología</option>
              <option value="cardiologia">Cardiología</option>
            </select>

            <h2>Motivo de la Consulta</h2>
            <textarea placeholder="Escribe el motivo de tu consulta" rows="4"></textarea>
          </div>

          <div class="calendario-section">
            <h2>Selecciona una Fecha</h2>
            <Calendario />
          </div>

          <div class="horario-section">
            <h2>Selecciona un Horario</h2>
            <HorarioSelector />

            <div class="action-buttons">
              <button class="confirm-button">Confirmar Cita</button>
              <button class="cancel-button">Cancelar</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Citas;
