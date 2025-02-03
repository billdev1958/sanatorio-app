import { Component } from "solid-js";
import { For } from "solid-js";

/*interface Appointments {
  patient: string;
  beneficiary?: string;
  serviceName: string;
  officeName: string;
  timeStart: string;
  timeEnd: string; 
  statusID: number;
  statusName: string;
}*/

interface ConsultaMedica {
  timeStart: string;
  timeEnd: string;
  paciente: string;
  servicio: string;
  consultorio: string;
  status: "pendiente" | "confirmada" | "cancelada";
  tipoPaciente: "titular" | "beneficiario";
}

const formatFechaConHoras = (timeStart: string, timeEnd: string) => {
  const startDate = new Date(timeStart);
  const endDate = new Date(timeEnd);

  const fecha = startDate.toISOString().split("T")[0]; // YYYY-MM-DD
  const horaInicio = startDate.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
  const horaFin = endDate.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });

  return `${fecha} ${horaInicio} - ${horaFin}`;
};

const consultas: ConsultaMedica[] = [
  {
    timeStart: "2024-02-15T08:00:00Z",
    timeEnd: "2024-02-15T09:00:00Z",
    paciente: "Billy Rivera Salinas",
    servicio: "Consulta General",
    consultorio: "105",
    status: "confirmada",
    tipoPaciente: "titular",
  },
  {
    timeStart: "2024-02-16T10:30:00Z",
    timeEnd: "2024-02-16T11:30:00Z",
    paciente: "Ana Pérez",
    servicio: "Odontología",
    consultorio: "203",
    status: "pendiente",
    tipoPaciente: "beneficiario",
  },
  {
    timeStart: "2024-03-10T14:00:00Z",
    timeEnd: "2024-03-10T15:00:00Z",
    paciente: "Carlos Rivera",
    servicio: "Cardiología",
    consultorio: "301",
    status: "cancelada",
    tipoPaciente: "beneficiario",
  },
  {
    timeStart: "2024-03-15T09:00:00Z",
    timeEnd: "2024-03-15T10:00:00Z",
    paciente: "Billy Rivera Salinas",
    servicio: "Dermatología",
    consultorio: "201",
    status: "confirmada",
    tipoPaciente: "titular",
  },
  {
    timeStart: "2024-04-05T07:30:00Z",
    timeEnd: "2024-04-05T08:30:00Z",
    paciente: "Carlos Rivera",
    servicio: "Oftalmología",
    consultorio: "305",
    status: "pendiente",
    tipoPaciente: "beneficiario",
  },
  {
    timeStart: "2024-04-22T11:00:00Z",
    timeEnd: "2024-04-22T12:00:00Z",
    paciente: "Billy Rivera Salinas",
    servicio: "Nutrición",
    consultorio: "102",
    status: "confirmada",
    tipoPaciente: "titular",
  },
];

const ConsultasHome: Component = () => {
  return (
    <div class="consultas-home">
      <h2>Mis Consultas Médicas</h2>

      <a href="/citas" class="btn-agendar">
        Agendar Nueva Consulta
      </a>

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
          <For each={consultas}>
            {(consulta) => {
              const fechaCompleta = formatFechaConHoras(consulta.timeStart, consulta.timeEnd);
              
              return (
                <tr>
                  <td>{fechaCompleta}</td>
                  <td>{consulta.paciente}</td>
                  <td>
                    <span class={`tipo-${consulta.tipoPaciente}`}>
                      {consulta.tipoPaciente === "titular" ? "Titular" : "Beneficiario"}
                    </span>
                  </td>
                  <td>{consulta.servicio}</td>
                  <td>{consulta.consultorio}</td>
                  <td class={`status-${consulta.status}`}>{consulta.status}</td>
                  <td>
                    <button class="btn-editar">Editar</button>
                    <button class="btn-confirmar">Confirmar</button>
                    <button class="btn-cancelar">Cancelar</button>
                  </td>
                </tr>
              );
            }}
          </For>
        </tbody>
      </table>
    </div>
  );
};

export default ConsultasHome;
