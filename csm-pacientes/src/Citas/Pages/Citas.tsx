import { createSignal, createEffect } from "solid-js";
import { useNavigate, useParams } from "@solidjs/router";
import Calendario from "../../components/Calendario";
import {
  getParamsForAppointment,
  getOfficeSchedules,
  registerAppointment,
  getAppointmentByID,
} from "../Services/CatalogServices";
import {
  Services,
  Shift,
  SchedulesAppointmentRequest,
  OfficeScheduleResponse,
  PatientAndBeneficiaries,
  Beneficiary,
  RegisterAppointmentRequest,
  AppointmentByID,
} from "../Models/Catalogs";
import { useAuth } from "../../services/AuthContext";

const Citas = () => {
  const paramsUrl = useParams(); // Para obtener parámetros de la URL
  const appointmentId = paramsUrl.id; // Si existe, se está editando una cita
  const navigate = useNavigate();
  const { token } = useAuth();

  // Estados para formulario y datos generales
  const [shift, setShift] = createSignal<Shift | null>(null);
  const [service, setService] = createSignal<Services | null>(null);
  const [fullDate, setFullDate] = createSignal<string | null>(null);
  const [schedules, setSchedules] = createSignal<OfficeScheduleResponse[]>([]);
  const [selectedSchedule, setSelectedSchedule] = createSignal<OfficeScheduleResponse | null>(null);
  const [selectedPatient, setSelectedPatient] = createSignal<string | null>(null);
  const [notes, setNotes] = createSignal<string | null>(null);
  const [symptoms, setSymptoms] = createSignal<string | null>(null);
  const [paramsData, setParamsData] = createSignal<{
    patients: PatientAndBeneficiaries;
    services: Services[];
    shifts: Shift[];
  } | null>(null);
  const [scheduleError, setScheduleError] = createSignal<string | null>(null);
  const [loading, setLoading] = createSignal(false);

  // Obtener parámetros generales (pacientes, servicios, turnos)
  createEffect(async () => {
    try {
      const response = await getParamsForAppointment(token() ?? undefined);
      if (response.data) {
        const { patients, services, shifts } = response.data;
        const normalizedPatients: PatientAndBeneficiaries = {
          ...patients,
          benefeciaries: patients?.benefeciaries ?? [],
        };
        setParamsData({
          patients: normalizedPatients,
          services,
          shifts,
        });
      } else {
        throw new Error("No se encontraron servicios, turnos o pacientes disponibles.");
      }
    } catch (error: any) {
      console.error("Error al obtener datos:", error);
      setScheduleError(error.message || "Error al obtener los datos.");
    }
  });

  // Manejadores para cambios en selects
  const handlePatientChange = (e: Event) => {
    const target = e.currentTarget as HTMLSelectElement;
    setSelectedPatient(target.value);
  };

  const handleServiceChange = (e: Event) => {
    const target = e.currentTarget as HTMLSelectElement;
    const selectedServiceId = parseInt(target.value);
    const currentParams = paramsData();
    if (currentParams) {
      const selectedService = currentParams.services.find((svc) => svc.id === selectedServiceId);
      setService(selectedService || null);
    }
  };

  const handleShiftChange = (e: Event) => {
    const target = e.currentTarget as HTMLSelectElement;
    const selectedShiftId = parseInt(target.value);
    const currentParams = paramsData();
    if (currentParams) {
      const selectedShift = currentParams.shifts.find((sft) => sft.id === selectedShiftId);
      setShift(selectedShift || null);
    }
  };

  // Obtener horarios disponibles (se ejecuta cuando se seleccionan servicio, turno, fecha y paciente)
  createEffect(async () => {
    if (!service() || !shift() || !fullDate() || !selectedPatient()) {
      return;
    }

    const appointmentData: SchedulesAppointmentRequest = {
      service: service()!.id,
      shift: shift()!.id,
      appointmentDate: fullDate()!,
    };

    try {
      setLoading(true);
      setScheduleError(null);

      const response = await getOfficeSchedules(appointmentData, token() ?? undefined);

      if (response.data && Array.isArray(response.data)) {
        setSchedules(response.data);
      } else {
        setSchedules([]);
        setScheduleError("No se encontraron horarios disponibles.");
      }
    } catch (error: any) {
      console.error("Error al obtener horarios:", error);
      setSchedules([]);
      setScheduleError(error.message || "Error al obtener los horarios.");
    } finally {
      setLoading(false);
    }
  });

  // Si existe un appointmentId, obtener la cita (modo edición) y rellenar el formulario
  createEffect(async () => {
    if (!appointmentId) return; // Solo si hay un ID

    try {
      const response = await getAppointmentByID(appointmentId, token() ?? undefined);
      if (response.data) {
        const appointment: AppointmentByID = response.data;
        // Si beneficiaryID viene con valor, se selecciona; de lo contrario, se selecciona el patientID.
        if (appointment.beneficiaryID) {
          setSelectedPatient(appointment.beneficiaryID);
        } else {
          setSelectedPatient(appointment.patientID);
        }
        // Si la estructura AppointmentByID incluye serviceID y shiftID, se usan para seleccionar servicio y turno
        if (paramsData()) {
          setService(paramsData()!.services.find((s) => s.id === appointment.serviceID) || null);
          setShift(paramsData()!.shifts.find((s) => s.id === appointment.shiftID) || null);
        }
        // Almacena el string ISO completo (incluyendo la parte horaria y milisegundos)
        setFullDate(new Date(appointment.timeStart).toISOString());
        setNotes(appointment.reason || "");
        setSymptoms(appointment.symptoms || "");
      }
    } catch (error: any) {
      console.error("Error al obtener la cita:", error);
      setScheduleError(error.message || "Error al cargar la cita.");
    }
  });

  // Función para confirmar la cita (registro)
  const confirmAppointment = async () => {
    if (!selectedSchedule() || !selectedPatient()) {
      alert("Debe seleccionar un horario y un paciente antes de continuar.");
      return;
    }
  
    if (!fullDate()) {
      alert("Debe seleccionar una fecha válida.");
      console.error("Error: fullDate es nulo o indefinido.");
      return;
    }
  
    try {
      const date = new Date(fullDate()!);
      const timeStart = new Date(date.toISOString().split("T")[0] + `T${selectedSchedule()!.timeStart}Z`);
      const timeEnd = new Date(date.toISOString().split("T")[0] + `T${selectedSchedule()!.timeEnd}Z`);
  
      if (isNaN(timeStart.getTime()) || isNaN(timeEnd.getTime())) {
        throw new Error("Los valores de timeStart o timeEnd no son válidos.");
      }
  
      const patientID = paramsData()!.patients.accountHolderID;
      const selectedPatientID = selectedPatient(); // ID del paciente seleccionado
  
      // Si el paciente seleccionado es el titular, beneficiaryID no se envía
      const appointmentData: RegisterAppointmentRequest = {
        scheduleID: selectedSchedule()!.id,
        patientID,
        ...(selectedPatientID !== patientID && { beneficiaryID: selectedPatientID }),
        timeStart: timeStart.toISOString(),
        timeEnd: timeEnd.toISOString(),
        reason: notes(),
        symptoms: symptoms(),
      };
  
      console.log("Datos de la cita:", appointmentData);
  
      const response = await registerAppointment(appointmentData, token() ?? undefined);
  
      if (response.data) {
        alert("Cita registrada exitosamente!");
        navigate("/");
      } else {
        alert("Error al registrar la cita.");
      }
    } catch (error: any) {
      console.error("Error al registrar la cita:", error);
      alert("Ocurrió un error al registrar la cita. Revise la consola para más detalles.");
    }
  };

  // Función para ajustar el alto del textarea según su contenido
  const autoResizeTextarea = (event: Event) => {
    const target = event.target as HTMLTextAreaElement;
    target.style.height = "auto";
    target.style.height = `${target.scrollHeight}px`;
  };

  return (
    <div class="citas-container">
      <div class="cita-card">
        <h1>{appointmentId ? "Editar Cita" : "Selecciona tu Cita"}</h1>

        <div class="cita-sections">
          <div class="left-section">
            <div class="form-section">
              <h2>Selecciona Paciente</h2>
              <select
                id="patient"
                required
                value={selectedPatient() || ""}
                onChange={handlePatientChange}
              >
                <option value="">-- Pacientes --</option>
                {paramsData()?.patients && (
                  <option value={paramsData()?.patients.accountHolderID}>
                    {paramsData()?.patients.fullName}
                  </option>
                )}
                {paramsData()?.patients?.benefeciaries?.length ? (
                  paramsData()?.patients?.benefeciaries.map((beneficiary: Beneficiary) => (
                    <option value={beneficiary.beneficiaryID}>
                      {beneficiary.fullName}
                    </option>
                  ))
                ) : (
                  <option disabled>No hay beneficiarios disponibles</option>
                )}
              </select>

              <h2>Selecciona Servicio</h2>
              <select
                id="service"
                required
                value={service()?.id ? service()?.id.toString() : ""}
                onChange={handleServiceChange}
              >
                <option value="">Servicios --</option>
                {paramsData()?.services.map((srv) => (
                  <option value={srv.id}>{srv.name}</option>
                ))}
              </select>

              <h2>Selecciona Turno</h2>
              <select
                id="turno"
                required
                value={shift()?.id ? shift()?.id.toString() : ""}
                onChange={handleShiftChange}
              >
                <option value="">Turno --</option>
                {paramsData()?.shifts.map((sft) => (
                  <option value={sft.id}>{sft.name}</option>
                ))}
              </select>

              <h2>Notas</h2>
              <textarea
                id="notes"
                class="text-input"
                placeholder="Agrega una nota opcional"
                value={notes() || ""}
                onInput={(e) => {
                  setNotes((e.target as HTMLTextAreaElement).value);
                  autoResizeTextarea(e);
                }}
              ></textarea>

              <h2>Síntomas</h2>
              <textarea
                id="symptoms"
                class="text-input"
                placeholder="Describe los síntomas opcionales"
                value={symptoms() || ""}
                onInput={(e) => {
                  setSymptoms((e.target as HTMLTextAreaElement).value);
                  autoResizeTextarea(e);
                }}
              ></textarea>
            </div>

            <div class="calendario-section">
              <h2>Selecciona una Fecha</h2>
              <Calendario
  selectedDate={fullDate() || undefined}
  onDateChange={(utcDate: string) => {
    console.log("Fecha en UTC:", utcDate);
    setFullDate(utcDate);
  }}
/>
              {fullDate() && (
                <p class="selected-date">Fecha seleccionada: {fullDate()}</p>
              )}
            </div>
          </div>

          <div class="right-section">
            <div class="schedules-section">
              <h2>Horarios Disponibles</h2>
              {loading() && (
                <div class="loading-message">Cargando horarios...</div>
              )}
              {!loading() && scheduleError() && (
                <div class="error-message">
                  <p>{scheduleError()}</p>
                </div>
              )}
              {!loading() && schedules().length > 0 && (
                <div class="schedules-grid">
                  {schedules().map((schedule) => {
                    const isDisabled = schedule.statusID !== 1;
                    return (
                      <div
                        class={`schedule-card ${
                          selectedSchedule() &&
                          selectedSchedule()!.id === schedule.id
                            ? "selected"
                            : ""
                        } ${isDisabled ? "disabled" : ""}`}
                        onClick={() => {
                          if (!isDisabled) {
                            setSelectedSchedule(schedule);
                          }
                        }}
                      >
                        <p class="time-label">Inicio:</p>
                        <p class="time-value">{schedule.timeStart}</p>
                        <p class="time-label">Fin:</p>
                        <p class="time-value">{schedule.timeEnd}</p>
                      </div>
                    );
                  })}
                </div>
              )}
            </div>

            <div class="action-buttons">
              <button
                type="button"
                class="confirm-button"
                onClick={confirmAppointment}
              >
                Confirmar Cita
              </button>
              <button
                type="button"
                class="cancel-button"
                onClick={() => navigate("/")}
              >
                Cancelar
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Citas;
