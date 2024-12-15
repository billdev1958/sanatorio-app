import { createSignal, onMount } from "solid-js";
import flatpickr from "flatpickr";
import "flatpickr/dist/flatpickr.min.css";

type FormData = {
  selectedDays: number[];
  timeStart: string;
  timeEnd: string;
  timeDuration: string; // Duración en formato hh:mm
  shiftID: number | "";
  serviceID: number | "";
  doctorID: string;
  timeSlots: string[]; // Horarios generados
};

const DAYS_OF_WEEK = [
  { value: 0, label: "Sunday" },
  { value: 1, label: "Monday" },
  { value: 2, label: "Tuesday" },
  { value: 3, label: "Wednesday" },
  { value: 4, label: "Thursday" },
  { value: 5, label: "Friday" },
  { value: 6, label: "Saturday" },
];

function RegisterOfficeScheduleForm() {
  const [formData, setFormData] = createSignal<FormData>({
    selectedDays: [],
    timeStart: "",
    timeEnd: "",
    timeDuration: "00:00",
    shiftID: "",
    serviceID: "",
    doctorID: "",
    timeSlots: [],
  });

  let timeStartPicker: HTMLInputElement | undefined;
  let timeEndPicker: HTMLInputElement | undefined;
  let timeDurationPicker: HTMLInputElement | undefined;

  onMount(() => {
    // Flatpickr para timeStart
    flatpickr(timeStartPicker!, {
      enableTime: true,
      noCalendar: true,
      dateFormat: "H:i",
      defaultDate: "08:00", // Valor predeterminado
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
      defaultDate: "17:00", // Valor predeterminado
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
      defaultDate: "00:00", // Empieza desde 00:00
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
            {DAYS_OF_WEEK.map((day) => (
              <button
                type="button"
                class={`day-button ${
                  formData().selectedDays.includes(day.value) ? "selected" : ""
                }`}
                onClick={() => toggleDaySelection(day.value)}
              >
                {day.label}
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
          <small>Specify the duration in hours and minutes.</small>
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
            <option value="1">Morning</option>
            <option value="2">Evening</option>
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
            <option value="101">General Consultation</option>
            <option value="102">Pediatrics</option>
            <option value="103">Dermatology</option>
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
            <option value="uuid-1">Dr. John Smith</option>
            <option value="uuid-2">Dr. Alice Johnson</option>
            <option value="uuid-3">Dr. Emily Davis</option>
          </select>
        </div>

        <button type="submit">Submit</button>
      </form>

      <style>
        {`
          .form-container {
            max-width: 400px;
            margin: 0 auto;
            padding: 1em;
            border: 1px solid #ccc;
            border-radius: 5px;
            background: #f9f9f9;
            box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
          }
          .form-group {
            margin-bottom: 1em;
          }
          label {
            display: block;
            margin-bottom: 0.5em;
            font-weight: bold;
          }
          input, select {
            width: 100%;
            padding: 0.5em;
            border: 1px solid #ccc;
            border-radius: 4px;
          }
          .day-selector {
            display: flex;
            flex-wrap: wrap;
            gap: 0.5em;
          }
          .day-button {
            padding: 0.5em 1em;
            border: 1px solid #007bff;
            border-radius: 4px;
            background: white;
            color: #007bff;
            cursor: pointer;
          }
          .day-button.selected {
            background: #007bff;
            color: white;
          }
          .day-button:hover {
            background: #0056b3;
            color: white;
          }
          button {
            padding: 0.7em 1.5em;
            background: #007bff;
            color: white;
            border: none;
            border-radius: 4px;
            cursor: pointer;
          }
          button:hover {
            background: #0056b3;
          }
          small {
            display: block;
            margin-top: 0.3em;
            color: #6c757d;
            font-size: 0.85em;
          }
        `}
      </style>
    </div>
  );
}

export default RegisterOfficeScheduleForm;
