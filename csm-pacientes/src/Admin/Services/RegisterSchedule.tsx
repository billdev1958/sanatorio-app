export interface DayOfWeek {
    id: number;
    name: string;
  }
  
  export interface CatShift {
    id: number;
    name: string;
  }
  
  export interface CatService {
    id: number;
    name: string;
  }
  
  export interface Office {
    office_id: number;
    office_name: string;
  }
  
  export interface Doctor {
    account_id: string;
    first_name: string;
    last_name_1: string;
    last_name_2: string;
  }
  
  export interface GetOfficeScheduleInfoResponse {
    day_of_week: DayOfWeek[];
    cat_shift: CatShift[];
    cat_services: CatService[];
    doctor: Doctor[] | null; // Puede ser null según el ejemplo proporcionado
    office: Office[];
  }
  
  export interface ApiResponse<T> {
    status: string;
    message?: string;
    data?: T;
    errors?: any;
  }
  
  // Tipo específico para la respuesta de este endpoint
  export type GetOfficeScheduleApiResponse = ApiResponse<GetOfficeScheduleInfoResponse>;
  
  export interface RegisterOfficeScheduleRequest {
    selectedDays: number[];
    timeStart: string;     // Formato HH:MM (ej. "09:00")
    timeEnd: string;       // Formato HH:MM (ej. "12:00")
    timeDuration: string;  // Duración en formato hh:mm (ej. "01:00")
    shiftID: string;       // ID del turno como string
    serviceID: string;     // ID del servicio como string
    doctorID: string;      // UUID del doctor
    officeID: number;      // ID de la oficina
    timeSlots: string[];   // Array de intervalos (ej. ["09:00 - 10:00", "10:00 - 11:00", ...])
  }