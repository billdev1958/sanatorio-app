import { Component } from 'solid-js';
import { For } from 'solid-js';

interface ConsultaMedica {
  fecha: string;
  servicio: string;
  consultorio: string;
  status: 'pendiente' | 'confirmada' | 'cancelada';
}

interface ListadoConsultasProps {
  consultas: ConsultaMedica[];
}

const ListadoConsultas: Component<ListadoConsultasProps> = (props) => {
  return (
    <div class="consultas-home"> {/* Usamos la clase consultas-home */}
      <h2>Mis Consultas Médicas</h2>

      <button class="btn-agendar">
        Agendar Nueva Consulta
      </button>

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
          <For each={props.consultas}>
            {(consulta) => (
              <tr>
                <td>{consulta.fecha}</td>
                <td>{consulta.servicio}</td>
                <td>{consulta.consultorio}</td>
                <td class={`status-${consulta.status}`}>
                  {consulta.status}
                </td>
                <td>
                  <div class="acciones-container"> 
                    <button class="btn-editar">Editar</button>
                    <button class="btn-confirmar">Confirmar</button>
                    <button class="btn-cancelar">Cancelar</button>
                  </div>
                </td>
              </tr>
            )}
          </For>
        </tbody>
      </table>
    </div>
  );
};

export default ListadoConsultas;