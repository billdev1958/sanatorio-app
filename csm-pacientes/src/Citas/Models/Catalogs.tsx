export interface Services{
    id: number 
    name: string
}

export interface Shift{
    id: number 
    name: string
}

export interface Beneficiary {
    beneficiaryID: string; // Corresponde a uuid.UUID en Go
    fullName: string;
  }
  
  export interface PatientAndBeneficiaries {
    accountHolderID: string; // Corresponde a uuid.UUID en Go
    fullName: string;
    benefeciaries: Beneficiary[];
  }

export interface SchedulesAppointmentRequest {
    shift: number; // ID del turno (matutino, vespertino, etc.)
    service: number; // ID del servicio (Medicina General, etc.)
    appointmentDate: string; // Fecha completa en formato YYYY-MM-DD
}

export interface OfficeScheduleResponse {
    id: number;
    timeStart: string;
    timeEnd: string;
    timeDuration: string;
    officeName: string;
    statusID: number;
  }

export interface RegisterAppointmentRequest {
    scheduleID: number; // ID del horario
    patientID: string; // UUID del paciente
    beneficiaryID?: string; // UUID del beneficiario (opcional)
    timeStart: string; // Fecha y hora de inicio en formato ISO 8601
    timeEnd: string; // Fecha y hora de fin en formato ISO 8601
    reason?: string; // Motivo de la cita (opcional)
    symptoms?: string; // Síntomas (opcional)
  }
  