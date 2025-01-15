import logoCMS from '../assets/logo_cms.png'; 

import { Component } from 'solid-js';
import Widget from '../components/Widget';

const Dashboard: Component = () => {
  return (
    <main class="content-area">
      <section class="welcome-section">
        <div class="welcome-content">
        <img src={logoCMS} alt="Logo de la Clínica" class="clinic-logo" />
        <h1>Bienvenido</h1>
          <p>Gestiona tus citas, beneficiarios, y revisa tus resultados de laboratorio de manera fácil y rápida.</p>
        </div>
      </section>

      {/* Sección de Widgets */}
      <section class="dashboard-widgets">
        <Widget 
          icon="fas fa-calendar-alt" 
          title="Agendar Cita" 
          description="Programa tu próxima consulta de manera rápida y sencilla." 
          link="/citas" 
          buttonText="Agendar Ahora" 
        />
        <Widget 
          icon="fas fa-user-plus" 
          title="Registrar Beneficiarios" 
          description="Gestiona a tus beneficiarios para un acceso más fácil a los servicios." 
          link="/beneficiary" 
          buttonText="Registrar" 
        />
        <Widget 
          icon="fas fa-flask" 
          title="Revisar citas" 
          description="Revisa todas tus citas y su status" 
          link="/consultas" 
          buttonText="Acceder" 
        />

      </section>
    </main>
  );
};

export default Dashboard;



