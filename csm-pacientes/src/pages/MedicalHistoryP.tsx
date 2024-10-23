// src/pages/HistorialMedicoPage.tsx
import { Component } from 'solid-js';
import HistorialMedico from '../components/MedicalHistory';

const HistorialMedicoPage: Component = () => {
  return (
    <div class="historial-medico-page">
      <h1>Historial Médico del Paciente</h1>
      <HistorialMedico />
    </div>
  );
};

export default HistorialMedicoPage;
