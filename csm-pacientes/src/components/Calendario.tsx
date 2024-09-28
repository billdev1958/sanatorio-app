import { onCleanup, onMount } from 'solid-js';
import flatpickr from 'flatpickr';
import 'flatpickr/dist/flatpickr.min.css'; // Importamos los estilos de flatpickr

const Calendario = () => {
  let calendarRef: HTMLInputElement | null = null; // Referencia al input del calendario

  onMount(() => {
    if (calendarRef) { // Verificamos que el ref no sea null
      const calendar = flatpickr(calendarRef, {
        inline: true, // Hace que el calendario siempre esté visible
        dateFormat: 'Y-m-d', // Formato de la fecha
        defaultDate: new Date(), // Fecha por defecto (hoy)
        onChange: (selectedDates) => {
          console.log('Fecha seleccionada:', selectedDates);
        },
      });

      onCleanup(() => {
        calendar.destroy(); // Limpieza cuando el componente se desmonta
      });
    }
  });

  return (
    <div>
      <input ref={el => calendarRef = el} type="text" style={{ display: 'none' }} /> {/* Este input está oculto, pero es necesario para flatpickr */}
    </div>
  );
};

export default Calendario;
