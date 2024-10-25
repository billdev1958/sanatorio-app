// src/components/PatientInfo.tsx
import { Component } from 'solid-js';

const PatientInfo: Component = () => {
  return (
    <div class="patient-info">
      <h2><i class="fas fa-user"></i> Patient Information</h2>
      <p><strong>Name:</strong> Reinaldos Prydden</p>
      <p><strong>Email:</strong> csword2@comcast.net</p>
      <p><strong>Birth Date:</strong> 1/10/2024</p>
      <p><strong>Height (cm):</strong> 4</p>
      <p><strong>Weight (kg):</strong> 4</p>
      <p><strong>Gender:</strong> Option 2</p>
    </div>
  );
};

export default PatientInfo;
