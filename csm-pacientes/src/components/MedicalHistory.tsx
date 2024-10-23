// src/components/HistorialMedicoForm.tsx
import { Component } from 'solid-js';

const HistorialMedicoForm: Component = () => {
  return (
    <div class="historial-medico-form">
      <div class="sidebar">
        <div class="title-box">
          <h1>Patient Medical History Form</h1>
        </div>
        <div class="patient-info">
          <h2><i class="fas fa-user"></i> Patient Information</h2>
          <p><strong>Name:</strong> Reinaldos Prydden</p>
          <p><strong>Email:</strong> csword2@comcast.net</p>
          <p><strong>Birth Date:</strong> 1/10/2024</p>
          <p><strong>Height (cm):</strong> 4</p>
          <p><strong>Weight (kg):</strong> 4</p>
          <p><strong>Gender:</strong> Option 2</p>
        </div>
      </div>
      <div class="content">
        <div class="info-panel">
          <h2><i class="fas fa-notes-medical"></i> Patient Medical History</h2>
          <p><strong>Drug Allergies:</strong> Proin eu mi.</p>
          <p><strong>Have you ever had:</strong> Option 2</p>
          <p><strong>Other illnesses:</strong> Cras non velit nec nisi vulputate nonummy.</p>
          <p><strong>Operations:</strong> Proin eu mi.</p>
          <p><strong>Current Medications:</strong> Proin eu mi.</p>
        </div>

        <div class="info-panel">
          <h2><i class="fas fa-heart"></i> Healthy & Unhealthy Habits</h2>
          <p><strong>Exercise:</strong> Option 2</p>
          <p><strong>Diet:</strong> Option 2</p>
          <p><strong>Alcohol Consumption:</strong> Option 2</p>
          <p><strong>Caffeine Consumption:</strong> Option 2</p>
          <p><strong>Do you smoke?:</strong> Option 2</p>
        </div>
      </div>
    </div>
  );
};

export default HistorialMedicoForm;
