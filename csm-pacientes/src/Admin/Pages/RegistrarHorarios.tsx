import { createSignal, onMount } from "solid-js";
import flatpickr from "flatpickr";
import "flatpickr/dist/flatpickr.min.css";
import api from "../Services/Api";
import {
  GetOfficeScheduleApiResponse,
  GetOfficeScheduleInfoResponse,
  DayOfWeek,
  CatShift,
  CatService,
  Doctor,
  Office,
} from "../Services/RegisterSchedule";

type FormData = {
  selectedDays: number[];
  timeStart: string;
  timeEnd: string;
  timeDuration: string; // Duración en formato hh:mm
  shiftID: number | "";
  serviceID: number | "";
  doctorID: string;
  officeID: number | ""; // Campo para seleccionar oficina
  timeSlots: string[]; // Horarios generados
};

// Definimos el tipo para el request del endpoint POST /v1/schedule
interface RegisterOfficeScheduleRequest {
  selectedDays: number[];
  timeStart: string;    // Formato HH:MM, ej: "09:00"
  timeEnd: string;      // Formato HH:MM, ej: "12:00"
  timeDuration: string; // Formato hh:mm, ej: "01:00"
  shiftID: string;      
  serviceID: string;    
  doctorID: string;     
  officeID: number;     
  timeSlots: string[];  
}

function RegisterOfficeScheduleForm() {
  const [daysOfWeek, setDaysOfWeek] = createSignal<DayOfWeek[]>([]);
  const [shifts, setShifts] = createSignal<CatShift[]>([]);
  const [services, setServices] = createSignal<CatService[]>([]);
  const [doctors, setDoctors] = createSignal<Doctor[] | null>(null);
  const [offices, setOffices] = createSignal<Office[]>([]);

  const [formData, setFormData] = createSignal<FormData>({
    selectedDays: [],
    timeStart: "",
    timeEnd: "",
    timeDuration: "00:00",
    shiftID: "",
    serviceID: "",
    doctorID: "",
    officeID: "",
    timeSlots: [],
  });

  let timeStartPicker: HTMLInputElement | undefined;
  let timeEndPicker: HTMLInputElement | undefined;
  let timeDurationPicker: HTMLInputElement | undefined;

  // Función para generar horarios (time slots)
  const generateTimeSlots = () => {
    const { timeStart, timeEnd, timeDuration } = formData();

    if (!timeStart || !timeEnd || !timeDuration) {
      alert("Por favor, completa los campos de tiempo correctamente.");
      return;
    }

    const start = new Date(timeStart);
    const end = new Date(timeEnd);
    const [durationHours, durationMinutes] = timeDuration.split(":").map(Number);

    if (isNaN(durationHours) || isNaN(durationMinutes)) {
      alert("Duración no válida. Usa el formato hh:mm.");
      return;
    }

    const slots: string[] = [];
    const durationMs = (durationHours * 60 + durationMinutes) * 60 * 1000;
    let currentTime = start;

    while (currentTime < end) {
      const nextTime = new Date(currentTime.getTime() + durationMs);
      if (nextTime > end) break;
      slots.push(`${formatTime(currentTime)} - ${formatTime(nextTime)}`);
      currentTime = nextTime;
    }

    setFormData({ ...formData(), timeSlots: slots });
  };

  // Formatea una hora en formato hh:mm
  const formatTime = (date: Date) => {
    return `${date.getHours().toString().padStart(2, "0")}:${date.getMinutes().toString().padStart(2, "0")}`;
  };

  onMount(async () => {
    try {
      const response = await api.get<GetOfficeScheduleApiResponse>("/admin/schedule");

      if (response.data.status === "success" && response.data.data) {
        const data: GetOfficeScheduleInfoResponse = response.data.data;

        setDaysOfWeek(data.day_of_week);
        setShifts(data.cat_shift);
        setServices(data.cat_services);
        setDoctors(data.doctor);
        setOffices(data.office);
      } else {
        console.error("Error en la respuesta de la API:", response.data.message);
      }
    } catch (error: any) {
      if (error.response?.status === 401) {
        console.error("No autorizado: falta un token válido.");
        alert("Sesión expirada o no autorizada. Por favor, inicia sesión nuevamente.");
      } else {
        console.error("Error al realizar la solicitud:", error);
        alert("Ocurrió un error al cargar los datos. Inténtalo más tarde.");
      }
    }

    flatpickr(timeStartPicker!, {
      enableTime: true,
      noCalendar: true,
      dateFormat: "H:i",
      defaultDate: "08:00",
      onChange: (selectedDates) => {
        setFormData({
          ...formData(),
          timeStart: selectedDates.length > 0 ? selectedDates[0].toISOString() : "",
        });
      },
    });

    flatpickr(timeEndPicker!, {
      enableTime: true,
      noCalendar: true,
      dateFormat: "H:i",
      defaultDate: "17:00",
      onChange: (selectedDates) => {
        setFormData({
          ...formData(),
          timeEnd: selectedDates.length > 0 ? selectedDates[0].toISOString() : "",
        });
      },
    });

    flatpickr(timeDurationPicker!, {
      enableTime: true,
      noCalendar: true,
      dateFormat: "H:i",
      time_24hr: true,
      defaultDate: "00:00",
      onChange: (selectedDates) => {
        setFormData({
          ...formData(),
          timeDuration: selectedDates.length > 0 ? selectedDates[0].toTimeString().slice(0, 5) : "",
        });
      },
    });
  });

  const toggleDaySelection = (day: number) => {
    const selectedDays = formData().selectedDays;
    if (selectedDays.includes(day)) {
      setFormData({
        ...formData(),
        selectedDays: selectedDays.filter((d) => d !== day),
      });
    } else {
      setFormData({
        ...formData(),
        selectedDays: [...selectedDays, day],
      });
    }
  };

  const handleSelectChange = (
    e: Event & { currentTarget: HTMLSelectElement | HTMLInputElement }
  ) => {
    const { name, value } = e.currentTarget;
    const newValue = name === "officeID" && value !== "" ? Number(value) : value;
    setFormData({ ...formData(), [name]: newValue });
  };

  const submitSchedule = async (data: RegisterOfficeScheduleRequest) => {
    try {
      const response = await api.post("/v1/schedule", data);
      console.log("Respuesta del servidor:", response.data);
      alert("Horario registrado con éxito");
    } catch (error: any) {
      console.error("Error al registrar el horario:", error);
      alert("Ocurrió un error al registrar el horario. Por favor, intenta nuevamente.");
    }
  };

  const handleSubmit = (e: Event): void => {
    e.preventDefault();
    generateTimeSlots();

    const currentFormData = formData();

    // Convertir timeStart, timeEnd a formato HH:MM
    const timeStartFormatted = new Date(currentFormData.timeStart).toTimeString().slice(0, 5);
    const timeEndFormatted = new Date(currentFormData.timeEnd).toTimeString().slice(0, 5);

    const payload: RegisterOfficeScheduleRequest = {
      selectedDays: currentFormData.selectedDays,
      timeStart: timeStartFormatted,
      timeEnd: timeEndFormatted,
      timeDuration: currentFormData.timeDuration,
      shiftID: String(currentFormData.shiftID),
      serviceID: String(currentFormData.serviceID),
      doctorID: currentFormData.doctorID,
      officeID: Number(currentFormData.officeID),
      timeSlots: currentFormData.timeSlots,
    };

    console.log("Formulario Enviado (Payload):");
    console.log(JSON.stringify(payload, null, 2));

    submitSchedule(payload);
  };

  return (
    <div class="form-container">
      <form onSubmit={handleSubmit}>
        <div class="form-group">
          <label>Select Days</label>
          <div class="day-selector">
            {daysOfWeek().map((day) => (
              <button
                type="button"
                class={`day-button ${
                  formData().selectedDays.includes(day.id) ? "selected" : ""
                }`}
                onClick={() => toggleDaySelection(day.id)}
              >
                {day.name}
              </button>
            ))}
          </div>
        </div>

        <div class="form-group">
          <label>Start Time</label>
          <input type="text" ref={timeStartPicker} />
        </div>

        <div class="form-group">
          <label>End Time</label>
          <input type="text" ref={timeEndPicker} />
        </div>

        <div class="form-group">
          <label>Duration (hh:mm)</label>
          <input type="text" ref={timeDurationPicker} />
        </div>

        <div class="form-group">
          <label for="shiftID">Shift</label>
          <select
            id="shiftID"
            name="shiftID"
            required
            value={formData().shiftID}
            onInput={(e) => handleSelectChange(e as any)}
          >
            <option value="">Select...</option>
            {shifts().map((shift) => (
              <option value={shift.id}>{shift.name}</option>
            ))}
          </select>
        </div>

        <div class="form-group">
          <label for="serviceID">Service</label>
          <select
            id="serviceID"
            name="serviceID"
            required
            value={formData().serviceID}
            onInput={(e) => handleSelectChange(e as any)}
          >
            <option value="">Select a service...</option>
            {services().map((service) => (
              <option value={service.id}>{service.name}</option>
            ))}
          </select>
        </div>

        <div class="form-group">
          <label for="doctorID">Doctor</label>
          <select
            id="doctorID"
            name="doctorID"
            required
            value={formData().doctorID}
            onInput={(e) => handleSelectChange(e as any)}
          >
            <option value="">Select a doctor...</option>
            {doctors() &&
              doctors()!.map((doc) => (
                <option value={doc.account_id}>
                  {doc.first_name} {doc.last_name_1} {doc.last_name_2}
                </option>
              ))}
          </select>
        </div>

        <div class="form-group">
          <label for="officeID">Office</label>
          <select
            id="officeID"
            name="officeID"
            required
            value={formData().officeID}
            onInput={(e) => handleSelectChange(e as any)}
          >
            <option value="">Select an office...</option>
            {offices().map((office) => (
              <option value={office.office_id}>{office.office_name}</option>
            ))}
          </select>
        </div>

        <button type="submit" class="submit-button">Submit</button>
      </form>

      <style>
        {`
          .form-container {
            max-width: 500px;
            margin: 0 auto;
            padding: 2em;
            background-color: #f9f9f9;
            border-radius: 8px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
          }

          .form-group {
            margin-bottom: 1.5em;
          }

          label {
            display: block;
            font-weight: bold;
            margin-bottom: 0.5em;
            font-size: 1.1em;
            color: #333;
          }

          input,
          select {
            width: 100%;
            padding: 0.7em;
            font-size: 1em;
            border: 1px solid #ccc;
            border-radius: 6px;
            box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.1);
          }

          .day-selector {
            display: flex;
            flex-wrap: wrap;
            gap: 0.5em;
          }

          .day-button {
            padding: 0.6em 1em;
            border: 1px solid #007bff;
            border-radius: 6px;
            background-color: white;
            color: #007bff;
            cursor: pointer;
            font-weight: bold;
          }

          .day-button.selected {
            background-color: #007bff;
            color: white;
          }

          .day-button:hover {
            background-color: #0056b3;
            color: white;
          }

          .submit-button {
            width: 100%;
            padding: 0.8em;
            font-size: 1.1em;
            font-weight: bold;
            color: white;
            background-color: #28a745;
            border: none;
            border-radius: 6px;
            cursor: pointer;
          }

          .submit-button:hover {
            background-color: #218838;
          }

          @media (max-width: 600px) {
            .form-container {
              padding: 1.5em;
            }

            label {
              font-size: 1em;
            }

            input,
            select {
              font-size: 0.9em;
            }

            .submit-button {
              font-size: 1em;
            }
          }
        `}
      </style>
    </div>
  );
}

export default RegisterOfficeScheduleForm;
