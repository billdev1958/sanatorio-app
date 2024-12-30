import { createSignal, createEffect } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import Calendario from '../components/Calendario';
import HorarioSelector from '../components/HorarioSelector';
import { getScheduleAppointment } from '../services/RegisterService'; // Servicio ajustado
import { ScheduleAppointment, ScheduleData } from '../models/Horarios';
import { useAuth } from '../services/AuthContext'; // Para obtener el token

const Citas = () => {
  const [shift, setShift] = createSignal<number | null>(null); // Turno
  const [service, setService] = createSignal<number | null>(null); // Servicio
  const [day, setDay] = createSignal<number | null>(null); // Día seleccionado
  const [horarios, setHorarios] = createSignal<ScheduleData[]>([]); // Lista de horarios disponibles
  const [selectedHorario, setSelectedHorario] = createSignal<ScheduleData | null>(null); // Horario seleccionado
  const [scheduleError, setScheduleError] = createSignal<string | null>(null); // Error
  const navigate = useNavigate();
  const { token } = useAuth(); // Obtener el token del contexto de autenticación

  // Efecto reactivo: actualiza los horarios automáticamente al cambiar día, turno o servicio
  createEffect(async () => {
    // Validar que todos los campos requeridos estén completos
    if (!shift() || !service() || day() === null) {
      setScheduleError('Por favor, completa todos los campos antes de buscar los horarios.');
      return;
    }

    setScheduleError(null); // Limpiar errores previos
    try {
      const appointment: ScheduleAppointment = {
        shift: shift()!,
        service: service()!,
        day: day()!,
      };

      // Llamada al servicio para obtener los horarios
      const response = await getScheduleAppointment(appointment, token() ?? undefined);
      setHorarios(response.data); // Actualiza los horarios disponibles
    } catch (error: any) {
      setScheduleError(error.message);
      setHorarios([]); // Limpia los horarios si hay error
    }
  });

  return (
    <div class="citas-container">
      <div class="cita-card">
        <h1>Selecciona tu Cita</h1>
        <div class="cita-sections">
          {/* Sección del servicio */}
          <div class="form-section">
            <h2>Selecciona Servicio</h2>
            <select
              id="service"
              required
              onInput={(e: InputEvent) => setService(parseInt((e.target as HTMLSelectElement).value))}
            >
              <option value="">Servicios --</option>
              <option value="25">Medicina General</option>
              <option value="26">Pediatría</option>
              <option value="27">Dermatología</option>
            </select>

            <h2>Selecciona Turno</h2>
            <select
              id="turno"
              required
              onInput={(e: InputEvent) => setShift((e.target as HTMLSelectElement).value === 'matutino' ? 1 : 2)}
            >
              <option value="">Turno --</option>
              <option value="matutino">Matutino</option>
              <option value="vespertino">Vespertino</option>
            </select>
          </div>

          {/* Sección del calendario */}
          <div class="calendario-section">
            <h2>Selecciona una Fecha</h2>
            <Calendario onDateChange={(selectedDate: Date) => setDay(selectedDate.getDay())} />
            {day() !== null && <p>Día seleccionado: {day()}</p>}
          </div>

          {/* Sección del horario */}
          {horarios().length > 0 ? (
            <div class="horario-section">
              <h2>Selecciona un Horario</h2>
              <HorarioSelector
                horarios={horarios()}
                onHorarioSeleccionado={(horario) => setSelectedHorario(horario)}
              />
            </div>
          ) : (
            <div class="horario-section">
              <p>No hay horarios disponibles para los criterios seleccionados.</p>
            </div>
          )}
        </div>

        {/* Mensaje de error */}
        {scheduleError() && (
          <div class="error-message">
            <p>{scheduleError()}</p>
          </div>
        )}

        {/* Botones de acción */}
        <div class="action-buttons">
          <button
            type="button"
            class="confirm-button"
            onClick={() => {
              if (selectedHorario()) {
                console.log('Cita confirmada:', selectedHorario());
                navigate('/success', { replace: true });
              } else {
                setScheduleError('Por favor, selecciona un horario antes de confirmar.');
              }
            }}
          >
            Confirmar Cita
          </button>
          <button type="button" class="cancel-button" onClick={() => navigate('/')}>
            Cancelar
          </button>
        </div>
      </div>
    </div>
  );
};

export default Citas;
