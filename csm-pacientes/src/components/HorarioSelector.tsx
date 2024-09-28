import { createSignal } from 'solid-js';

const HorarioSelector = () => {
  const [selectedTime, setSelectedTime] = createSignal<string | null>(null);

  // Horarios predefinidos
  const horarios = [
    '08:00 AM', '09:00 AM', '10:00 AM', 
    '11:00 AM', '12:00 PM', '01:00 PM', 
    '02:00 PM', '03:00 PM', '04:00 PM'
  ];

  return (
    <div class="horario-container">
      <div class="horarios-grid">
        {horarios.map((hora) => (
          <div
            class={`horario-card ${selectedTime() === hora ? 'selected' : ''}`} 
            onClick={() => setSelectedTime(hora)}
          >
            {hora}
          </div>
        ))}
      </div>
    </div>
  );
};

export default HorarioSelector;
