// Modelo para un día de la semana
export interface DayOfWeek {
    id: number;
    name: string;
  }
  
  // Modelo para un turno
  export interface CatShift {
    id: number;
    name: string;
  }
  
  // Modelo para un servicio
  export interface CatService {
    id: number;
    name: string;
  }
  
  // Modelo para una oficina
  export interface Office {
    office_id: number;
    office_name: string;
  }
  
  // Modelo para la respuesta principal del servicio
  export interface GetOfficeScheduleInfoResponse {
    status: string; // "success" o "error"
    message: string; // Mensaje de la respuesta
    data: {
      day_of_week: DayOfWeek[];
      cat_shift: CatShift[];
      cat_services: CatService[];
      doctor: null | any; // Doctor puede ser null u otro modelo en el futuro
      office: Office[];
    };
    errors?: any; // Campo opcional para errores
  }
  