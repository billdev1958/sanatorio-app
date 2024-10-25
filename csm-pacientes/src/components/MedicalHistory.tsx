// src/components/HistorialMedicoForm.tsx
import { Component } from 'solid-js';
import PatientInfo from './PatientInfo';
import InfoPanel from './InfoPanel';

const HistorialMedicoForm: Component = () => {
  const medicalHistoryFields = [
    { label: 'Antecedentes Heredofamiliares', value: 'Diabetes en la familia.' },
    { label: 'Antecedentes Personales No Patológicos', value: 'No fuma, no consume alcohol.' },
    { label: 'Antecedentes Personales Patológicos', value: 'Hipertensión.' },
    { label: 'Antecedentes Gineco-Obstétricos / Androgénicos', value: 'No aplica.' },
    { label: 'Padecimiento Actual', value: 'Dolor en la parte inferior del abdomen.' },
  ];

  const systemInterrogationFields = [
    { label: 'Cardiovascular', value: 'Sin problemas.' },
    { label: 'Respiratorio', value: 'Sin dificultades.' },
    { label: 'Gastrointestinal', value: 'Reflujo ocasional.' },
    { label: 'Genitourinario', value: 'Sin anomalías.' },
    { label: 'Hemático y Linfático', value: 'No hay antecedentes.' },
    { label: 'Endocrino', value: 'Función tiroidea normal.' },
    { label: 'Nervioso', value: 'Sin síntomas.' },
    { label: 'Musculoesquelético', value: 'Dolor leve en las articulaciones.' },
    { label: 'Piel, Mucosas y Anexos', value: 'Piel seca en los codos.' },
  ];

  const vitalSignsFields = [
    { label: 'T/A', value: '120/80 mmHg' },
    { label: 'Temperatura', value: '36.7 °C' },
    { label: 'Frecuencia Cardíaca', value: '72 bpm' },
    { label: 'Frecuencia Respiratoria', value: '16 rpm' },
    { label: 'Peso (kg)', value: '70 kg' },
    { label: 'Talla (m)', value: '1.75 m' },
    { label: 'IMC', value: '22.9' },
  ];

  const physicalExamFields = [
    { label: 'Habitus Exterior', value: 'Paciente alerta, bien nutrido.' },
    { label: 'Cabeza', value: 'Normocefálica, sin lesiones.' },
    { label: 'Cuello', value: 'Sin adenopatías.' },
    { label: 'Tórax', value: 'Simétrico, sin masas.' },
    { label: 'Abdomen', value: 'Blando, depresible.' },
    { label: 'Genitales', value: 'No aplica.' },
    { label: 'Extremidades', value: 'Sin edemas.' },
    { label: 'Piel', value: 'Piel seca.' },
  ];

  const diagnosticFields = [
    { label: 'Diagnósticos o Problemas Clínicos', value: 'Gastritis.' },
  ];

  const treatmentFields = [
    { label: 'Tratamiento Farmacológico', value: 'Omeprazol 20mg cada 24 horas.' },
  ];

  const prognosisFields = [
    { label: 'Pronóstico', value: 'Bueno.' },
    { label: 'Médico', value: 'Dr. Juan Pérez' },
    { label: 'Cédula Profesional', value: '1234567890' },
    { label: 'Cédula Especialidad', value: '0987654321' },
  ];

  return (
    <div class="historial-medico-form">
      <div class="sidebar">
        <div class="title-box">
          <h1>Historia Clínica General</h1>
        </div>
        <PatientInfo />
      </div>
      <div class="content">
        <InfoPanel title="Historial Médico" fields={medicalHistoryFields} />
        <InfoPanel title="Interrogatorio por Aparatos y Sistemas" fields={systemInterrogationFields} />
        <InfoPanel title="Signos Vitales" fields={vitalSignsFields} />
        <InfoPanel title="Exploración Física" fields={physicalExamFields} />
        <InfoPanel title="Diagnósticos" fields={diagnosticFields} />
        <InfoPanel title="Tratamiento" fields={treatmentFields} />
        <InfoPanel title="Pronóstico" fields={prognosisFields} />
      </div>
    </div>
  );
};

export default HistorialMedicoForm;
