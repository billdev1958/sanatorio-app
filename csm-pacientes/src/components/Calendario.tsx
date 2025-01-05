import { onCleanup, onMount } from 'solid-js';
import flatpickr from 'flatpickr';
import 'flatpickr/dist/flatpickr.min.css'; // Importamos los estilos de flatpickr

// Definimos las props que acepta el componente
interface CalendarioProps {
  onDateChange: (selectedDate: Date) => void; // Callback para enviar la fecha seleccionada
}

const Calendario = (props: CalendarioProps) => {
  let calendarRef: HTMLInputElement | null = null; // Referencia al input del calendario

  onMount(() => {
    if (calendarRef) { // Verificamos que el ref no sea null
      const calendar = flatpickr(calendarRef, {
        inline: true, // Hace que el calendario siempre esté visible
        dateFormat: 'Y-m-d', // Formato de la fecha
        minDate: 'today', // Deshabilita los días anteriores a hoy
        defaultDate: new Date(), // Fecha por defecto (hoy)
        onChange: (selectedDates) => {
          if (selectedDates.length > 0) {
            props.onDateChange(selectedDates[0]); // Llamamos al callback con la fecha seleccionada
          }
        },
      });

      onCleanup(() => {
        calendar.destroy(); // Limpieza cuando el componente se desmonta
      });
    }
  });

  return (
    <div>
      <input ref={(el) => (calendarRef = el)} type="text" style={{ display: 'none' }} /> {/* Este input está oculto, pero es necesario para flatpickr */}
    </div>
  );
};

export default Calendario;
