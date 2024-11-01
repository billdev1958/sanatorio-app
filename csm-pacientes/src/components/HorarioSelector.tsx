import { createSignal } from 'solid-js';

const HorarioSelector = () => {
  const [selectedTime, setSelectedTime] = createSignal<string | null>(null);

  // Generar horarios con intervalos de 35 minutos
  const generarHorarios = () => {
    const horarios = [
      { inicio: '08:00 AM', fin: '08:35 AM' },
      { inicio: '08:35 AM', fin: '09:10 AM' },
      { inicio: '09:10 AM', fin: '09:45 AM' },
      { inicio: '09:45 AM', fin: '10:20 AM' },
      { inicio: '10:20 AM', fin: '10:55 AM' },
      { inicio: '10:55 AM', fin: '11:30 AM' },
      { inicio: '11:30 AM', fin: '12:05 PM' },
      { inicio: '12:05 PM', fin: '12:40 PM' },
      { inicio: '12:40 PM', fin: '01:15 PM' },
      { inicio: '01:15 PM', fin: '01:50 PM' },
      { inicio: '01:50 PM', fin: '02:25 PM' },
      { inicio: '02:25 PM', fin: '03:00 PM' },
      { inicio: '03:00 PM', fin: '03:35 PM' },
      { inicio: '03:35 PM', fin: '04:10 PM' },
      { inicio: '04:10 PM', fin: '04:45 PM' }
    ];

    return horarios;
  };

  const horarios = generarHorarios();

  return (
    <div class="horario-container">
      <div class="horarios-grid">
        {horarios.map(({ inicio, fin }) => (
          <div
            class={`horario-card ${selectedTime() === `${inicio} - ${fin}` ? 'selected' : ''}`} 
            onClick={() => setSelectedTime(`${inicio} - ${fin}`)}
          >
            {`${inicio} - ${fin}`}
          </div>
        ))}
      </div>
    </div>
  );
};

export default HorarioSelector;
