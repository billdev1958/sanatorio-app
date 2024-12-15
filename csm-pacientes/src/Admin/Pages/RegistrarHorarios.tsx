import { createSignal, onMount } from "solid-js";
import flatpickr from "flatpickr";
import "flatpickr/dist/flatpickr.min.css";
import api from "../Services/Api"; // Importa el cliente API configurado
import {
  DayOfWeek,
  CatShift,
  CatService,
  Office,
  GetOfficeScheduleInfoResponse,
} from "../Services/RegisterSchedule"; // Importa los modelos

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

function RegisterOfficeScheduleForm() {
  // Estados para los datos dinámicos obtenidos de la API
  const [daysOfWeek, setDaysOfWeek] = createSignal<DayOfWeek[]>([]);
  const [shifts, setShifts] = createSignal<CatShift[]>([]);
  const [services, setServices] = createSignal<CatService[]>([]);
  const [offices, setOffices] = createSignal<Office[]>([]); // Nuevo estado para oficinas

  const [formData, setFormData] = createSignal<FormData>({
    selectedDays: [],
    timeStart: "",
    timeEnd: "",
    timeDuration: "00:00",
    shiftID: "",
    serviceID: "",
    doctorID: "",
    officeID: "", // Campo para oficina
    timeSlots: [],
  });

  let timeStartPicker: HTMLInputElement | undefined;
  let timeEndPicker: HTMLInputElement | undefined;
  let timeDurationPicker: HTMLInputElement | undefined;

  // Consumir API al montar el componente
  onMount(async () => {
    try {
      // Llama al servicio
      const response = await api.get<GetOfficeScheduleInfoResponse>("/admin/schedule");
      const { day_of_week, cat_shift, cat_services, office } = response.data.data;

      // Actualiza los estados con los datos obtenidos
      setDaysOfWeek(day_of_week);
      setShifts(cat_shift);
      setServices(cat_services);
      setOffices(office); // Oficinas

    } catch (error) {
      console.error("Error al obtener datos del servicio:", error);
    }

    // Flatpickr para timeStart
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

    // Flatpickr para timeEnd
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

    // Flatpickr para timeDuration
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
    setFormData({ ...formData(), [name]: value });
  };

  const handleSubmit = (e: Event): void => {
    e.preventDefault();
    console.log("Form Data:", JSON.stringify(formData(), null, 2));
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

        <button type="submit">Submit</button>
      </form>
    </div>
  );
}

export default RegisterOfficeScheduleForm;
