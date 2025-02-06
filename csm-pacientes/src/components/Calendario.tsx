import { createEffect, onCleanup, onMount } from 'solid-js';
import flatpickr from 'flatpickr';
import 'flatpickr/dist/flatpickr.min.css'; // Importamos los estilos de flatpickr

// Actualizamos la interfaz para que acepte la propiedad "selectedDate"
interface CalendarioProps {
  onDateChange: (selectedDate: string) => void; // Callback para enviar la fecha seleccionada en UTC
  selectedDate?: string; // Fecha seleccionada por defecto (en formato ISO)
}

const Calendario = (props: CalendarioProps) => {
  let calendarRef: HTMLInputElement | null = null;
  let calendarInstance: flatpickr.Instance; // Guardaremos la instancia de flatpickr

  onMount(() => {
    if (calendarRef) {
      calendarInstance = flatpickr(calendarRef, {
        inline: true, // Hace que el calendario siempre esté visible
        dateFormat: 'Y-m-d', // Formato de la fecha
        minDate: 'today', // Deshabilita los días anteriores a hoy
        defaultDate: props.selectedDate ? new Date(props.selectedDate) : new Date(), // Fecha por defecto
        onChange: (selectedDates) => {
          if (selectedDates.length > 0) {
            // Convertir la fecha seleccionada a formato ISO UTC
            const utcDate = new Date(selectedDates[0]).toISOString();
            props.onDateChange(utcDate); // Llamamos al callback con la fecha en UTC
          }
        },
      });

      onCleanup(() => {
        calendarInstance.destroy(); // Limpieza cuando el componente se desmonta
      });
    }
  });

  // Este efecto actualiza la fecha de la instancia si cambia la propiedad "selectedDate"
  createEffect(() => {
    if (calendarInstance && props.selectedDate) {
      calendarInstance.setDate(new Date(props.selectedDate), false);
    }
  });

  return (
    <div>
      <input ref={(el) => (calendarRef = el)} type="text" style={{ display: 'none' }} />
    </div>
  );
};

export default Calendario;
