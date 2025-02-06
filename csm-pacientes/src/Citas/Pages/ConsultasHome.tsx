import { createSignal, createEffect } from "solid-js";
import { getPatientAppointments } from "../Services/CatalogServices";
import { Appointment } from "../Models/Catalogs";
import { useAuth } from "../../services/AuthContext";
import { useNavigate } from "@solidjs/router";

const ConsultasHome = () => {
  const [appointments, setAppointments] = createSignal<Appointment[]>([]);
  const [error, setError] = createSignal<string | null>(null);
  const [loading, setLoading] = createSignal<boolean>(false);

  const { token } = useAuth();
  const navigate = useNavigate();

  createEffect(async () => {
    setLoading(true);
    try {
      const data = await getPatientAppointments(token() ?? undefined);
      setAppointments(data);
      setError(null);
    } catch (err: any) {
      console.error("Error al obtener citas:", err);
      setError(err.message || "Error al obtener el historial de citas.");
      setAppointments([]);
    } finally {
      setLoading(false);
    }
  });

  const formatFechaConHoras = (timeStart: string, timeEnd: string): string => {
    const startDate = new Date(timeStart);
    const endDate = new Date(timeEnd);
    const fecha = startDate.toISOString().split("T")[0]; // "YYYY-MM-DD"
    const horaInicio = startDate.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
    const horaFin = endDate.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
    return `${fecha} ${horaInicio} - ${horaFin}`;
  };

  return (
    <div class="consultas-home">
      <h2>Mis Consultas Médicas</h2>
      <a href="/citas" class="btn-agendar">
        Agendar Nueva Consulta
      </a>

      {loading() && <p>Cargando citas...</p>}
      {error() && <p class="error-message">Error: {error()}</p>}

      {!loading() && !error() && (
        <table>
          <thead>
            <tr>
              <th>Fecha y Hora</th>
              <th>Paciente</th>
              <th>Tipo</th>
              <th>Servicio</th>
              <th>Consultorio</th>
              <th>Status</th>
              <th>Acciones</th>
            </tr>
          </thead>
          <tbody>
            {appointments().map((appointment) => {
              const fechaCompleta = formatFechaConHoras(appointment.TimeStart, appointment.TimeEnd);
              // Se determina el tipo de paciente: si BeneficiaryID es null es "Titular", de lo contrario "Beneficiario"
              const tipoPaciente = appointment.BeneficiaryID === null ? "Titular" : "Beneficiario";
              return (
                <tr>
                  <td>{fechaCompleta}</td>
                  <td>{appointment.PatientName}</td>
                  <td>
                    <span class={`tipo-${tipoPaciente.toLowerCase()}`}>
                      {tipoPaciente}
                    </span>
                  </td>
                  <td>{appointment.ServiceName}</td>
                  <td>{appointment.OfficeName}</td>
                  <td class={`status-${appointment.StatusName.toLowerCase()}`}>
                    {appointment.StatusName}
                  </td>
                  <td>
                    <button
                      class="btn-editar"
                      onClick={() => navigate("/citas/" + appointment.AppointmentID)}
                    >
                      Editar
                    </button>
                    <button class="btn-confirmar">Confirmar</button>
                    <button class="btn-cancelar">Cancelar</button>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default ConsultasHome;
