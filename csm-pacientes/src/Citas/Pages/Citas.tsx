import { createResource, createSignal, createEffect, Suspense } from "solid-js";
import { useNavigate, useParams } from "@solidjs/router";
import Calendario from "../../components/Calendario";
import {
  getParamsForAppointment,
  getOfficeSchedules,
  registerAppointment,
  getAppointmentByID,
} from "../Services/CatalogServices";
import type {
  Services,
  Shift,
  SchedulesAppointmentRequest,
  OfficeScheduleResponse,
  Beneficiary,
  RegisterAppointmentRequest,
  AppointmentByID,
} from "../Models/Catalogs";
import { useAuth } from "../../services/AuthContext";

const Citas = () => {
  const paramsUrl = useParams();
  const appointmentId = paramsUrl.id;
  const navigate = useNavigate();
  const { token } = useAuth();

  const [shift, setShift] = createSignal<Shift | null>(null);
  const [service, setService] = createSignal<Services | null>(null);
  const [fullDate, setFullDate] = createSignal<string | null>(null);
  const [schedules, setSchedules] = createSignal<OfficeScheduleResponse[]>([]);
  const [selectedSchedule, setSelectedSchedule] = createSignal<OfficeScheduleResponse | null>(null);
  const [selectedPatient, setSelectedPatient] = createSignal<string | null>(null);
  const [notes, setNotes] = createSignal<string>("");
  const [symptoms, setSymptoms] = createSignal<string>("");
  const [scheduleError, setScheduleError] = createSignal<string | null>(null);
  const [loading, setLoading] = createSignal(false);

  const [paramsData] = createResource(
    () => token() ?? undefined,
    getParamsForAppointment
  );


  const [appointmentData] = createResource(
    () => (appointmentId && paramsData() ? appointmentId : undefined),
    (id) => getAppointmentByID(id, token() ?? undefined)
  );

  createEffect(() => {
    if (paramsData.loading || appointmentData.loading) return;

    if (appointmentData() && paramsData()?.data) {
      const app: AppointmentByID = appointmentData()!.data!;
      setSelectedPatient(app.beneficiaryID ? app.beneficiaryID : app.patientID);
      setService(
        paramsData()?.data?.services?.find((s) => s.id === app.serviceID) || null
      );
      setShift(
        paramsData()?.data?.shifts?.find((s) => s.id === app.shiftID) || null
      );
      setFullDate(new Date(app.timeStart).toISOString());
      setNotes(app.reason || "");
      setSymptoms(app.symptoms || "");
    }
  });

  const handlePatientChange = (e: Event) => {
    const target = e.currentTarget as HTMLSelectElement;
    setSelectedPatient(target.value);
  };

  const handleServiceChange = (e: Event) => {
    const target = e.currentTarget as HTMLSelectElement;
    const selectedServiceId = parseInt(target.value);
    if (paramsData()?.data?.services) {
      const selectedService = paramsData()?.data?.services.find(
        (svc) => svc.id === selectedServiceId
      );
      setService(selectedService || null);
    }
  };

  const handleShiftChange = (e: Event) => {
    const target = e.currentTarget as HTMLSelectElement;
    const selectedShiftId = parseInt(target.value);
    if (paramsData()?.data?.shifts) {
      const selectedShift = paramsData()?.data?.shifts.find(
        (sft) => sft.id === selectedShiftId
      );
      setShift(selectedShift || null);
    }
  };

  createEffect(async () => {
    if (!service() || !shift() || !fullDate() || !selectedPatient()) return;

    const appointmentRequest: SchedulesAppointmentRequest = {
      service: service()!.id,
      shift: shift()!.id,
      appointmentDate: fullDate()!,
    };

    try {
      setLoading(true);
      setScheduleError(null);

      const response = await getOfficeSchedules(
        appointmentRequest,
        token() ?? undefined
      );
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

  const confirmAppointment = async () => {
    if (!selectedSchedule() || !selectedPatient()) {
      alert("Debe seleccionar un horario y un paciente antes de continuar.");
      return;
    }
    if (!fullDate() || !paramsData()?.data) {
      alert("Debe seleccionar una fecha válida y tener los datos cargados.");
      return;
    }
    try {
      const date = new Date(fullDate()!);
      const timeStart = new Date(
        date.toISOString().split("T")[0] + `T${selectedSchedule()!.timeStart}Z`
      );
      const timeEnd = new Date(
        date.toISOString().split("T")[0] + `T${selectedSchedule()!.timeEnd}Z`
      );

      const patientID = paramsData()?.data?.patients?.accountHolderID;
      if (!patientID) {
        alert("No se encontró el ID del paciente principal.");
        return;
      }
      const selectedPatientID = selectedPatient();
      const appointmentDataToRegister: RegisterAppointmentRequest = {
        scheduleID: selectedSchedule()!.id,
        patientID,
        ...(selectedPatientID !== patientID && { beneficiaryID: selectedPatientID }),
        timeStart: timeStart.toISOString(),
        timeEnd: timeEnd.toISOString(),
        reason: notes(),
        symptoms: symptoms(),
      };

      console.log("Datos de la cita:", appointmentDataToRegister);
      const response = await registerAppointment(
        appointmentDataToRegister,
        token() ?? undefined
      );
      if (response.data) {
        alert("¡Cita registrada exitosamente!");
        navigate("/");
      } else {
        alert("Error al registrar la cita.");
      }
    } catch (error: any) {
      console.error("Error al registrar la cita:", error);
      alert("Ocurrió un error al registrar la cita. Revise la consola para más detalles.");
    }
  };

  const autoResizeTextarea = (event: Event) => {
    const target = event.target as HTMLTextAreaElement;
    target.style.height = "auto";
    target.style.height = `${target.scrollHeight}px`;
  };

  return (
    <Suspense fallback={<div>Cargando datos...</div>}>
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
                  {paramsData()?.data?.patients && (
                    <option value={paramsData()?.data?.patients.accountHolderID}>
                      {paramsData()?.data?.patients.fullName}
                    </option>
                  )}
                  {paramsData()?.data?.patients?.benefeciaries?.length ? (
                    paramsData()?.data?.patients.benefeciaries.map(
                      (beneficiary: Beneficiary) => (
                        <option value={beneficiary.beneficiaryID}>
                          {beneficiary.fullName}
                        </option>
                      )
                    )
                  ) : (
                    <option disabled>No hay beneficiarios disponibles</option>
                  )}
                </select>

                <h2>Selecciona Servicio</h2>
                <select
                  id="service"
                  required
                  value={service()?.id ? service()!.id.toString() : ""}
                  onChange={handleServiceChange}
                >
                  <option value="">-- Servicios --</option>
                  {paramsData()?.data?.services?.map((srv) => (
                    <option value={srv.id}>{srv.name}</option>
                  ))}
                </select>

                <h2>Selecciona Turno</h2>
                <select
                  id="turno"
                  required
                  value={shift()?.id ? shift()!.id.toString() : ""}
                  onChange={handleShiftChange}
                >
                  <option value="">-- Turno --</option>
                  {paramsData()?.data?.shifts?.map((sft) => (
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
                {loading() && <div class="loading-message">Cargando horarios...</div>}
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
                <button type="button" class="confirm-button" onClick={confirmAppointment}>
                  Confirmar Cita
                </button>
                <button type="button" class="cancel-button" onClick={() => navigate("/")}>
                  Cancelar
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Suspense>
  );
};

export default Citas;
