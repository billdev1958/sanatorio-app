import { Component } from "solid-js";
import { For } from "solid-js";

interface ConsultaMedica {
  fecha: string;
  servicio: string;
  consultorio: string;
  status: "pendiente" | "confirmada" | "cancelada";
}

const consultas: ConsultaMedica[] = [
  {
    fecha: "2024-01-20",
    servicio: "Consulta General",
    consultorio: "105",
    status: "confirmada",
  },
  {
    fecha: "2024-02-15",
    servicio: "Odontología",
    consultorio: "203",
    status: "pendiente",
  },
  {
    fecha: "2024-03-10",
    servicio: "Cardiología",
    consultorio: "301",
    status: "cancelada",
  },
  {
    fecha: "2024-03-15",
    servicio: "Dermatología",
    consultorio: "201",
    status: "confirmada",
  },
  {
    fecha: "2024-04-05",
    servicio: "Oftalmología",
    consultorio: "305",
    status: "pendiente",
  },
  {
    fecha: "2024-04-22",
    servicio: "Nutrición",
    consultorio: "102",
    status: "confirmada",
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
            <th>Fecha de Cita</th>
            <th>Servicio</th>
            <th>Consultorio</th>
            <th>Status</th>
            <th>Acciones</th>
          </tr>
        </thead>
        <tbody>
          <For each={consultas}>
            {(consulta) => (
              <tr>
                <td>{consulta.fecha}</td>
                <td>{consulta.servicio}</td>
                <td>{consulta.consultorio}</td>
                <td class={`status-${consulta.status}`}>{consulta.status}</td>
                <td>
                  <button class="btn-editar">Editar</button>
                  <button class="btn-confirmar">Confirmar</button>
                  <button class="btn-cancelar">Cancelar</button>
                </td>
              </tr>
            )}
          </For>
        </tbody>
      </table>
    </div>
  );
};

export default ConsultasHome;
