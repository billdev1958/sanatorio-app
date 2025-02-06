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
    scheduleID: number;
    patientID: string;
    beneficiaryID?: string | null; // Allow null or undefined
    timeStart: string;
    timeEnd: string;
    reason?: string | null;
    symptoms?: string | null;
}

export interface Appointment {
  AccountID: string;
  PatientID: string;
  BeneficiaryID: string | null;
  PatientName: string;
  OfficeName: string;
  ServiceName: string;
  TimeStart: string; // Formato ISO 8601, por ejemplo: "2025-02-05T10:00:00Z"
  TimeEnd: string;   // Formato ISO 8601, por ejemplo: "2025-02-05T11:00:00Z"
  StatusName: string;
}

export interface AppointmentByID {
  patientID: string;
  beneficiaryID?: string | null;
  timeStart: string;
  timeEnd: string;
  serviceID?: number;
  shiftID?: number;
  reason?: string;
  symptoms?: string;
}