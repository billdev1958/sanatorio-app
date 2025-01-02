import { createSignal, createEffect } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import Calendario from '../../components/Calendario';
import { getParamsForAppointment } from '../Services/CatalogServices';
import { Services, Shift } from '../Models/Catalogs';
import { useAuth } from '../../services/AuthContext';

const Citas = () => {
  const [shift, setShift] = createSignal<Shift | null>(null);
  const [service, setService] = createSignal<Services | null>(null);
  const [day, setDay] = createSignal<number | null>(null);
  const [scheduleError, setScheduleError] = createSignal<string | null>(null);
  const [params, setParams] = createSignal<{ services: Services[]; shifts: Shift[] } | null>(null);
  const navigate = useNavigate();
  const { token } = useAuth();

  createEffect(async () => {
    try {
      const response = await getParamsForAppointment(token() ?? undefined);
      console.log('Respuesta de la API:', response);

      if (response.data) {
        setParams(response.data);
      } else {
        throw new Error('Datos no encontrados en la respuesta de la API');
      }
    } catch (error: any) {
      console.error('Error al obtener datos:', error);
      setScheduleError(error.message || 'Error al obtener los parámetros');
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

  return (
    <div class="citas-container">
      <div class="cita-card">
        <h1>Selecciona tu Cita</h1>
        <div class="cita-sections">
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
            <Calendario onDateChange={(selectedDate: Date) => setDay(selectedDate.getDay())} />
            {day() !== null && <p>Día seleccionado: {day()}</p>}
          </div>

          {scheduleError() && (
            <div class="error-message">
              <p>{scheduleError()}</p>
            </div>
          )}

          <div class="action-buttons">
            <button
              type="button"
              class="confirm-button"
              onClick={() =>
                console.log('Datos de la cita:', {
                  service: service(),
                  shift: shift(),
                  day: day(),
                })
              }
            >
              Confirmar Cita
            </button>
            <button type="button" class="cancel-button" onClick={() => navigate('/')}>
              Cancelar
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Citas;
