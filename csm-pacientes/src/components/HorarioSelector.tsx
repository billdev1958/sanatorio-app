import { createSignal, For } from 'solid-js';
import { ScheduleData } from '../models/Horarios'; // Modelo para el horario de la API

interface HorarioSelectorProps {
  horarios: ScheduleData[]; // Lista de horarios disponibles
  onHorarioSeleccionado: (horario: ScheduleData) => void; // Callback para el horario seleccionado
}

const HorarioSelector = (props: HorarioSelectorProps) => {
  const [selectedHorario, setSelectedHorario] = createSignal<ScheduleData | null>(null);

  const seleccionarHorario = (horario: ScheduleData) => {
    setSelectedHorario(horario);
    props.onHorarioSeleccionado(horario); // Llama al callback con el horario seleccionado
  };

  return (
    <div class="horario-container">
      <div class="horarios-grid">
        <For each={props.horarios}>
          {(horario) => (
            <div
              class={`horario-card ${
                selectedHorario()?.ID === horario.ID ? 'selected' : ''
              }`}
              onClick={() => seleccionarHorario(horario)}
            >
              {`${horario.time_start} - ${horario.time_end}`}
            </div>
          )}
        </For>
      </div>
    </div>
  );
};

export default HorarioSelector;
