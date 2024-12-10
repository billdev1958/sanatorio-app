import { createSignal } from "solid-js";

const ScheduleFilter = () => {
  const [officeName, setOfficeName] = createSignal("");
  const [serviceName, setServiceName] = createSignal("");
  const [doctorName, setDoctorName] = createSignal("");
  const [dayOfWeek, setDayOfWeek] = createSignal("");

  // Options for service name and days of the week
  const serviceOptions = [
    "Audiometría",
    "Acupuntura",
    "Densitometría",
    "Electrocardiografía",
    "Ginecología",
    "Laboratorio clínico",
  ];

  const dayOptions = [
    "Lunes",
    "Martes",
    "Miércoles",
    "Jueves",
    "Viernes",
    "Sábado",
  ];

  return (
    <div class="filter-tab">
      <div class="filter-group">
        <label for="officeName">Nombre Oficina</label>
        <input
          type="text"
          id="officeName"
          value={officeName()}
          onInput={(e) => setOfficeName(e.currentTarget.value)}
          placeholder="Filtrar por nombre de oficina"
        />
      </div>

      <div class="filter-group">
        <label for="serviceName">Nombre Servicio</label>
        <select
          id="serviceName"
          value={serviceName()}
          onChange={(e) => setServiceName(e.currentTarget.value)}
        >
          <option value="">Seleccionar servicio</option>
          {serviceOptions.map((service) => (
            <option value={service}>{service}</option>
          ))}
        </select>
      </div>

      <div class="filter-group">
        <label for="doctorName">Nombre Doctor</label>
        <input
          type="text"
          id="doctorName"
          value={doctorName()}
          onInput={(e) => setDoctorName(e.currentTarget.value)}
          placeholder="Filtrar por nombre de doctor"
        />
      </div>

      <div class="filter-group">
        <label for="dayOfWeek">Día de la Semana</label>
        <select
          id="dayOfWeek"
          value={dayOfWeek()}
          onChange={(e) => setDayOfWeek(e.currentTarget.value)}
        >
          <option value="">Seleccionar día</option>
          {dayOptions.map((day) => (
            <option value={day}>{day}</option>
          ))}
        </select>
      </div>
    </div>
  );
};

export default ScheduleFilter;
