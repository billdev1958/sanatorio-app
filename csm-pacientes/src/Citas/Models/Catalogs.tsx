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
    beneficiaries: Beneficiary[];
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

 