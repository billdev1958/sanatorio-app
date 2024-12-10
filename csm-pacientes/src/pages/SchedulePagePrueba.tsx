import { createSignal } from "solid-js";
import RegisterOfficeSchedule from "../components/SchedulePrueba";
import ScheduleFilter from "../components/FilterTab";

interface Schedule {
  id: number;
  officeName: string;
  serviceName: string;
  doctorName: string;
  day: string;
  timeStart: string;
  timeEnd: string;
}

const SchedulesPage = () => {
  const [schedules, setSchedules] = createSignal<Schedule[]>([
    {
      id: 1,
      officeName: "Consultorio 1",
      serviceName: "Audiometría",
      doctorName: "Juan Pérez",
      day: "Lunes",
      timeStart: "08:00",
      timeEnd: "12:00",
    },
    {
      id: 2,
      officeName: "Consultorio 2",
      serviceName: "Acupuntura",
      doctorName: "Ana Gómez",
      day: "Martes",
      timeStart: "10:00",
      timeEnd: "14:00",
    },
  ]);

  const [showRegisterForm, setShowRegisterForm] = createSignal(false);

  const handleAddSchedule = () => {
    setShowRegisterForm(true);
  };

  const handleCloseForm = () => {
    setShowRegisterForm(false);
  };

  return (
    <div class="schedule-page">
      <ScheduleFilter />

      <div class="schedule-cards-wrapper">
        {schedules().map((schedule) => (
          <div class="schedule-card">
            <div class="schedule-card-content">
              <h2 class="schedule-office-name">{schedule.officeName}</h2>
              <p class="schedule-info">
                <strong>Servicio:</strong> {schedule.serviceName}
              </p>
              <p class="schedule-info">
                <strong>Doctor:</strong> {schedule.doctorName}
              </p>
              <p class="schedule-info">
                <strong>Día:</strong> {schedule.day}
              </p>
              <p class="schedule-info">
                <strong>Hora de Inicio:</strong> {schedule.timeStart}
              </p>
              <p class="schedule-info">
                <strong>Hora de Fin:</strong> {schedule.timeEnd}
              </p>
            </div>
          </div>
        ))}
      </div>

      <button class="btn-add-schedule" onClick={handleAddSchedule}>
        Registrar Nuevo Horario
      </button>

      {showRegisterForm() && (
        <div class="register-form-container">
          <RegisterOfficeSchedule onCancel={handleCloseForm} />
        </div>
      )}
    </div>
  );
};

export default SchedulesPage;
