import { createSignal, createEffect } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import Calendario from '../../components/Calendario';
import { getParamsForAppointment, getOfficeSchedules } from '../Services/CatalogServices';
import { Services, Shift, SchedulesAppointmentRequest, OfficeScheduleResponse } from '../Models/Catalogs';
import { useAuth } from '../../services/AuthContext';

const Citas = () => {
  const [shift, setShift] = createSignal<Shift | null>(null);
  const [service, setService] = createSignal<Services | null>(null);
  const [fullDate, setFullDate] = createSignal<string | null>(null);
  const [schedules, setSchedules] = createSignal<OfficeScheduleResponse[]>([]);
  const [selectedSchedule, setSelectedSchedule] = createSignal<OfficeScheduleResponse | null>(null);
  const [params, setParams] = createSignal<{ services: Services[]; shifts: Shift[] } | null>(null);
  const [scheduleError, setScheduleError] = createSignal<string | null>(null);
  const [loading, setLoading] = createSignal(false);

  const navigate = useNavigate();
  const { token } = useAuth();

  createEffect(async () => {
    try {
      const response = await getParamsForAppointment(token() ?? undefined);
      if (response.data) {
        setParams(response.data);
      } else {
        throw new Error('No se encontraron servicios o turnos disponibles.');
      }
    } catch (error: any) {
      console.error('Error al obtener servicios y turnos:', error);
      setScheduleError(error.message || 'Error al obtener los parámetros.');
    }
  });

  const handleServiceChange = (e: InputEvent) => {
    const selectedServiceId = parseInt((e.target as HTMLSelectElement).value);
    const currentParams = params();
    if (currentParams) {
      const selectedService = currentParams.services.find((service) => service.id === selectedServiceId);
      setService(selectedService || null);
    }
  };

  const handleShiftChange = (e: InputEvent) => {
    const selectedShiftId = parseInt((e.target as HTMLSelectElement).value);
    const currentParams = params();
    if (currentParams) {
      const selectedShift = currentParams.shifts.find((shift) => shift.id === selectedShiftId);
      setShift(selectedShift || null);
    }
  };

  createEffect(async () => {
    if (!service() || !shift() || !fullDate()) {
      return;
    }

    const appointmentData: SchedulesAppointmentRequest = {
      service: service()!.id,
      shift: shift()!.id,
      appointmentDate: fullDate()!,
    };

    try {
      setLoading(true);
      setScheduleError(null); // Reiniciar errores
      const response = await getOfficeSchedules(appointmentData, token() ?? undefined);
      if (response.data && Array.isArray(response.data)) {
        setSchedules(response.data);
      } else {
        setSchedules([]);
        setScheduleError('No se encontraron horarios disponibles.');
      }
    } catch (error: any) {
      console.error('Error al obtener horarios:', error);
      setSchedules([]);
      setScheduleError(error.message || 'Error al obtener los horarios.');
    } finally {
      setLoading(false);
    }
  });

  return (
    <div class="citas-container">
      <div class="cita-card">
        <h1>Selecciona tu Cita</h1>

        <div class="cita-sections">
          {/* Columna izquierda: Formulario y calendario */}
          <div class="left-section">
            <div class="form-section">
              <h2>Selecciona Servicio</h2>
              <select id="service" required onInput={handleServiceChange}>
                <option value="">Servicios --</option>
                {params()?.services.map((service) => (
                  <option value={service.id}>{service.name}</option>
                ))}
              </select>

              <h2>Selecciona Turno</h2>
              <select id="turno" required onInput={handleShiftChange}>
                <option value="">Turno --</option>
                {params()?.shifts.map((shift) => (
                  <option value={shift.id}>{shift.name}</option>
                ))}
              </select>
            </div>

            <div class="calendario-section">
              <h2>Selecciona una Fecha</h2>
              <Calendario
                onDateChange={(selectedDate: Date) =>
                  setFullDate(selectedDate.toISOString().split('T')[0])
                }
              />
              {fullDate() && <p class="selected-date">Fecha seleccionada: {fullDate()}</p>}
            </div>
          </div>

          {/* Columna derecha: Horarios */}
          <div class="right-section">
            <div class="schedules-section">
              <h2>Horarios Disponibles</h2>
              {loading() && <div class="loading-message">Cargando horarios...</div>}

              {!loading() && scheduleError() && (
                <div class="error-message">
                  <p>{scheduleError()}</p>
                </div>
              )}

              {!loading() && schedules().length > 0 && (
                <div class="schedules-grid">
                  {schedules().map((schedule) => (
                    <div
                      class={`schedule-card ${
                        selectedSchedule() && selectedSchedule()!.id === schedule.id ? 'selected' : ''
                      }`}
                      onClick={() => setSelectedSchedule(schedule)}
                    >
                      <p class="time-label">Inicio:</p>
                      <p class="time-value">{schedule.timeStart}</p>
                      <p class="time-label">Fin:</p>
                      <p class="time-value">{schedule.timeEnd}</p>
                    </div>
                  ))}
                </div>
              )}
            </div>

            {/* Botones debajo de los horarios */}
            <div class="action-buttons">
              <button
                type="button"
                class="confirm-button"
                onClick={() => {
                  if (selectedSchedule()) {
                    console.log('Horario seleccionado:', selectedSchedule());
                  } else {
                    alert('Debe seleccionar un horario antes de continuar.');
                  }
                }}
              >
                Confirmar Horario
              </button>
              <button type="button" class="cancel-button" onClick={() => navigate('/')}>
                Cancelar
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Citas;
