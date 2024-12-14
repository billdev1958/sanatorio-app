import logoCMS from '../../assets/logo_cms.png'; 

import { Component } from 'solid-js';
import Widget from '../../components/Widget.tsx';

const DashboardAdmin: Component = () => {
  return (
    <main class="content-area">
      <section class="welcome-section">
        <div class="welcome-content">
        <img src={logoCMS} alt="Logo de la Clínica" class="clinic-logo" />
        <h1>BIENVENIDO AL GESTOR DE HORARIOS DE LA CLINICA MULTIDISCIPLINARIA DE SALUD</h1>
          <p>Gestiona tus citas, beneficiarios, y revisa tus resultados de laboratorio de manera fácil y rápida.</p>
        </div>
      </section>

      {/* Sección de Widgets */}
      <section class="dashboard-widgets">
        <Widget 
          icon="fas fa-calendar-alt" 
          title="Registra doctor" 
          description="Registra doctores." 
          link="/admin/doctor" 
          buttonText="Agendar Ahora" 
        />
        <Widget 
          icon="fas fa-user-plus" 
          title="Registras oficinas" 
          description="Gestiona los consultorios." 
          link="/admin/offices" 
          buttonText="Registrar" 
        />
        <Widget 
          icon="fas fa-flask" 
          title="Registra horarios" 
          description="Gestiona los horarios de las consultas medicas." 
          link="/admin/schedule" 
          buttonText="Acceder" 
        />

      </section>
    </main>
  );
};

export default DashboardAdmin;



