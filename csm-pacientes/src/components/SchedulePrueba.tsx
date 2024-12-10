import { createSignal, onMount } from "solid-js";
import flatpickr from "flatpickr";
import "flatpickr/dist/flatpickr.css";

// Definición de tipos
interface Office {
  id: number;
  name: string;
}

interface Doctor {
  id: string;
  firstName: string;
  lastName1: string;
  lastName2: string;
}

interface Service {
  id: number;
  name: string;
}

interface Day {
  id: number;
  name: string;
}

interface FormData {
  officeId: number;
  doctorId: string;
  serviceId: number;
  day: number;
  timeStart: string;
  timeEnd: string;
}

interface RegisterOfficeScheduleProps {
  onCancel: () => void;
}

const RegisterOfficeSchedule = (props: RegisterOfficeScheduleProps) => {
  // Definición de señales tipadas
  const [offices, setOffices] = createSignal<Office[]>([
    { id: 1, name: "Consultorio 1" },
    { id: 2, name: "Consultorio 2" },
    { id: 3, name: "Consultorio 3" },
  ]);
  const [doctors, setDoctors] = createSignal<Doctor[]>([
    { id: "doc1", firstName: "Juan", lastName1: "Pérez", lastName2: "Lopez" },
    { id: "doc2", firstName: "Ana", lastName1: "Gómez", lastName2: "Martínez" },
  ]);
  const [services, setServices] = createSignal<Service[]>([
    { id: 1, name: "Audiometría" },
    { id: 2, name: "Acupuntura" },
  ]);
  const [days, setDays] = createSignal<Day[]>([
    { id: 1, name: "Lunes" },
    { id: 2, name: "Martes" },
    { id: 3, name: "Miércoles" },
    { id: 4, name: "Jueves" },
    { id: 5, name: "Viernes" },
    { id: 6, name: "Sábado" },
  ]);
  const [timeStart, setTimeStart] = createSignal<string>("");
  const [timeEnd, setTimeEnd] = createSignal<string>("");

  // Inicialización del selector de hora con Flatpickr
  const initializeFlatpickr = () => {
    flatpickr("#timeStart", {
      enableTime: true,
      noCalendar: true,
      dateFormat: "H:i",
      time_24hr: true,
      onChange: ([selectedDate]) => setTimeStart(selectedDate?.toTimeString().slice(0, 5) || ""),
    });

    flatpickr("#timeEnd", {
      enableTime: true,
      noCalendar: true,
      dateFormat: "H:i",
      time_24hr: true,
      onChange: ([selectedDate]) => setTimeEnd(selectedDate?.toTimeString().slice(0, 5) || ""),
    });
  };

  // Se ejecuta al montar el componente
  onMount(() => {
    initializeFlatpickr();
  });

  // Maneja el envío del formulario
  const handleSubmit = (e: Event) => {
    e.preventDefault();
    const target = e.target as HTMLFormElement;

    const formData: FormData = {
      officeId: parseInt(target.office.value),
      doctorId: target.doctor.value,
      serviceId: parseInt(target.service.value),
      day: parseInt(target.day.value),
      timeStart: timeStart(),
      timeEnd: timeEnd(),
    };

    console.log("Form Submitted:", formData);
    alert(`Formulario enviado:\n${JSON.stringify(formData, null, 2)}`);
  };

  return (
    <div class="register-form-content">
      <form onSubmit={handleSubmit} class="form">
        <div class="form-group">
          <label for="office">Office</label>
          <select id="office" name="office" required>
            <option value="" disabled selected>
              Select an office
            </option>
            {offices().map((office) => (
              <option value={office.id}>{office.name}</option>
            ))}
          </select>
        </div>

        <div class="form-group">
          <label for="doctor">Doctor</label>
          <select id="doctor" name="doctor" required>
            <option value="" disabled selected>
              Select a doctor
            </option>
            {doctors().map((doctor) => (
              <option value={doctor.id}>
                {doctor.firstName} {doctor.lastName1} {doctor.lastName2}
              </option>
            ))}
          </select>
        </div>

        <div class="form-group">
          <label for="service">Service</label>
          <select id="service" name="service" required>
            <option value="" disabled selected>
              Select a service
            </option>
            {services().map((service) => (
              <option value={service.id}>{service.name}</option>
            ))}
          </select>
        </div>

        <div class="form-group">
          <label for="day">Day</label>
          <select id="day" name="day" required>
            <option value="" disabled selected>
              Select a day
            </option>
            {days().map((day) => (
              <option value={day.id}>{day.name}</option>
            ))}
          </select>
        </div>

        <div class="form-group">
          <label for="timeStart">Start Time</label>
          <input
            type="text"
            id="timeStart"
            name="timeStart"
            placeholder="Select start time"
            required
          />
        </div>

        <div class="form-group">
          <label for="timeEnd">End Time</label>
          <input
            type="text"
            id="timeEnd"
            name="timeEnd"
            placeholder="Select end time"
            required
          />
        </div>

        <div class="form-actions">
          <button type="submit" class="btn-submit">Register</button>
          <button type="button" class="btn-cancel" onClick={props.onCancel}>
            Cancelar
          </button>
        </div>
      </form>
    </div>
  );
};

export default RegisterOfficeSchedule;
